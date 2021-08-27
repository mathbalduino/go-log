package loxeLog

import (
	"reflect"
	"testing"
)

func TestSyncHooks(t *testing.T) {
	t.Run("Should return a new equal instance of logger", func(t *testing.T) {
		l := &Logger{}
		newLog := l.SyncHooks(nil)
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(newLog).Pointer() {
			t.Fatalf("Expected a different instance")
		}
	})
	t.Run("The new instance should point to the same configuration", func(t *testing.T) {
		c := &Configuration{}
		l := &Logger{configuration: c}
		newLog := l.SyncHooks(nil)
		if reflect.ValueOf(l.configuration).Pointer() != reflect.ValueOf(newLog.configuration).Pointer() {
			t.Fatalf("Expected the same configuration")
		}
	})
	t.Run("The new instance should be equivalent, but with new slices/maps", func(t *testing.T) {
		fnD := func(log Log) interface{} { return nil }
		fnE := func(log Log) interface{} { return nil }
		fnF := func(log Log) interface{} { return nil }
		fnG := func(log Log) interface{} { return nil }
		outH := func(lvl uint64, msg string, fields LogFields) {}
		outI := func(lvl uint64, msg string, fields LogFields) {}
		l := &Logger{
			fields:     LogFields{"a": "aaa", "b": "bbb", "c": "ccc"},
			syncHooks:  Hooks{"d": fnD, "e": fnE},
			asyncHooks: Hooks{"f": fnF, "g": fnG},
			outputs:    []Output{outH, outI},
		}
		newLog := l.SyncHooks(nil)
		if reflect.ValueOf(l.fields).Pointer() == reflect.ValueOf(newLog.fields).Pointer() {
			t.Fatalf("Expected to be a different fields map")
		}
		if !reflect.DeepEqual(l.fields, newLog.fields) {
			t.Fatalf("Expected an equivalent fields map")
		}
		if reflect.ValueOf(l.syncHooks).Pointer() == reflect.ValueOf(newLog.syncHooks).Pointer() {
			t.Fatalf("Expected to be a different syncHooks map")
		}
		if len(l.syncHooks) != len(newLog.syncHooks) {
			t.Fatalf("Expected an equivalent syncHooks map")
		}
		for k, v := range l.syncHooks {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.syncHooks[k]).Pointer() {
				t.Fatalf("Expected an equivalent syncHooks map")
			}
		}
		if reflect.ValueOf(l.asyncHooks).Pointer() == reflect.ValueOf(newLog.asyncHooks).Pointer() {
			t.Fatalf("Expected to be a different asyncHooks map")
		}
		if len(l.asyncHooks) != len(newLog.asyncHooks) {
			t.Fatalf("Expected an equivalent asyncHooks map")
		}
		for k, v := range l.asyncHooks {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.asyncHooks[k]).Pointer() {
				t.Fatalf("Expected an equivalent asyncHooks map")
			}
		}
		if reflect.ValueOf(l.outputs).Pointer() == reflect.ValueOf(newLog.outputs).Pointer() {
			t.Fatalf("Expected to be a different outputs slice")
		}
		if len(l.outputs) != len(newLog.outputs) {
			t.Fatalf("Expected an equivalent outputs slice")
		}
		for k, v := range l.outputs {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.outputs[k]).Pointer() {
				t.Fatalf("Expected an equivalent outputs slice")
			}
		}
	})
	t.Run("Merge the given SyncHooks with the logger sync hooks, overriding duplicates", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		fnB := func(log Log) interface{} { return nil }
		fnC := func(log Log) interface{} { return nil }
		fnD := func(log Log) interface{} { return nil }
		fnE := func(log Log) interface{} { return nil }
		syncHooks := Hooks{"b": fnD, "d": fnE}
		l := &Logger{syncHooks: Hooks{"a": fnA, "b": fnB, "c": fnC}}
		newLog := l.SyncHooks(syncHooks)
		if reflect.ValueOf(newLog.syncHooks["a"]).Pointer() != reflect.ValueOf(fnA).Pointer() ||
			reflect.ValueOf(newLog.syncHooks["b"]).Pointer() != reflect.ValueOf(fnD).Pointer() ||
			reflect.ValueOf(newLog.syncHooks["c"]).Pointer() != reflect.ValueOf(fnC).Pointer() ||
			reflect.ValueOf(newLog.syncHooks["d"]).Pointer() != reflect.ValueOf(fnE).Pointer() {
			t.Fatalf("Expected the merge of the logger syncHooks + argument")
		}
	})
}

