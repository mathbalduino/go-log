package logger

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLog(t *testing.T) {
	t.Run("Should return immediately if the lvl is not enabled", raceFreeTest(func(t *testing.T) {
		oldHandleLog := handleLog
		defer func() { handleLog = oldHandleLog }()

		calls := 0
		handleLog = func(log Log) { calls += 1 }
		c := make(chan Log)
		m := &mockAsyncScheduler{mockNextChannel: func() chan<- Log {
			calls += 1
			return c
		}}
		l := &logger{configuration: &Configuration{AsyncScheduler: m, LvlsEnabled: LvlTrace}}
		l.Log(LvlDebug, "", nil)
		if len(c) != 0 || calls != 0 {
			t.Fatalf("Expected to return immediately")
		}
	}, wHandleLog))
	t.Run("Should give the correct log to the next channel if it's async, and not call handleLog", raceFreeTest(func(t *testing.T) {
		oldHandleLog := handleLog
		defer func() { handleLog = oldHandleLog }()
		handleLogCalls := 0
		handleLog = func(log Log) { handleLogCalls += 1 }

		lvl := LvlTrace
		msg := "some msg"
		adHocFields := []LogFields{
			{"a": "aaa", "b": "bbb", "c": "ccc"},
			{"d": "ddd", "e": "eee"},
			{"f": "fff", "g": "ggg", "h": "hhh"},
		}
		fnI := func(log Log) interface{} { return "iii" }
		fnJ := func(log Log) interface{} { return "jjj" }
		fnK := func(log Log) interface{} { return "kkk" }
		calls := 0
		c := make(chan Log, 1)
		m := &mockAsyncScheduler{mockNextChannel: func() chan<- Log {
			calls += 1
			return c
		}}
		l := &logger{
			configuration: &Configuration{AsyncScheduler: m, LvlsEnabled: lvl},
			preHooks:      Hooks{"i": fnI, "j": fnJ, "k": fnK},
		}

		l.Log(lvl, msg, adHocFields)
		receivedLog := <-c
		if calls != 1 {
			t.Fatalf("Expected to call NextChannel")
		}
		if handleLogCalls != 0 {
			t.Fatalf("Expected to not call handleLog")
		}
		if receivedLog.lvl != lvl {
			t.Fatalf("Expected to create a log with the correct lvl")
		}
		if receivedLog.msg != msg {
			t.Fatalf("Expected to create a log with the correct msg")
		}
		if reflect.ValueOf(receivedLog.logger).Pointer() != reflect.ValueOf(l).Pointer() {
			t.Fatalf("Expected to create a log that points to the correct src")
		}
		if !reflect.DeepEqual(receivedLog.preFields, LogFields{"i": "iii", "j": "jjj", "k": "kkk"}) {
			t.Fatalf("Expected to create a log with the correct pre fields")
		}
		if !reflect.DeepEqual(receivedLog.adHocFields, adHocFields) {
			t.Fatalf("Expected to create a log with the correct adHoc fields")
		}
		if receivedLog.postFields != nil {
			t.Fatalf("Expected postFields to be nil")
		}
	}, wHandleLog))
	t.Run("Should give the correct log to handleLog directly, if it's sync (nil AsyncScheduler)", func(t *testing.T) {
		lvl := LvlTrace
		msg := "some msg"
		adHocFields := []LogFields{
			{"a": "aaa", "b": "bbb", "c": "ccc"},
			{"d": "ddd", "e": "eee"},
			{"f": "fff", "g": "ggg", "h": "hhh"},
		}
		fnI := func(log Log) interface{} { return "iii" }
		fnJ := func(log Log) interface{} { return "jjj" }
		fnK := func(log Log) interface{} { return "kkk" }
		l := &logger{
			configuration: &Configuration{LvlsEnabled: lvl},
			preHooks:      Hooks{"i": fnI, "j": fnJ, "k": fnK},
		}
		calls := 0
		handleLog = func(receivedLog Log) {
			calls += 1
			if receivedLog.lvl != lvl {
				t.Fatalf("Expected to create a log with the correct lvl")
			}
			if receivedLog.msg != msg {
				t.Fatalf("Expected to create a log with the correct msg")
			}
			if reflect.ValueOf(receivedLog.logger).Pointer() != reflect.ValueOf(l).Pointer() {
				t.Fatalf("Expected to create a log that points to the correct src")
			}
			if !reflect.DeepEqual(receivedLog.preFields, LogFields{"i": "iii", "j": "jjj", "k": "kkk"}) {
				t.Fatalf("Expected to create a log with the correct pre fields")
			}
			if !reflect.DeepEqual(receivedLog.adHocFields, adHocFields) {
				t.Fatalf("Expected to create a log with the correct adHoc fields")
			}
			if receivedLog.postFields != nil {
				t.Fatalf("Expected postFields to be nil")
			}
		}

		l.Log(lvl, msg, adHocFields)
		if calls != 1 {
			t.Fatalf("Expected to call handleLog")
		}
	})
	t.Run("Trace method should create a log with LvlTrace and forward msg and adHocFields", func(t *testing.T) {
		msg := "some msg"
		adHocFields := []LogFields{
			{"a": "aaa", "b": "bbb", "c": "ccc"},
			{"d": "ddd", "e": "eee"},
			{"f": "fff", "g": "ggg", "h": "hhh"},
		}
		c := make(chan Log, 1)
		m := &mockAsyncScheduler{mockNextChannel: func() chan<- Log { return c }}
		l := &logger{configuration: &Configuration{AsyncScheduler: m, LvlsEnabled: LvlTrace}}
		l.Trace(msg, adHocFields...)
		receivedLog := <-c
		if receivedLog.lvl != LvlTrace {
			t.Fatalf("Expected to create a log with the correct lvl")
		}
		if receivedLog.msg != msg {
			t.Fatalf("Expected to create a log with the correct msg")
		}
		if !reflect.DeepEqual(receivedLog.adHocFields, adHocFields) {
			t.Fatalf("Expected to create a log with the correct adHoc fields")
		}
	})
	t.Run("Debug method should create a log with LvlDebug and forward msg and adHocFields", func(t *testing.T) {
		msg := "some msg"
		adHocFields := []LogFields{
			{"a": "aaa", "b": "bbb", "c": "ccc"},
			{"d": "ddd", "e": "eee"},
			{"f": "fff", "g": "ggg", "h": "hhh"},
		}
		c := make(chan Log, 1)
		m := &mockAsyncScheduler{mockNextChannel: func() chan<- Log { return c }}
		l := &logger{configuration: &Configuration{AsyncScheduler: m, LvlsEnabled: LvlDebug}}
		l.Debug(msg, adHocFields...)
		receivedLog := <-c
		if receivedLog.lvl != LvlDebug {
			t.Fatalf("Expected to create a log with the correct lvl")
		}
		if receivedLog.msg != msg {
			t.Fatalf("Expected to create a log with the correct msg")
		}
		if !reflect.DeepEqual(receivedLog.adHocFields, adHocFields) {
			t.Fatalf("Expected to create a log with the correct adHoc fields")
		}
	})
	t.Run("Info method should create a log with LvlInfo and forward msg and adHocFields", func(t *testing.T) {
		msg := "some msg"
		adHocFields := []LogFields{
			{"a": "aaa", "b": "bbb", "c": "ccc"},
			{"d": "ddd", "e": "eee"},
			{"f": "fff", "g": "ggg", "h": "hhh"},
		}
		c := make(chan Log, 1)
		m := &mockAsyncScheduler{mockNextChannel: func() chan<- Log { return c }}
		l := &logger{configuration: &Configuration{AsyncScheduler: m, LvlsEnabled: LvlInfo}}
		l.Info(msg, adHocFields...)
		receivedLog := <-c
		if receivedLog.lvl != LvlInfo {
			t.Fatalf("Expected to create a log with the correct lvl")
		}
		if receivedLog.msg != msg {
			t.Fatalf("Expected to create a log with the correct msg")
		}
		if !reflect.DeepEqual(receivedLog.adHocFields, adHocFields) {
			t.Fatalf("Expected to create a log with the correct adHoc fields")
		}
	})
	t.Run("Warn method should create a log with LvlWarn and forward msg and adHocFields", func(t *testing.T) {
		msg := "some msg"
		adHocFields := []LogFields{
			{"a": "aaa", "b": "bbb", "c": "ccc"},
			{"d": "ddd", "e": "eee"},
			{"f": "fff", "g": "ggg", "h": "hhh"},
		}
		c := make(chan Log, 1)
		m := &mockAsyncScheduler{mockNextChannel: func() chan<- Log { return c }}
		l := &logger{configuration: &Configuration{AsyncScheduler: m, LvlsEnabled: LvlWarn}}
		l.Warn(msg, adHocFields...)
		receivedLog := <-c
		if receivedLog.lvl != LvlWarn {
			t.Fatalf("Expected to create a log with the correct lvl")
		}
		if receivedLog.msg != msg {
			t.Fatalf("Expected to create a log with the correct msg")
		}
		if !reflect.DeepEqual(receivedLog.adHocFields, adHocFields) {
			t.Fatalf("Expected to create a log with the correct adHoc fields")
		}
	})
	t.Run("Error method should create a log with LvlError and forward msg and adHocFields", func(t *testing.T) {
		msg := "some msg"
		adHocFields := []LogFields{
			{"a": "aaa", "b": "bbb", "c": "ccc"},
			{"d": "ddd", "e": "eee"},
			{"f": "fff", "g": "ggg", "h": "hhh"},
		}
		c := make(chan Log, 1)
		m := &mockAsyncScheduler{mockNextChannel: func() chan<- Log { return c }}
		l := &logger{configuration: &Configuration{AsyncScheduler: m, LvlsEnabled: LvlError}}
		l.Error(msg, adHocFields...)
		receivedLog := <-c
		if receivedLog.lvl != LvlError {
			t.Fatalf("Expected to create a log with the correct lvl")
		}
		if receivedLog.msg != msg {
			t.Fatalf("Expected to create a log with the correct msg")
		}
		if !reflect.DeepEqual(receivedLog.adHocFields, adHocFields) {
			t.Fatalf("Expected to create a log with the correct adHoc fields")
		}
	})
	t.Run("Fatal method should create a log with LvlFatal and forward msg and adHocFields", func(t *testing.T) {
		msg := "some msg"
		adHocFields := []LogFields{
			{"a": "aaa", "b": "bbb", "c": "ccc"},
			{"d": "ddd", "e": "eee"},
			{"f": "fff", "g": "ggg", "h": "hhh"},
		}
		c := make(chan Log, 1)
		m := &mockAsyncScheduler{mockNextChannel: func() chan<- Log { return c }}
		l := &logger{configuration: &Configuration{AsyncScheduler: m, LvlsEnabled: LvlFatal}}
		l.Fatal(msg, adHocFields...)
		receivedLog := <-c
		if receivedLog.lvl != LvlFatal {
			t.Fatalf("Expected to create a log with the correct lvl")
		}
		if receivedLog.msg != msg {
			t.Fatalf("Expected to create a log with the correct msg")
		}
		if !reflect.DeepEqual(receivedLog.adHocFields, adHocFields) {
			t.Fatalf("Expected to create a log with the correct adHoc fields")
		}
	})
	t.Run("ErrorFrom should create a log with LvlError, call ErrorParser and forward msg, error fields (inserted before the adHocs) and adHocFields", func(t *testing.T) {
		msg := "some error msg"
		adHocFields := []LogFields{
			{"a": "aaa", "b": "bbb", "c": "ccc"},
			{"d": "ddd", "e": "eee"},
			{"f": "fff", "g": "ggg", "h": "hhh"},
		}
		c := make(chan Log, 1)
		m := &mockAsyncScheduler{mockNextChannel: func() chan<- Log { return c }}
		e := fmt.Errorf("some random string")
		calls := 0
		l := &logger{configuration: &Configuration{AsyncScheduler: m, LvlsEnabled: LvlError, ErrorParser: func(err error) (string, LogFields) {
			calls += 1
			if e != err {
				t.Fatalf("Expected the correct error")
			}
			return msg, LogFields{"error": e}
		}}}
		l.ErrorFrom(e, adHocFields...)
		receivedLog := <-c
		if calls != 1 {
			t.Fatalf("Expected to call ErrorParser")
		}
		if receivedLog.lvl != LvlError {
			t.Fatalf("Expected to create a log with the correct lvl")
		}
		if receivedLog.msg != msg {
			t.Fatalf("Expected to create a log with the correct msg")
		}
		if !reflect.DeepEqual(receivedLog.adHocFields, append([]LogFields{{"error": e}}, adHocFields...)) {
			t.Fatalf("Expected to create a log with the correct adHoc fields")
		}
	})
	t.Run("ErrorFrom should create a log with LvlError, call ErrorParser and forward msg, error fields (inserted before the adHocs) and adHocFields", func(t *testing.T) {
		msg := "some error msg"
		adHocFields := []LogFields{
			{"a": "aaa", "b": "bbb", "c": "ccc"},
			{"d": "ddd", "e": "eee"},
			{"f": "fff", "g": "ggg", "h": "hhh"},
		}
		c := make(chan Log, 1)
		m := &mockAsyncScheduler{mockNextChannel: func() chan<- Log { return c }}
		e := fmt.Errorf("some random string")
		calls := 0
		l := &logger{configuration: &Configuration{AsyncScheduler: m, LvlsEnabled: LvlError, ErrorParser: func(err error) (string, LogFields) {
			calls += 1
			if e != err {
				t.Fatalf("Expected the correct error")
			}
			return msg, nil
		}}}
		l.ErrorFrom(e, adHocFields...)
		receivedLog := <-c
		if calls != 1 {
			t.Fatalf("Expected to call ErrorParser")
		}
		if receivedLog.lvl != LvlError {
			t.Fatalf("Expected to create a log with the correct lvl")
		}
		if receivedLog.msg != msg {
			t.Fatalf("Expected to create a log with the correct msg")
		}
		if !reflect.DeepEqual(receivedLog.adHocFields, adHocFields) {
			t.Fatalf("Expected to create a log with the correct adHoc fields")
		}
	})
	t.Run("FatalFrom should create a log with LvlFatal, call ErrorParser and forward msg, error fields (inserted before the adHocs) and adHocFields", func(t *testing.T) {
		msg := "some error msg"
		adHocFields := []LogFields{
			{"a": "aaa", "b": "bbb", "c": "ccc"},
			{"d": "ddd", "e": "eee"},
			{"f": "fff", "g": "ggg", "h": "hhh"},
		}
		c := make(chan Log, 1)
		m := &mockAsyncScheduler{mockNextChannel: func() chan<- Log { return c }}
		e := fmt.Errorf("some random string")
		calls := 0
		l := &logger{configuration: &Configuration{AsyncScheduler: m, LvlsEnabled: LvlFatal, ErrorParser: func(err error) (string, LogFields) {
			calls += 1
			if e != err {
				t.Fatalf("Expected the correct error")
			}
			return msg, LogFields{"error": e}
		}}}
		l.FatalFrom(e, adHocFields...)
		receivedLog := <-c
		if calls != 1 {
			t.Fatalf("Expected to call ErrorParser")
		}
		if receivedLog.lvl != LvlFatal {
			t.Fatalf("Expected to create a log with the correct lvl")
		}
		if receivedLog.msg != msg {
			t.Fatalf("Expected to create a log with the correct msg")
		}
		if !reflect.DeepEqual(receivedLog.adHocFields, append([]LogFields{{"error": e}}, adHocFields...)) {
			t.Fatalf("Expected to create a log with the correct adHoc fields")
		}
	})
	t.Run("FatalFrom should create a log with LvlFatal, call ErrorParser and forward msg, error fields (inserted before the adHocs) and adHocFields", func(t *testing.T) {
		msg := "some error msg"
		adHocFields := []LogFields{
			{"a": "aaa", "b": "bbb", "c": "ccc"},
			{"d": "ddd", "e": "eee"},
			{"f": "fff", "g": "ggg", "h": "hhh"},
		}
		c := make(chan Log, 1)
		m := &mockAsyncScheduler{mockNextChannel: func() chan<- Log { return c }}
		e := fmt.Errorf("some random string")
		calls := 0
		l := &logger{configuration: &Configuration{AsyncScheduler: m, LvlsEnabled: LvlFatal, ErrorParser: func(err error) (string, LogFields) {
			calls += 1
			if e != err {
				t.Fatalf("Expected the correct error")
			}
			return msg, nil
		}}}
		l.FatalFrom(e, adHocFields...)
		receivedLog := <-c
		if calls != 1 {
			t.Fatalf("Expected to call ErrorParser")
		}
		if receivedLog.lvl != LvlFatal {
			t.Fatalf("Expected to create a log with the correct lvl")
		}
		if receivedLog.msg != msg {
			t.Fatalf("Expected to create a log with the correct msg")
		}
		if !reflect.DeepEqual(receivedLog.adHocFields, adHocFields) {
			t.Fatalf("Expected to create a log with the correct adHoc fields")
		}
	})
}

type mockAsyncScheduler struct {
	mockNextChannel func() chan<- Log
	mockShutdown    func()
}

func (f *mockAsyncScheduler) NextChannel() chan<- Log {
	if f.mockNextChannel != nil {
		return f.mockNextChannel()
	}
	return nil
}

func (f *mockAsyncScheduler) Shutdown() {
	if f.mockShutdown != nil {
		f.mockShutdown()
	}
}
