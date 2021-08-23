package loxeLog

import (
	"context"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestDefaultAsyncScheduler(t *testing.T) {
	t.Run("Should return nil if its given 0 go routines", func(t *testing.T) {
		a := DefaultAsyncScheduler(0, 0)
		if a != nil {
			t.Fatalf("Expected to be nil")
		}
	})
	// TODO: completar o teste abaixo setando um time.Sleep alto no primeiro e ir baixando até o ultimo,
	//		 pegando o ID da go routine. No final, nenhum ID pode ser repetido e devem ser 5
	t.Run("Should spawn n go routines", func(t *testing.T) {
		a := DefaultAsyncScheduler(5, 0).(*asyncScheduler)
		calls := 0
		a.chans[1] <- Log{
			logger: &Logger{
				configuration: &Configuration{},
				outputs: []Output{func(lvl_ uint64, msg_ string, fields_ LogFields) {
					calls += 1
				}},
				fields: LogFields{},
			},
		}
		if calls != 1 {
			t.Fatalf("Expected to spawn the go routines")
		}
		a.Shutdown()
	})
	t.Run("Should create a channel for every go routine", func(t *testing.T) {
		nChans := uint64(5)
		chansCap := uint64(3)
		a := DefaultAsyncScheduler(nChans, 3).(*asyncScheduler)
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
	t.Run("Should set a not nil waitGroup", func(t *testing.T) {
		a := DefaultAsyncScheduler(1, 0).(*asyncScheduler)
		if a.wg == nil {
			t.Fatalf("Expected to be not nil")
		}
		a.Shutdown()
	})
}

// TODO: criar um caso de teste para cada metodo do asyncScheduler

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