func TestRawSyncHooks(t *testing.T) {
	t.Run("Should return a new equal instance of logger", func(t *testing.T) {
		l := &Logger{}
		newLog := l.RawSyncHooks(nil)
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(newLog).Pointer() {
			t.Fatalf("Expected a different instance")
		}
	})
	t.Run("The new instance should point to the same configuration", func(t *testing.T) {
		c := &Configuration{}
		l := &Logger{configuration: c}
		newLog := l.RawSyncHooks(nil)
		if reflect.ValueOf(l.configuration).Pointer() != reflect.ValueOf(newLog.configuration).Pointer() {
			t.Fatalf("Expected the same configuration")
		}
	})
	t.Run("The new instance should be equivalent, but with new slices/maps (except the syncHooks)", func(t *testing.T) {
		fnF := func(log Log) interface{} { return nil }
		fnG := func(log Log) interface{} { return nil }
		outH := func(lvl uint64, msg string, fields LogFields) {}
		outI := func(lvl uint64, msg string, fields LogFields) {}
		l := &Logger{
			fields:     LogFields{"a": "aaa", "b": "bbb", "c": "ccc"},
			asyncHooks: Hooks{"f": fnF, "g": fnG},
			outputs:    []Output{outH, outI},
		}
		newLog := l.RawSyncHooks(nil)
		if reflect.ValueOf(l.fields).Pointer() == reflect.ValueOf(newLog.fields).Pointer() {
			t.Fatalf("Expected to be a different fields map")
		}
		if !reflect.DeepEqual(l.fields, newLog.fields) {
			t.Fatalf("Expected an equivalent fields map")
		}
		if reflect.ValueOf(l.asyncHooks).Pointer() == reflect.ValueOf(newLog.asyncHooks).Pointer() {
			t.Fatalf("Expected to be a different asyncHooks map")
		}
		if len(l.asyncHooks) != len(newLog.asyncHooks) {
			t.Fatalf("Expected an equivalent asyncHooks map")
		}
		for k, v := range l.asyncHooks {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.asyncHooks[k]).Pointer() {
				t.Fatalf("Expected an equivalent asyncHooks map")
			}
		}
		if reflect.ValueOf(l.outputs).Pointer() == reflect.ValueOf(newLog.outputs).Pointer() {
			t.Fatalf("Expected to be a different outputs slice")
		}
		if len(l.outputs) != len(newLog.outputs) {
			t.Fatalf("Expected an equivalent outputs slice")
		}
		for k, v := range l.outputs {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.outputs[k]).Pointer() {
				t.Fatalf("Expected an equivalent outputs slice")
			}
		}
	})
	t.Run("Should set the Logger syncHooks, ignoring the previous one", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		fnB := func(log Log) interface{} { return nil }
		fnC := func(log Log) interface{} { return nil }
		fnD := func(log Log) interface{} { return nil }
		fnE := func(log Log) interface{} { return nil }
		syncHooks := Hooks{"b": fnD, "d": fnE}
		l := &Logger{syncHooks: Hooks{"a": fnA, "b": fnB, "c": fnC}}
		newLog := l.RawSyncHooks(syncHooks)
		if len(newLog.syncHooks) != 2 ||
			reflect.ValueOf(newLog.syncHooks["b"]).Pointer() != reflect.ValueOf(fnD).Pointer() ||
			reflect.ValueOf(newLog.syncHooks["d"]).Pointer() != reflect.ValueOf(fnE).Pointer() {
			t.Fatalf("Expected the override the value of the logger syncHooks completelly")
		}
	})
}

