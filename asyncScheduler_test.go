package loxeLog

import (
	"context"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var contextWithCancelMutex = &sync.Mutex{}

func TestDefaultAsyncScheduler(t *testing.T) {
	t.Run("Should return nil if its given 0 go routines", func(t *testing.T) {
		a := DefaultAsyncScheduler(0, 0)
		if a != nil {
			t.Fatalf("Expected to be nil")
		}
	})
	t.Run("Should spawn n go routines", func(t *testing.T) {
		// There's expected to exist 5 go routines, that will be chained together. The 0, 1, 2, and 3 will
		// try to lock a mutex that can only be unlocked by the next go routine. Example:
		// 		GoRoutine 0:
		//			locks mutex 0
		//			UNlocks test suite mutex
		// 		GoRoutine 1:
		//			locks mutex 1
		//			UNlocks mutex 0
		// 		GoRoutine 2:
		//			locks mutex 2
		//			UNlocks mutex 1
		// 		GoRoutine 3:
		//			locks mutex 3
		//			UNlocks mutex 2
		// 		GoRoutine 4:
		//			UNlocks mutex 3
		//
		// This way, it's required that there's at least 5 go routines, otherwise it will deadlock
		// Note that it's expected that the scheduling between the go routines is correct

		a := DefaultAsyncScheduler(5, 0).(*asyncScheduler)
		locks := []*sync.Mutex{{}, {}, {}, {}}
		locks[0].Lock()
		locks[1].Lock()
		locks[2].Lock()
		locks[3].Lock()
		nextIdx := 0
		callOrder := []int{0, 0, 0, 0, 0}

		testLock := sync.Mutex{}
		testLock.Lock()
		a.chans[0] <- Log{
			logger: &Logger{
				configuration: &Configuration{},
				outputs: []Output{func(lvl_ uint64, msg_ string, fields_ LogFields) {
					locks[0].Lock()
					callOrder[nextIdx] += 0
					nextIdx += 1
					testLock.Unlock()
				}},
			},
		}
		a.chans[1] <- Log{
			logger: &Logger{
				configuration: &Configuration{},
				outputs: []Output{func(lvl_ uint64, msg_ string, fields_ LogFields) {
					locks[1].Lock()
					callOrder[nextIdx] += 1
					nextIdx += 1
					locks[0].Unlock()
				}},
			},
		}
		a.chans[2] <- Log{
			logger: &Logger{
				configuration: &Configuration{},
				outputs: []Output{func(lvl_ uint64, msg_ string, fields_ LogFields) {
					locks[2].Lock()
					callOrder[nextIdx] += 2
					nextIdx += 1
					locks[1].Unlock()
				}},
			},
		}
		a.chans[3] <- Log{
			logger: &Logger{
				configuration: &Configuration{},
				outputs: []Output{func(lvl_ uint64, msg_ string, fields_ LogFields) {
					locks[3].Lock()
					callOrder[nextIdx] += 3
					nextIdx += 1
					locks[2].Unlock()
				}},
			},
		}
		a.chans[4] <- Log{
			logger: &Logger{
				configuration: &Configuration{},
				outputs: []Output{func(lvl_ uint64, msg_ string, fields_ LogFields) {
					callOrder[nextIdx] += 4
					nextIdx += 1
					locks[3].Unlock()
				}},
			},
		}

		testLock.Lock()
		if callOrder[0] != 4 || callOrder[1] != 3 || callOrder[2] != 2 || callOrder[3] != 1 || callOrder[4] != 0 {
			t.Fatalf("Expected to spawn the go routines")
		}
		a.Shutdown()
	})
	t.Run("Should create a channel for every go routine", func(t *testing.T) {
		nChans := uint64(5)
		chansCap := uint64(3)
		a := DefaultAsyncScheduler(nChans, chansCap).(*asyncScheduler)
		if len(a.chans) != int(nChans) {
			t.Fatalf("Expected a right-sized channels slice")
		}
		for i, chanI := range a.chans {
			if cap(chanI) != int(chansCap) {
				t.Fatalf("Expected to create a channel with the given capacity")
			}
			for j := i + 1; j < len(a.chans); j++ {
				if reflect.ValueOf(chanI).Pointer() == reflect.ValueOf(a.chans[j]).Pointer() {
					t.Fatalf("Not expected to repeat the channels inside the slice")
				}
			}
		}
		a.Shutdown()
	})
	t.Run("Should set the first chan as the next one", func(t *testing.T) {
		a := DefaultAsyncScheduler(1, 0).(*asyncScheduler)
		if a.nextChan != 0 {
			t.Fatalf("Expected to be zero")
		}
		a.Shutdown()
	})
	t.Run("Should set a not nil cancelFn", func(t *testing.T) {
		a := DefaultAsyncScheduler(1, 0).(*asyncScheduler)
		if a.cancelFn == nil {
			t.Fatalf("Expected to be not nil")
		}
		a.Shutdown()
	})
	t.Run("Should set the cancelFn to the one returned by the contextWithCancel", func(t *testing.T) {
		contextWithCancelMutex.Lock()
		defer contextWithCancelMutex.Unlock()

		ctx, cancelFn := context.WithCancel(context.Background())
		contextWithCancel = func(parent context.Context) (context.Context, context.CancelFunc) {
			return ctx, cancelFn
		}
		a := DefaultAsyncScheduler(1, 0).(*asyncScheduler)
		if reflect.ValueOf(cancelFn).Pointer() != reflect.ValueOf(a.cancelFn).Pointer() {
			t.Fatalf("Expected to be the same cancelFn")
		}
		a.Shutdown()
		contextWithCancel = context.WithCancel
	})
	// TODO: Criar caso de teste para garantir que o contexto certo está sendo passado para AsyncHandleLog
	t.Run("Should set a not nil waitGroup", func(t *testing.T) {
		a := DefaultAsyncScheduler(1, 0).(*asyncScheduler)
		if a.wg == nil {
			t.Fatalf("Expected to be not nil")
		}
		a.Shutdown()
	})
}

func TestShutdown(t *testing.T) {
	t.Run("Should call the cancel function, notifying the go routines context to exit", func(t *testing.T) {
		wg := &sync.WaitGroup{}
		calls := 0
		a := &asyncScheduler{wg: wg, cancelFn: func() {
			calls += 1
		}}
		a.Shutdown()
		if calls != 1 {
			t.Fatalf("Expected to call the cancel function")
		}
	})
	t.Run("Should call the wg.Wait() method", func(t *testing.T) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		a := &asyncScheduler{wg: wg, cancelFn: func() {}}
		time.AfterFunc(time.Second, func() { wg.Done() })
		c := make(chan bool)
		go func() {
			a.Shutdown()
			c <- true
		}()
		waited := false
		for {
			select {
			case <-c:
				if !waited {
					t.Fatalf("Expected to wait at wg.Wait()")
				}
				return
			case <-time.After(time.Second):
				waited = true
			}
		}
	})
}

