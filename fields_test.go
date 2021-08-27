package loxeLog

import (
	"reflect"
	"testing"
)

func TestFields(t *testing.T) {
	t.Run("Should return a new equal instance of logger", func(t *testing.T) {
		l := &Logger{}
		newLog := l.Fields(nil)
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(newLog).Pointer() {
			t.Fatalf("Expected a different instance")
		}
	})
	t.Run("The new instance should point to the same configuration", func(t *testing.T) {
		c := &Configuration{}
		l := &Logger{configuration: c}
		newLog := l.Fields(nil)
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
		newLog := l.Fields(nil)
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
	t.Run("Merge the given LogFields with the logger fields, overriding duplicates", func(t *testing.T) {
		fields := LogFields{"a": "AAA", "d": "ddd", "e": "eee"}
		l := &Logger{fields: LogFields{"a": "aaa", "b": "bbb", "c": "ccc"}}
		newLog := l.Fields(fields)
		if newLog.fields["a"] != "AAA" || newLog.fields["b"] != "bbb" || newLog.fields["c"] != "ccc" ||
			newLog.fields["d"] != "ddd" || newLog.fields["e"] != "eee" {
			t.Fatalf("Expected the merge of the logger fields + argument")
		}
	})
}

func TestRawFields(t *testing.T) {
	t.Run("Should return a new equal instance of logger", func(t *testing.T) {
		l := &Logger{}
		newLog := l.RawFields(nil)
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(newLog).Pointer() {
			t.Fatalf("Expected a different instance")
		}
	})
	t.Run("The new instance should point to the same configuration", func(t *testing.T) {
		c := &Configuration{}
		l := &Logger{configuration: c}
		newLog := l.RawFields(nil)
		if reflect.ValueOf(l.configuration).Pointer() != reflect.ValueOf(newLog.configuration).Pointer() {
			t.Fatalf("Expected the same configuration")
		}
	})
	t.Run("The new instance should be equivalent, but with new slices/maps (except the fields)", func(t *testing.T) {
		fnD := func(log Log) interface{} { return nil }
		fnE := func(log Log) interface{} { return nil }
		fnF := func(log Log) interface{} { return nil }
		fnG := func(log Log) interface{} { return nil }
		outH := func(lvl uint64, msg string, fields LogFields) {}
		outI := func(lvl uint64, msg string, fields LogFields) {}
		l := &Logger{
			syncHooks:  Hooks{"d": fnD, "e": fnE},
			asyncHooks: Hooks{"f": fnF, "g": fnG},
			outputs:    []Output{outH, outI},
		}
		newLog := l.RawFields(nil)
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
	t.Run("Should set the Logger fields, ignoring the previous one", func(t *testing.T) {
		fields := LogFields{"a": "AAA", "d": "ddd", "e": "eee"}
		l := &Logger{fields: LogFields{"a": "aaa", "b": "bbb", "c": "ccc"}}
		newLog := l.RawFields(fields)
		if len(newLog.fields) != 3 || newLog.fields["a"] != "AAA" || newLog.fields["d"] != "ddd" || newLog.fields["e"] != "eee" {
			t.Fatalf("Expected to override the value of the logger fields completelly")
		}
	})
}