func TestAsyncHooks(t *testing.T) {
	t.Run("Should return a new equal instance of logger", func(t *testing.T) {
		l := &Logger{}
		newLog := l.AsyncHooks(nil)
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(newLog).Pointer() {
			t.Fatalf("Expected a different instance")
		}
	})
	t.Run("The new instance should point to the same configuration", func(t *testing.T) {
		c := &Configuration{}
		l := &Logger{configuration: c}
		newLog := l.AsyncHooks(nil)
		if reflect.ValueOf(l.configuration).Pointer() != reflect.ValueOf(newLog.configuration).Pointer() {
			t.Fatalf("Expected the same configuration")
		}
	})
	t.Run("The new instance should be equivalent, but with new slices/maps", func(t *testing.T) {
		fnD := func(log Log) interface{} { return nil }
		fnE := func(log Log) interface{} { return nil }
		fnF := func(log Log) interface{} { return nil }
		fnG := func(log Log) interface{} { return nil }
		outH := func(lvl uint64, msg string, fields LogFields) {}
		outI := func(lvl uint64, msg string, fields LogFields) {}
		l := &Logger{
			fields:     LogFields{"a": "aaa", "b": "bbb", "c": "ccc"},
			syncHooks:  Hooks{"d": fnD, "e": fnE},
			asyncHooks: Hooks{"f": fnF, "g": fnG},
			outputs:    []Output{outH, outI},
		}
		newLog := l.AsyncHooks(nil)
		if reflect.ValueOf(l.fields).Pointer() == reflect.ValueOf(newLog.fields).Pointer() {
			t.Fatalf("Expected to be a different fields map")
		}
		if !reflect.DeepEqual(l.fields, newLog.fields) {
			t.Fatalf("Expected an equivalent fields map")
		}
		if reflect.ValueOf(l.syncHooks).Pointer() == reflect.ValueOf(newLog.syncHooks).Pointer() {
			t.Fatalf("Expected to be a different syncHooks map")
		}
		if len(l.syncHooks) != len(newLog.syncHooks) {
			t.Fatalf("Expected an equivalent syncHooks map")
		}
		for k, v := range l.syncHooks {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.syncHooks[k]).Pointer() {
				t.Fatalf("Expected an equivalent syncHooks map")
			}
		}
		if reflect.ValueOf(l.asyncHooks).Pointer() == reflect.ValueOf(newLog.asyncHooks).Pointer() {
			t.Fatalf("Expected to be a different asyncHooks map")
		}
		if len(l.asyncHooks) != len(newLog.asyncHooks) {
			t.Fatalf("Expected an equivalent asyncHooks map")
		}
		for k, v := range l.asyncHooks {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.asyncHooks[k]).Pointer() {
				t.Fatalf("Expected an equivalent asyncHooks map")
			}
		}
		if reflect.ValueOf(l.outputs).Pointer() == reflect.ValueOf(newLog.outputs).Pointer() {
			t.Fatalf("Expected to be a different outputs slice")
		}
		if len(l.outputs) != len(newLog.outputs) {
			t.Fatalf("Expected an equivalent outputs slice")
		}
		for k, v := range l.outputs {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.outputs[k]).Pointer() {
				t.Fatalf("Expected an equivalent outputs slice")
			}
		}
	})
	t.Run("Merge the given AsyncHooks with the logger async hooks, overriding duplicates", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		fnB := func(log Log) interface{} { return nil }
		fnC := func(log Log) interface{} { return nil }
		fnD := func(log Log) interface{} { return nil }
		fnE := func(log Log) interface{} { return nil }
		asyncHooks := Hooks{"b": fnD, "d": fnE}
		l := &Logger{asyncHooks: Hooks{"a": fnA, "b": fnB, "c": fnC}}
		newLog := l.AsyncHooks(asyncHooks)
		if reflect.ValueOf(newLog.asyncHooks["a"]).Pointer() != reflect.ValueOf(fnA).Pointer() ||
			reflect.ValueOf(newLog.asyncHooks["b"]).Pointer() != reflect.ValueOf(fnD).Pointer() ||
			reflect.ValueOf(newLog.asyncHooks["c"]).Pointer() != reflect.ValueOf(fnC).Pointer() ||
			reflect.ValueOf(newLog.asyncHooks["d"]).Pointer() != reflect.ValueOf(fnE).Pointer() {
			t.Fatalf("Expected the merge of the logger asyncHooks + argument")
		}
	})
}

