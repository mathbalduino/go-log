package logger

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Since I decided to cover the most use cases as it's possible, all the tests that
// override/read the value of some global variable MUST lock/unlock it
//
// It is done using the "raceFreeTest" high-order function, passing the locks at the
// end of the function, as a variadic argument
//
// If the function A is set to an internal global variable, and some test are overriding
// it, all tests that call this function A MUST pass the READ lock to the "raceFreeTest"
// high-order function, while the test that is writing to it passes the WRITE lock.
//
// Maybe it's too much complexity, but it was the only way to do some important test checks.
// Hope it pays off! kkkkkkkkkk

func TestDefaultAsyncScheduler(t *testing.T) {
	t.Run("Should return nil if its given 0 go routines", func(t *testing.T) {
		a := DefaultAsyncScheduler(0, 0)
		if a != nil {
			t.Fatalf("Expected to be nil")
		}
	})
	t.Run("Should create a unique channel (w/ the correct cap) for every go routine", raceFreeTest(func(t *testing.T) {
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
	}, rContextWithCancel, rNewWaitGroup, rAsyncHandleLog, rHandleLog))
	t.Run("Should set the first chan as the next one (nextChan field first value)", raceFreeTest(func(t *testing.T) {
		a := DefaultAsyncScheduler(1, 0).(*asyncScheduler)
		if a.nextChan != 0 {
			t.Fatalf("Expected to be zero")
		}
		a.Shutdown()
	}, rContextWithCancel, rNewWaitGroup, rAsyncHandleLog, rHandleLog))
	t.Run("Should set the cancelFn to the one returned by the contextWithCancel", raceFreeTest(func(t *testing.T) {
		ctx, cancelFn := context.WithCancel(context.Background())
		contextWithCancel = func(parent context.Context) (context.Context, context.CancelFunc) { return ctx, cancelFn }
		defer func() { contextWithCancel = context.WithCancel }()

		a := DefaultAsyncScheduler(1, 0).(*asyncScheduler)
		if reflect.ValueOf(cancelFn).Pointer() != reflect.ValueOf(a.cancelFn).Pointer() {
			t.Fatalf("Expected to be the same cancelFn")
		}
		a.Shutdown()
	}, wContextWithCancel, rNewWaitGroup, rAsyncHandleLog, rHandleLog))
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
		defer func() { newWaitGroup = oldNewWaitGroup }()
		newWaitGroup = func() WaitGroup { return mockWG }

		oldAsyncHandleLog := asyncHandleLog
		defer func() { asyncHandleLog = oldAsyncHandleLog }()
		wg := &sync.WaitGroup{}
		wg.Add(nGoRoutines)
		asyncHandleLog = func(ctx context.Context, c <-chan Log, _ WaitGroup) error {
			wg.Done()
			return nil
		}

		a := DefaultAsyncScheduler(uint64(nGoRoutines), 0).(*asyncScheduler)
		wg.Wait()
		if a.wg == nil {
			t.Fatalf("Expected to be not nil")
		}
		if calls != 1 {
			t.Fatalf("Expected to call wg.Add() incrementing the counter")
		}
		a.Shutdown()
	}, wNewWaitGroup, rContextWithCancel, rHandleLog, wAsyncHandleLog))
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
		oldAsyncHandleLog := asyncHandleLog
		defer func() { asyncHandleLog = oldAsyncHandleLog }()
		asyncHandleLog = func(ctx context.Context, c <-chan Log, wg WaitGroup) error {
			idx := atomic.AddUint64(&i, 1)
			if idx < 5 {
				locks[idx].Lock()
			}
			locks[idx-1].Unlock()
			return nil
		}

		DefaultAsyncScheduler(5, 1)
		locks[0].Lock() // waits for the chain reaction
		if i != 5 {
			t.Fatalf("Expected to spawn 5 go routines")
		}
	}, wAsyncHandleLog, rContextWithCancel, rNewWaitGroup))
	t.Run("Should give the correct context interface to the spawned go routines", raceFreeTest(func(t *testing.T) {
		realCtx, cancelFn := context.WithCancel(context.Background())
		contextWithCancel = func(parent context.Context) (context.Context, context.CancelFunc) { return realCtx, cancelFn }
		defer func() { contextWithCancel = context.WithCancel }()

		nGoRoutines := uint64(5)
		wg := &sync.WaitGroup{} // just to make the testSuite wait for the go routines
		wg.Add(int(nGoRoutines))
		oldAsyncHandleLog := asyncHandleLog
		defer func() { asyncHandleLog = oldAsyncHandleLog }()
		asyncHandleLog = func(givenCtx context.Context, _ <-chan Log, _ WaitGroup) error {
			if givenCtx != realCtx {
				t.Fatalf("Wrong context given")
			}
			wg.Done()
			return nil
		}

		DefaultAsyncScheduler(nGoRoutines, 0)
		wg.Wait()
	}, wAsyncHandleLog, wContextWithCancel, rNewWaitGroup))
	t.Run("Should give the correct unique channel to each spawned go routine", raceFreeTest(func(t *testing.T) {
		nGoRoutines := uint64(5)
		i := uint64(0)
		wg := &sync.WaitGroup{}
		wg.Add(int(nGoRoutines))
		channels := make([]<-chan Log, nGoRoutines)
		oldAsyncHandleLog := asyncHandleLog
		defer func() { asyncHandleLog = oldAsyncHandleLog }()
		asyncHandleLog = func(_ context.Context, c <-chan Log, _ WaitGroup) error {
			idx := atomic.AddUint64(&i, 1) - 1
			channels[idx] = c
			wg.Done()
			return nil
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
	}, wAsyncHandleLog, rNewWaitGroup, rContextWithCancel))
	t.Run("Should give the correct waitGroup to the spawned go routines", raceFreeTest(func(t *testing.T) {
		nGoRoutines := uint64(5)
		testSuiteWG := &sync.WaitGroup{}
		var givenWG WaitGroup
		testSuiteWG.Add(int(nGoRoutines))
		oldAsyncHandleLog := asyncHandleLog
		defer func() { asyncHandleLog = oldAsyncHandleLog }()
		asyncHandleLog = func(_ context.Context, _ <-chan Log, receivedWG WaitGroup) error {
			if receivedWG == nil {
				t.Fatalf("Not expected to pass nil WaitGroup to the spawned go routine")
			}
			if givenWG == nil {
				givenWG = receivedWG
			} else if reflect.ValueOf(givenWG).Pointer() != reflect.ValueOf(receivedWG).Pointer() {
				t.Fatalf("Expected to pass the same WaitGroup to all the spawned go routines")
			}
			testSuiteWG.Done()
			return nil
		}

		DefaultAsyncScheduler(nGoRoutines, 0)
		testSuiteWG.Wait()
	}, wAsyncHandleLog, rNewWaitGroup, rContextWithCancel))
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
		defer func() { atomicAddUint64 = atomic.AddUint64 }()

		a.NextChannel()
		if calls != 1 {
			t.Fatalf("Expected to call atomic.AddUint64")
		}
	}, wAtomicAddUint64))
	t.Run("Should round-robin the channels every call, starting from the first one", raceFreeTest(func(t *testing.T) {
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
	}, rAtomicAddUint64))
}

