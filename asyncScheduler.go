package loxeLog

import "sync/atomic"

type AsyncScheduler interface {
	NextChannel() chan<- Log
}

func DefaultAsyncScheduler(nGoRoutines uint64, chanCap uint64) AsyncScheduler {
	if nGoRoutines == 0 {
		return nil
	}

	scheduler := &asyncScheduler{
		make([]chan Log, nGoRoutines),
		0,
	}
	for i := range scheduler.chans {
		scheduler.chans[i] = make(chan Log, chanCap)
		go AsyncHandleLog(scheduler.chans[i])
	}

	return scheduler
}

// -----

type asyncScheduler struct {
	chans    []chan Log
	nextChan uint64
}

func (a *asyncScheduler) NextChannel() chan<- Log {
	currChannel := (atomic.AddUint64(&a.nextChan, 1) - 1) % uint64(len(a.chans)) // TODO: avoid overflow
	return a.chans[currChannel]
}

func AsyncHandleLog(c <-chan Log) {
	if c == nil {
		return
	}

	for {
		log := <-c
		handleLog(log)
	}
}