func TestRawAsyncHooks(t *testing.T) {
	t.Run("Should return a new equal instance of logger", func(t *testing.T) {
		l := &Logger{}
		newLog := l.RawAsyncHooks(nil)
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(newLog).Pointer() {
			t.Fatalf("Expected a different instance")
		}
	})
	t.Run("The new instance should point to the same configuration", func(t *testing.T) {
		c := &Configuration{}
		l := &Logger{configuration: c}
		newLog := l.RawAsyncHooks(nil)
		if reflect.ValueOf(l.configuration).Pointer() != reflect.ValueOf(newLog.configuration).Pointer() {
			t.Fatalf("Expected the same configuration")
		}
	})
	t.Run("The new instance should be equivalent, but with new slices/maps (except the asyncHooks)", func(t *testing.T) {
		fnF := func(log Log) interface{} { return nil }
		fnG := func(log Log) interface{} { return nil }
		outH := func(lvl uint64, msg string, fields LogFields) {}
		outI := func(lvl uint64, msg string, fields LogFields) {}
		l := &Logger{
			fields:     LogFields{"a": "aaa", "b": "bbb", "c": "ccc"},
			syncHooks: Hooks{"f": fnF, "g": fnG},
			outputs:    []Output{outH, outI},
		}
		newLog := l.RawAsyncHooks(nil)
		if reflect.ValueOf(l.fields).Pointer() == reflect.ValueOf(newLog.fields).Pointer() {
			t.Fatalf("Expected to be a different fields map")
		}
		if !reflect.DeepEqual(l.fields, newLog.fields) {
			t.Fatalf("Expected an equivalent fields map")
		}
		if reflect.ValueOf(l.syncHooks).Pointer() == reflect.ValueOf(newLog.syncHooks).Pointer() {
			t.Fatalf("Expected to be a different syncHooks map")
		}
		if len(l.syncHooks) != len(newLog.syncHooks) {
			t.Fatalf("Expected an equivalent syncHooks map")
		}
		for k, v := range l.syncHooks {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.syncHooks[k]).Pointer() {
				t.Fatalf("Expected an equivalent syncHooks map")
			}
		}
		if reflect.ValueOf(l.outputs).Pointer() == reflect.ValueOf(newLog.outputs).Pointer() {
			t.Fatalf("Expected to be a different outputs slice")
		}
		if len(l.outputs) != len(newLog.outputs) {
			t.Fatalf("Expected an equivalent outputs slice")
		}
		for k, v := range l.outputs {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.outputs[k]).Pointer() {
				t.Fatalf("Expected an equivalent outputs slice")
			}
		}
	})
	t.Run("Should set the Logger asyncHooks, ignoring the previous one", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		fnB := func(log Log) interface{} { return nil }
		fnC := func(log Log) interface{} { return nil }
		fnD := func(log Log) interface{} { return nil }
		fnE := func(log Log) interface{} { return nil }
		asyncHooks := Hooks{"b": fnD, "d": fnE}
		l := &Logger{asyncHooks: Hooks{"a": fnA, "b": fnB, "c": fnC}}
		newLog := l.RawAsyncHooks(asyncHooks)
		if len(newLog.asyncHooks) != 2 ||
			reflect.ValueOf(newLog.asyncHooks["b"]).Pointer() != reflect.ValueOf(fnD).Pointer() ||
			reflect.ValueOf(newLog.asyncHooks["d"]).Pointer() != reflect.ValueOf(fnE).Pointer() {
			t.Fatalf("Expected the override the value of the logger asyncHooks completelly")
		}
	})
}
