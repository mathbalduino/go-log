package loxeLog

import (
	"context"
	"reflect"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	rContextWithCancel = newResource()
	rAsyncHandleLog    = newResource()
	rNewWaitGroup      = newResource()
	rAtomicAddUint64   = newResource()
)

func TestDefaultAsyncScheduler(t *testing.T) {
	t.Run("Should return nil if its given 0 go routines", func(t *testing.T) {
		a := DefaultAsyncScheduler(0, 0)
		if a != nil {
			t.Fatalf("Expected to be nil")
		}
	})
	t.Run("Should create a unique channel (w/ the correct cap) for every go routine", func(t *testing.T) {
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
	t.Run("Should set the cancelFn to the one returned by the contextWithCancel", raceFreeTest(func(t *testing.T) {
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
	}, rContextWithCancel))
	t.Run("Should set a not nil waitGroup and increment it's counter by the number of goRoutines", raceFreeTest(func(t *testing.T) {
		calls := 0
		nGoRoutines := 3
		mockWG := &mockWaitGroup{mockAdd: func(i int) {
			calls += 1
			if i != nGoRoutines {
				t.Fatalf("Wrong delta added to the WaitGroup")
			}
		}}
		oldNewWaitGroup := newWaitGroup
		newWaitGroup = func() WaitGroup { return mockWG }

		a := DefaultAsyncScheduler(uint64(nGoRoutines), 0).(*asyncScheduler)
		if a.wg == nil {
			t.Fatalf("Expected to be not nil")
		}
		if calls != 1 {
			t.Fatalf("Expected to call wg.Add() incrementing the counter")
		}
		a.Shutdown()
		newWaitGroup = oldNewWaitGroup
	}, rNewWaitGroup))
	t.Run("Should spawn the correct number of go routines", raceFreeTest(func(t *testing.T) {
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
		AsyncHandleLog = func(ctx context.Context, c <-chan Log, wg WaitGroup) {
			idx := atomic.AddUint64(&i, 1)
			if idx < 5 {
				locks[idx].Lock()
			}
			locks[idx-1].Unlock()
		}

		DefaultAsyncScheduler(5, 1)
		locks[0].Lock() // waits for the chain reaction
		if i != 5 {
			t.Fatalf("Expected to spawn only 5 go routines")
		}
		AsyncHandleLog = oldAsyncHandleLog
	}, rAsyncHandleLog))
	t.Run("Should give the correct context interface to the spawned go routines", raceFreeTest(func(t *testing.T) {
		realCtx, cancelFn := context.WithCancel(context.Background())
		contextWithCancel = func(parent context.Context) (context.Context, context.CancelFunc) {
			return realCtx, cancelFn
		}
		nGoRoutines := uint64(5)
		wg := &sync.WaitGroup{}
		wg.Add(int(nGoRoutines))
		oldAsyncHandleLog := AsyncHandleLog
		AsyncHandleLog = func(givenCtx context.Context, _ <-chan Log, _ WaitGroup) {
			if givenCtx != realCtx {
				t.Fatalf("Wrong context given")
			}
			wg.Done()
		}

		DefaultAsyncScheduler(nGoRoutines, 0)
		wg.Wait()
		AsyncHandleLog = oldAsyncHandleLog
	}, rAsyncHandleLog, rContextWithCancel))
	t.Run("Should give the correct unique channel to each spawned go routine", raceFreeTest(func(t *testing.T) {
		nGoRoutines := uint64(5)
		i := uint64(0)
		wg := &sync.WaitGroup{}
		wg.Add(int(nGoRoutines))
		channels := make([]<-chan Log, nGoRoutines)
		oldAsyncHandleLog := AsyncHandleLog
		AsyncHandleLog = func(_ context.Context, c <-chan Log, _ WaitGroup) {
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
	}, rAsyncHandleLog))
	t.Run("Should give the correct waitGroup to the spawned go routines", raceFreeTest(func(t *testing.T) {
		nGoRoutines := uint64(5)
		testSuiteWG := &sync.WaitGroup{}
		var givenWG WaitGroup
		testSuiteWG.Add(int(nGoRoutines))
		oldAsyncHandleLog := AsyncHandleLog
		AsyncHandleLog = func(_ context.Context, _ <-chan Log, receivedWG WaitGroup) {
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
	}, rAsyncHandleLog))
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
		calls := 0
		wg := &mockWaitGroup{mockWait: func() { calls += 1 }}
		a := &asyncScheduler{wg: wg, cancelFn: func() {}}
		a.Shutdown()
		if calls != 1 {
			t.Fatalf("Expected to call wg.Wait() one time")
		}
	})
}

func TestNextChannel(t *testing.T) {
	t.Run("Should call the atomic.AddUint64 to increment the counter by one", raceFreeTest(func(t *testing.T) {
		chans := []chan Log{make(chan Log), make(chan Log), make(chan Log)}
		a := &asyncScheduler{chans: chans}
		calls := 0
		atomicAddUint64 = func(addr *uint64, delta uint64) (new uint64) {
			calls += 1
			if delta != 1 {
				t.Fatalf("Expected to increment the counter by one")
			}
			if reflect.ValueOf(addr).Pointer() != reflect.ValueOf(&a.nextChan).Pointer() {
				t.Fatalf("Wrong pointer given to atomic.AddUint64")
			}
			return 1
		}
		a.NextChannel()
		if calls != 1 {
			t.Fatalf("Expected to call atomic.AddUint64")
		}
		atomicAddUint64 = atomic.AddUint64
	}, rAtomicAddUint64))
	t.Run("Should round-robin the channels every call, starting from the first one", func(t *testing.T) {
		chans := []chan Log{make(chan Log), make(chan Log), make(chan Log)}
		a := &asyncScheduler{chans: chans}
		for _, chanI := range chans {
			chanJ := a.NextChannel()
			if reflect.ValueOf(chanI).Pointer() != reflect.ValueOf(chanJ).Pointer() {
				t.Fatalf("Expected to return the correct next channel")
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
	// TODO: Passar um mock do WaitGroup e usar ele
	// TODO: Retornar erros nos edge cases
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

type mockWaitGroup struct {
	mockWait func()
	mockDone func()
	mockAdd  func(i int)
}

func (f *mockWaitGroup) Wait() {
	if f.mockWait != nil {
		f.mockWait()
	}
}
func (f *mockWaitGroup) Done() {
	if f.mockDone != nil {
		f.mockDone()
	}
}
func (f *mockWaitGroup) Add(i int) {
	if f.mockAdd != nil {
		f.mockAdd(i)
	}
}

// Tests run in parallel, so it's required to control the concurrency over
// global variables (such as the mocked functions). The functions below handle
// it

type testResource struct {
	index uint64 // unique ID
	mutex *sync.Mutex
}

// generates a new mutex + unique ID
var newResource = func() func() testResource {
	i := uint64(0)
	return func() testResource {
		return testResource{
			atomic.AddUint64(&i, 1) - 1,
			&sync.Mutex{},
		}
	}
}()

// sort the resources in ID ascending order, and call "Lock". "Unlock" in the reverse order
func raceFreeTest(fn func(*testing.T), resources ...testResource) func(*testing.T) {
	return func(t *testing.T) {
		sort.Slice(resources, func(i, j int) bool { return resources[i].index < resources[j].index })
		for _, resource := range resources {
			resource.mutex.Lock()
			defer resource.mutex.Unlock() // Safe to call inside the loop
		}
		fn(t)
	}
}