var atomicAddUint64Mutex = &sync.Mutex{}

func TestNextChannel(t *testing.T) {
	t.Run("Should call the atomic.AddUint64 to increment the counter by one", func(t *testing.T) {
		atomicAddUint64Mutex.Lock()
		defer atomicAddUint64Mutex.Unlock()

		chans := []chan Log{make(chan Log), make(chan Log), make(chan Log)}
		a := &asyncScheduler{chans: chans}
		calls := 0
		atomicAddUint64 = func(addr *uint64, delta uint64) (new uint64) {
			calls += 1
			if delta != 1 {
				t.Fatalf("Expected to increment the counter by one")
			}
			if reflect.ValueOf(addr).Pointer() != reflect.ValueOf(&a.nextChan).Pointer() {
				t.Fatalf("Wrong atomic.AddUint64 pointer")
			}
			return 1
		}
		a.NextChannel()
		if calls != 1 {
			t.Fatalf("Expected to call atomic.AddUint64")
		}
		atomicAddUint64 = atomic.AddUint64
	})
	t.Run("Should round-robin the channels for every call", func(t *testing.T) {
		chans := []chan Log{make(chan Log), make(chan Log), make(chan Log)}
		a := &asyncScheduler{chans: chans}
		for _, chanI := range chans {
			chanJ := a.NextChannel()
			if reflect.ValueOf(chanI).Pointer() != reflect.ValueOf(chanJ).Pointer() {
				t.Fatalf("Expected to return the right next channel")
			}
		}
		// Another round
		for _, chanI := range chans {
			chanJ := a.NextChannel()
			if reflect.ValueOf(chanI).Pointer() != reflect.ValueOf(chanJ).Pointer() {
				t.Fatalf("Expected to return the right next channel")
			}
		}
	})
}

func TestAsyncHandleLog(t *testing.T) {
	t.Run("If the given context is nil, return immediately, and calling wg.Done()", func(t *testing.T) {
		c := make(chan bool)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			AsyncHandleLog(nil, make(<-chan Log), wg)
			c <- true
		}()
		select {
		case <-c:
			wg.Wait() // not expected to deadlock
			return
		case <-time.After(time.Second):
			t.Fatalf("Expected to not enter the loop")
		}
	})
	t.Run("If the given chan is nil, return immediately, and calling wg.Done()", func(t *testing.T) {
		c := make(chan bool)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			AsyncHandleLog(context.Background(), nil, wg)
			c <- true
		}()
		select {
		case <-c:
			wg.Wait() // not expected to deadlock
			return
		case <-time.After(time.Second):
			t.Fatalf("Expected to not enter the loop")
		}
	})
	t.Run("If the given WaitGroup is nil, return immediately", func(t *testing.T) {
		c := make(chan bool)
		go func() {
			AsyncHandleLog(context.Background(), make(<-chan Log), nil)
			c <- true
		}()
		select {
		case <-c:
			return
		case <-time.After(time.Second):
			t.Fatalf("Expected to not enter the loop")
		}
	})
	t.Run("Should handle Logs passed by the channel, sending them to the output", func(t *testing.T) {
		logChan := make(chan Log)
		ctx, cancelFn := context.WithCancel(context.Background())
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go AsyncHandleLog(ctx, logChan, wg)

		lvl := LvlDebug
		msg := "Some msg"
		fields := LogFields{"a": "aaa", "b": "bbb", "d": "ddd"}
		calls := 0
		log := Log{
			lvl: lvl,
			msg: msg,
			logger: &Logger{
				configuration: &Configuration{},
				outputs: []Output{func(lvl_ uint64, msg_ string, fields_ LogFields) {
					calls += 1
					if lvl_ != lvl || msg_ != msg || !reflect.DeepEqual(fields_, fields_) {
						t.Fatalf("Expected to pass the correct Log to the outputs")
					}
				}},
				fields: fields,
			},
		}
		logChan <- log

		if calls != 1 {
			t.Fatalf("Expected to give the Log to the outputs")
		}
		cancelFn()
	})
	t.Run("Should return when the context is done, decrementing the semaphore", func(t *testing.T) {
		c := make(chan bool)
		ctx, cancelFn := context.WithCancel(context.Background())
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			AsyncHandleLog(ctx, make(chan Log), wg)
			c <- true
		}()

		cancelFn()
		select {
		case <-c:
			wg.Wait() // not expected to deadlock
			return
		case <-time.After(time.Second):
			t.Fatalf("Expected to exit")
		}
	})
}
