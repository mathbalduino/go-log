package logger

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("If there's some validation errors with the configuration, panic", func(t *testing.T) {
		calls := 0
		c := make(chan bool)
		go func() {
			defer func() {
				e := recover()
				if e == nil {
					t.Fatalf("Expected some error")
				}
				calls += 1
				if e.(error) != ErrLvlMsgSameKey {
					t.Fatalf("Expected the correct error")
				}
				c <- true
			}()

			New(Configuration{}) // same lvl and msg field name
		}()

		<-c
		if calls != 1 {
			t.Fatalf("Expected to call deferred function")
		}
	})
	t.Run("Should set a pointer to the given configuration and nil to the others", func(t *testing.T) {
		errorParser := func(err error) (string, LogFields) { return "", nil }
		lvlsEnabled := LvlFatal
		lvl := "lvl"
		msg := "msg"
		m := &mockAsyncScheduler{}
		l := New(Configuration{m, lvl, msg, lvlsEnabled, errorParser})
		if l.configuration == nil || l.outputs != nil || l.preHooks != nil || l.postHooks != nil || l.fields != nil {
			t.Fatalf("Expected to be nil")
		}
		if l.configuration.AsyncScheduler != AsyncScheduler(m) {
			t.Fatalf("Expected the correct AsyncScheduler")
		}
		if l.configuration.LvlFieldName != lvl {
			t.Fatalf("Expected the correct LvlFieldName")
		}
		if l.configuration.MsgFieldName != msg {
			t.Fatalf("Expected the correct MsgFieldName")
		}
		if l.configuration.LvlsEnabled != lvlsEnabled {
			t.Fatalf("Expected the correct LvlsEnabled")
		}
		if reflect.ValueOf(errorParser).Pointer() != reflect.ValueOf(l.configuration.ErrorParser).Pointer() {
			t.Fatalf("Expected the correct ErrorParser")
		}
	})
}

func TestNewDefault(t *testing.T) {
	t.Run("Should use the default configuration", func(t *testing.T) {
		c := DefaultConfig()
		l := NewDefault()
		if l.configuration.AsyncScheduler != c.AsyncScheduler {
			t.Fatalf("Expected the correct AsyncScheduler")
		}
		if l.configuration.LvlFieldName != c.LvlFieldName {
			t.Fatalf("Expected the correct LvlFieldName")
		}
		if l.configuration.MsgFieldName != c.MsgFieldName {
			t.Fatalf("Expected the correct MsgFieldName")
		}
		if l.configuration.LvlsEnabled != c.LvlsEnabled {
			t.Fatalf("Expected the correct LvlsEnabled")
		}
		if reflect.ValueOf(l.configuration.ErrorParser).Pointer() != reflect.ValueOf(c.ErrorParser).Pointer() {
			t.Fatalf("Expected the correct ErrorParser")
		}
	})
	t.Run("Should set two outputs, the first should log to stdout and the second panic on LvlFatal", raceFreeTest(func(t *testing.T) {
		l := NewDefault()
		if len(l.outputs) != 2 {
			t.Fatalf("Expected to set some output")
		}
		if reflect.ValueOf(l.outputs[0]).Pointer() != reflect.ValueOf(OutputAnsiToStdout).Pointer() {
			t.Fatal("The first output is wrong")
		}
		if reflect.ValueOf(l.outputs[1]).Pointer() != reflect.ValueOf(OutputPanicOnFatal).Pointer() {
			t.Fatal("The second output is wrong")
		}
	}, wStdOut))
}
