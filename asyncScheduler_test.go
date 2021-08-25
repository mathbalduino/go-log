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
var asyncHandleLogMutex = &sync.Mutex{}

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
		asyncHandleLogMutex.Lock()
		defer asyncHandleLogMutex.Unlock()

		// There's expected to exist 5 go routines, plus the TestSuite go routine,
		// that will be chained together. The 0, 1, 2, and 3 will try to lock a mutex
		// that can only be unlocked by the next go routine. Example:
		// 		TestSuite GoRoutine:
		//			locks mutex 0
		// 		GoRoutine 0:
		//			locks mutex 1
		// 		GoRoutine 1:
		//			locks mutex 2
		// 		GoRoutine 2:
		//			locks mutex 3
		// 		GoRoutine 3:
		//			locks mutex 4
		// 		GoRoutine 4:
		//			UNlocks mutex 4
		//			Last go routine alive and available (the others are locked)
		// 		GoRoutine 3:
		//			Unlocked by GoRoutine 4, UNlocks mutex 3
		// 		GoRoutine 2:
		//			Unlocked by GoRoutine 3, UNlocks mutex 2
		// 		GoRoutine 1:
		//			Unlocked by GoRoutine 2, UNlocks mutex 1
		// 		GoRoutine 0:
		//			Unlocked by GoRoutine 1, UNlocks mutex 0 (test suite)
		// This way, it's required that there's at least 5 go routines, otherwise it will deadlock

		// Start the test with the mutexes locked
		locks := []*sync.Mutex{{}, {}, {}, {}, {}}
		locks[0].Lock()
		locks[1].Lock()
		locks[2].Lock()
		locks[3].Lock()
		locks[4].Lock()

		i := uint64(0)
		oldAsyncHandleLog := AsyncHandleLog
		AsyncHandleLog = func(ctx context.Context, c <-chan Log, wg *sync.WaitGroup) {
			idx := atomic.AddUint64(&i, 1)
			if idx < 5 {
				locks[idx].Lock()
			}
			locks[idx - 1].Unlock()
		}

		DefaultAsyncScheduler(5, 1)
		locks[0].Lock() // waits for the chain reaction
		if i != 5 {
			t.Fatalf("Expected to spawn only 5 go routines")
		}
		AsyncHandleLog = oldAsyncHandleLog
	})
	t.Run("Should give the correct context interface to the spawned go routines", func(t *testing.T) {
		asyncHandleLogMutex.Lock()
		contextWithCancelMutex.Lock()
		defer func() {
			contextWithCancelMutex.Unlock()
			asyncHandleLogMutex.Unlock()
		}()

		realCtx, cancelFn := context.WithCancel(context.Background())
		contextWithCancel = func(parent context.Context) (context.Context, context.CancelFunc) {
			return realCtx, cancelFn
		}
		nGoRoutines := uint64(5)
		wg := &sync.WaitGroup{}
		wg.Add(int(nGoRoutines))
		oldAsyncHandleLog := AsyncHandleLog
		AsyncHandleLog = func(givenCtx context.Context, _ <-chan Log, _ *sync.WaitGroup) {
			if givenCtx != realCtx {
				t.Fatalf("Wrong context given")
			}
			wg.Done()
		}

		DefaultAsyncScheduler(nGoRoutines, 0)
		wg.Wait()
		AsyncHandleLog = oldAsyncHandleLog
	})
	t.Run("Should give the correct channel to each spawned go routine", func(t *testing.T) {
		asyncHandleLogMutex.Lock()
		defer asyncHandleLogMutex.Unlock()

		nGoRoutines := uint64(5)
		i := uint64(0)
		wg := &sync.WaitGroup{}
		wg.Add(int(nGoRoutines))
		channels := make([]<-chan Log, nGoRoutines)
		oldAsyncHandleLog := AsyncHandleLog
		AsyncHandleLog = func(_ context.Context, c <-chan Log, _ *sync.WaitGroup) {
			idx := atomic.AddUint64(&i, 1) - 1
			channels[idx] = c
			wg.Done()
		}

		DefaultAsyncScheduler(nGoRoutines, 0)
		wg.Wait()
		for i, chanI := range channels {
			for j := i + 1; j < len(channels); j++ {
				if reflect.ValueOf(chanI).Pointer() == reflect.ValueOf(channels[j]).Pointer() {
					t.Fatalf("Every go routine is expected to have a unique channel")
				}
			}
		}
		AsyncHandleLog = oldAsyncHandleLog
	})
	t.Run("Should give the correct waitGroup to the spawned go routines", func(t *testing.T) {
		asyncHandleLogMutex.Lock()
		defer asyncHandleLogMutex.Unlock()

		nGoRoutines := uint64(5)
		testSuiteWG := &sync.WaitGroup{}
		var givenWG *sync.WaitGroup
		testSuiteWG.Add(int(nGoRoutines))
		oldAsyncHandleLog := AsyncHandleLog
		AsyncHandleLog = func(_ context.Context, _ <-chan Log, receivedWG *sync.WaitGroup) {
			if receivedWG == nil {
				t.Fatalf("Not expected to pass nil WaitGroup to the spawned go routine")
			}
			if givenWG == nil {
				givenWG = receivedWG
			} else if reflect.ValueOf(givenWG).Pointer() != reflect.ValueOf(receivedWG).Pointer() {
				t.Fatalf("Expected to pass the same WaitGroup to all the spawned go routines")
			}
			testSuiteWG.Done()
		}

		DefaultAsyncScheduler(nGoRoutines, 0)
		testSuiteWG.Wait()
		AsyncHandleLog = oldAsyncHandleLog
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
		time.AfterFunc(time.Millisecond * 100, func() { wg.Done() })
		a.Shutdown() // deadlock in case of failure
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
