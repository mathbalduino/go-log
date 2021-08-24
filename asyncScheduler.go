package loxeLog

import (
	"context"
	"sync"
	"sync/atomic"
)

// AsyncScheduler is used as a source of channels
// that are used to send new Logs to worker goroutines,
// handling Logs in an asynchronous way
type AsyncScheduler interface {
	// NextChannel must return a valid, non-nil,
	// receive-only channel
	NextChannel() chan<- Log

	// Shutdown must send a signal (and wait response)
	// to the running go routines, exiting them
	Shutdown()
}

// DefaultAsyncScheduler will create one channel by goroutine, with the given
// capacity, and setup a goroutine that will handle newly created Logs.
//
// Note that if nGoRoutines is zero, nothing happens and the returned AsyncScheduler
// will be nil
func DefaultAsyncScheduler(nGoRoutines uint64, chanCap uint64) AsyncScheduler {
	if nGoRoutines == 0 {
		return nil
	}

	ctx, cancelFn := contextWithCancel(context.Background())
	scheduler := &asyncScheduler{
		make([]chan Log, nGoRoutines),
		0,
		cancelFn,
		&sync.WaitGroup{},
	}
	scheduler.wg.Add(int(nGoRoutines))
	for i := range scheduler.chans {
		scheduler.chans[i] = make(chan Log, chanCap)
		go AsyncHandleLog(ctx, scheduler.chans[i], scheduler.wg)
	}

	return scheduler
}

var contextWithCancel = context.WithCancel

// -----

// asyncScheduler is a default implementation for
// async log handling, that uses a round-robin like
// scheduling scheme
type asyncScheduler struct {
	// For every channel, there is a goroutine
	chans []chan Log

	// Stores the next goroutine responsible to
	// handle the next Log.
	//
	// Note that this variable can overflow, but
	// it's not a big deal, just apply mod(n_channels)
	nextChan uint64

	// When called, will close the go routines context
	// Done() channel, exiting them
	cancelFn context.CancelFunc

	// Used to wait for the go routines exit
	wg *sync.WaitGroup
}

// Shutdown will call the cancel function, closing the go
// routines context channel, and wait for them to exit (via waitGroup)
func (a *asyncScheduler) Shutdown() {
	a.cancelFn()
	a.wg.Wait()
}

// NextChannel selects the next channel to be used,
// using a round-robin-like scheduling scheme, applying
// some mod operation to avoid overflow issues
func (a *asyncScheduler) NextChannel() chan<- Log {
	currChannel := (atomicAddUint64(&a.nextChan, 1) - 1) % uint64(len(a.chans))
	return a.chans[currChannel]
}

// just to ease tests
var atomicAddUint64 = atomic.AddUint64

// AsyncHandleLog will wait on the given send-only
// channel, and forwarding any received Log to the
// internal "handleLog" function.
//
// Note that this function must be used to implement
// custom async strategies, since it's the only way
// to access the internal "handleLog" function
//
// Using "var" just to ease tests
var AsyncHandleLog = func(ctx context.Context, c <-chan Log, wg *sync.WaitGroup) {
	if wg == nil {
		return
	}

	defer wg.Done()
	if ctx == nil || c == nil {
		return
	}

	for {
		select {
		case log := <-c:
			handleLog(log)
		case <-ctx.Done():
			return
		}
	}
}
