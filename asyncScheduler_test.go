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
	t.Run("Should create a channel for every go routine", func(t *testing.T) {
		nGoRoutines := uint64(5)
		chansCap := uint64(3)
		a := DefaultAsyncScheduler(nGoRoutines, chansCap).(*asyncScheduler)
		if len(a.chans) != int(nGoRoutines) {
			t.Fatalf("Expected one channel for every go routine")
		}
		for i, chanI := range a.chans {
			if cap(chanI) != int(chansCap) {
				t.Fatalf("Expected to create the channels with the correct capacity")
			}
			for j := i + 1; j < len(a.chans); j++ {
				if reflect.ValueOf(chanI).Pointer() == reflect.ValueOf(a.chans[j]).Pointer() {
					t.Fatalf("Every go routine is expected to have a unique channel")
				}
			}
		}
		a.Shutdown()
	})
	t.Run("Should set the first chan as the next one (nextChan field first value)", func(t *testing.T) {
		a := DefaultAsyncScheduler(1, 0).(*asyncScheduler)
		if a.nextChan != 0 {
			t.Fatalf("Expected to be zero")
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
	t.Run("Should set a not nil waitGroup and add the number of goRoutines", func(t *testing.T) {
		nGoRoutines := 3
		a := DefaultAsyncScheduler(3, 0).(*asyncScheduler)
		if a.wg == nil {
			t.Fatalf("Expected to be not nil")
		}

		panicked := false
		doneCalls := 0
		c := make(chan bool)
		go func() {
			defer func() {
				e := recover()
				if e == nil {
					t.Fatalf("Expected to panic when the Done decrements the counter beolow zero")
				}
				panicked = true
				c <- true
			}()

			// Add one to cause an exception
			for i := 0; i < nGoRoutines+1; i++ {
				a.wg.Done()
				doneCalls += 1
			}
		}()
		<-c
		if !panicked {
			t.Fatalf("Expected to panic when the Done decrements the counter beolow zero")
		}
		if doneCalls != nGoRoutines {
			t.Fatalf("Expected to decrement the counter %d times", nGoRoutines)
		}

		// Restore the WaitGroup to avoid another exception on Shutdown (plus one because of on-purpose overflow above)
		a.wg.Add(nGoRoutines + 1)
		a.Shutdown()
	})
	t.Run("Should spawn the correct number of go routines", func(t *testing.T) {
		// TODO: Mock the AsyncHadleLog function and use a switch statement. Avoid using the channels
		// There's expected to exist 5 go routines, that will be chained together. The 0, 1, 2, and 3 will
		// try to lock a mutex that can only be unlocked by the next go routine. Example:
		// 		GoRoutine 0:
		//			locks mutex 0
		// 		GoRoutine 1:
		//			locks mutex 1
		// 		GoRoutine 2:
		//			locks mutex 2
		// 		GoRoutine 3:
		//			locks mutex 3
		// 		GoRoutine 4:
		//			UNlocks mutex 3
		//			Last go routine alive and available (the others are locked)
		// 		GoRoutine 3:
		//			Unlocked by GoRoutine 4, UNlocks mutex 2
		// 		GoRoutine 2:
		//			Unlocked by GoRoutine 3, UNlocks mutex 1
		// 		GoRoutine 1:
		//			Unlocked by GoRoutine 2, UNlocks mutex 0
		// 		GoRoutine 0:
		//			Unlocked by GoRoutine 1, UNlocks TestSuite GoRoutine
		// This way, it's required that there's at least 5 go routines, otherwise it will deadlock

		a := DefaultAsyncScheduler(5, 1).(*asyncScheduler)
		locks := []*sync.Mutex{{}, {}, {}, {}}

		// Start the test with the mutexes locked
		testSuiteLock := sync.Mutex{}
		testSuiteLock.Lock()
		locks[0].Lock()
		locks[1].Lock()
		locks[2].Lock()
		locks[3].Lock()

		// Race conditions are not expected over "calls" variable (go routines are chained)
		// Trouble using "for" statement. Hardcoded just to get it done
		calls := 0
		a.chans[0] <- Log{
			logger: &Logger{
				configuration: &Configuration{},
				outputs: []Output{func(lvl_ uint64, msg_ string, fields_ LogFields) {
					locks[0].Lock()
					calls += 1
					testSuiteLock.Unlock()
				}},
			},
		}
		a.chans[1] <- Log{
			logger: &Logger{
				configuration: &Configuration{},
				outputs: []Output{func(lvl_ uint64, msg_ string, fields_ LogFields) {
					locks[1].Lock()
					calls += 1
					locks[0].Unlock()
				}},
			},
		}
		a.chans[2] <- Log{
			logger: &Logger{
				configuration: &Configuration{},
				outputs: []Output{func(lvl_ uint64, msg_ string, fields_ LogFields) {
					locks[2].Lock()
					calls += 1
					locks[1].Unlock()
				}},
			},
		}
		a.chans[3] <- Log{
			logger: &Logger{
				configuration: &Configuration{},
				outputs: []Output{func(lvl_ uint64, msg_ string, fields_ LogFields) {
					locks[3].Lock()
					calls += 1
					locks[2].Unlock()
				}},
			},
		}
		a.chans[4] <- Log{
			logger: &Logger{
				configuration: &Configuration{},
				outputs: []Output{func(lvl_ uint64, msg_ string, fields_ LogFields) {
					calls += 1
					locks[3].Unlock()
				}},
			},
		}

		// Expected to be UNlocked by the first spawned go routine
		testSuiteLock.Lock()
		if calls != 5 {
			t.Fatalf("Expected to spawn the given number of go routines")
		}
		a.Shutdown()
	})
	t.Run("Should give the correct context interface to the spawned go routines", func(t *testing.T) {
		t.Fatalf("implement-me")
	})
	t.Run("Should give the correct channel to each spawned go routine", func(t *testing.T) {
		t.Fatalf("implement-me")
	})
	t.Run("Should give the correct waitGroup to the spawned go routines", func(t *testing.T) {
		t.Fatalf("implement-me")
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
		// TODO: Melhorar esse teste. Ta ruim, as vezes pode falhar mesmo estando correto por causa dos timers
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
			case <-time.After(time.Millisecond * 500):
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
	t.Run("Should round-robin the channels for every call, starting from the first one", func(t *testing.T) {
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
	// TODO: usar a tecnica de causar um panic ao decrementar o counter do waitGroup pra zero
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
	t.Run("If the given context is nil, call wg.Done() and return immediately", func(t *testing.T) {
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
	t.Run("If the given chan is nil, call wg.Done() and return immediately", func(t *testing.T) {
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
	t.Run("Should forward the logs received via channel to the 'handleLog' function", func(t *testing.T) {
		// TODO: mock handleLog function
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
		case <-time.After(time.Second * 10): // safe-exit for deadlocks
			t.Fatalf("Deadlock!")
		}
	})
}