func TestAsyncHandleLog(t *testing.T) {
	t.Run("Should call the package internal variable directly, forwarding the arguments", raceFreeTest(func(t *testing.T) {
		oldAsyncHandleLog := asyncHandleLog
		defer func() { asyncHandleLog = oldAsyncHandleLog }()

		calls := 0
		expectedCtx := context.Background()
		expectedChan := make(chan Log)
		var expectedWg WaitGroup = &sync.WaitGroup{}
		expectedErr := fmt.Errorf("some error")
		asyncHandleLog = func(receivedCtx context.Context, receivedChan <-chan Log, receivedWg WaitGroup) error {
			calls += 1
			if receivedCtx != expectedCtx {
				t.Fatalf("Wrong context")
			}
			if reflect.ValueOf(receivedChan).Pointer() != reflect.ValueOf(expectedChan).Pointer() {
				t.Fatalf("Wrong channel")
			}
			if receivedWg != expectedWg {
				t.Fatalf("Wrong waitGroup")
			}
			return expectedErr
		}

		receivedErr := AsyncHandleLog(expectedCtx, expectedChan, expectedWg)
		if expectedErr != receivedErr {
			t.Fatalf("Wrong error")
		}
		if calls != 1 {
			t.Fatalf("Expected to call the internal implementation one time")
		}
	}, wAsyncHandleLog))
	t.Run("If the given WaitGroup is nil, return immediately", raceFreeTest(func(t *testing.T) {
		e := AsyncHandleLog(context.Background(), make(<-chan Log), nil)
		if e != ErrNilWaitGroup {
			t.Fatalf("Expected to return the correct error")
		}
	}, rAsyncHandleLog, rHandleLog))
	t.Run("If the given context is nil, call wg.Done() and return immediately", raceFreeTest(func(t *testing.T) {
		calls := 0
		wg := &mockWaitGroup{mockDone: func() { calls += 1 }}
		e := AsyncHandleLog(nil, make(<-chan Log), wg)
		if e != ErrNilCtx {
			t.Fatalf("Expected to return the correct error")
		}
		if calls != 1 {
			t.Fatalf("Expected to call wg.Done() before exiting")
		}
	}, rAsyncHandleLog, rHandleLog))
	t.Run("If the given channel is nil, call wg.Done() and return immediately", raceFreeTest(func(t *testing.T) {
		calls := 0
		wg := &mockWaitGroup{mockDone: func() { calls += 1 }}
		e := AsyncHandleLog(context.Background(), nil, wg)
		if e != ErrNilChan {
			t.Fatalf("Expected to return the correct error")
		}
		if calls != 1 {
			t.Fatalf("Expected to call wg.Done() before exiting")
		}
	}, rAsyncHandleLog, rHandleLog))
	t.Run("Should forward the logs received via channel to the 'handleLog' function", raceFreeTest(func(t *testing.T) {
		expectedLog := Log{
			lvl:         LvlDebug,
			msg:         "Some msg",
			logger:      &Logger{},
			preFields:   LogFields{"a": "aaa", "b": "bbb", "c": "ccc"},
			adHocFields: []LogFields{{"d": "ddd", "e": "eee", "f": "fff"}},
		}
		calls := 0
		oldHandleLog := handleLog
		defer func() { handleLog = oldHandleLog }()
		handleLog = func(receivedLog Log) {
			calls += 1
			if !reflect.DeepEqual(receivedLog, expectedLog) {
				t.Fatalf("Expected to receive a different log")
			}
		}

		c := make(chan Log)
		go AsyncHandleLog(context.Background(), c, &mockWaitGroup{})
		c <- expectedLog
		if calls != 1 {
			t.Fatalf("Expected to call handleLog")
		}
	}, wHandleLog, rAsyncHandleLog))
	t.Run("Should return when the context is done, decrementing the counter and returning nil", raceFreeTest(func(t *testing.T) {
		ctx, cancelFn := context.WithCancel(context.Background())
		calls := 0
		wg := &mockWaitGroup{mockDone: func() { calls += 1 }}
		time.AfterFunc(time.Millisecond*500, cancelFn)
		e := AsyncHandleLog(ctx, make(chan Log), wg)
		if e != nil {
			t.Fatalf("Error is expected to be nil")
		}
		if calls != 1 {
			t.Fatalf("Expected to decrement the waitGroup counter")
		}
	}, rAsyncHandleLog, rHandleLog))
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
