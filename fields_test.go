package logger

import (
	"reflect"
	"testing"
)

func TestField(t *testing.T) {
	t.Run("Should return nil if the given key doesn't exists", func(t *testing.T) {
		l := &Logger{fields: LogFields{"a": "aaa", "b": "bbb", "c": "ccc"}}
		v := l.Field("d")
		if v != nil {
			t.Fatalf("Expected to be nil")
		}
		v = l.Field("zzz")
		if v != nil {
			t.Fatalf("Expected to be nil")
		}
	})
	t.Run("Should return the value of the given key", func(t *testing.T) {
		l := &Logger{fields: LogFields{"a": "aaa", "b": "bbb", "c": "ccc"}}
		v := l.Field("a")
		if v == nil || v.(string) != "aaa" {
			t.Fatalf("Expected to be the correct value")
		}
		v = l.Field("c")
		if v == nil || v.(string) != "ccc" {
			t.Fatalf("Expected to be the correct value")
		}
	})
}

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
			fields:    LogFields{"a": "aaa", "b": "bbb", "c": "ccc"},
			preHooks:  Hooks{"d": fnD, "e": fnE},
			postHooks: Hooks{"f": fnF, "g": fnG},
			outputs:   []Output{outH, outI},
		}
		newLog := l.Fields(nil)
		if reflect.ValueOf(l.fields).Pointer() == reflect.ValueOf(newLog.fields).Pointer() {
			t.Fatalf("Expected to be a different fields map")
		}
		if !reflect.DeepEqual(l.fields, newLog.fields) {
			t.Fatalf("Expected an equivalent fields map")
		}
		if reflect.ValueOf(l.preHooks).Pointer() == reflect.ValueOf(newLog.preHooks).Pointer() {
			t.Fatalf("Expected to be a different preHooks map")
		}
		if len(l.preHooks) != len(newLog.preHooks) {
			t.Fatalf("Expected an equivalent preHooks map")
		}
		for k, v := range l.preHooks {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.preHooks[k]).Pointer() {
				t.Fatalf("Expected an equivalent preHooks map")
			}
		}
		if reflect.ValueOf(l.postHooks).Pointer() == reflect.ValueOf(newLog.postHooks).Pointer() {
			t.Fatalf("Expected to be a different postHooks map")
		}
		if len(l.postHooks) != len(newLog.postHooks) {
			t.Fatalf("Expected an equivalent postHooks map")
		}
		for k, v := range l.postHooks {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.postHooks[k]).Pointer() {
				t.Fatalf("Expected an equivalent postHooks map")
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
			preHooks:  Hooks{"d": fnD, "e": fnE},
			postHooks: Hooks{"f": fnF, "g": fnG},
			outputs:   []Output{outH, outI},
		}
		newLog := l.RawFields(nil)
		if reflect.ValueOf(l.preHooks).Pointer() == reflect.ValueOf(newLog.preHooks).Pointer() {
			t.Fatalf("Expected to be a different preHooks map")
		}
		if len(l.preHooks) != len(newLog.preHooks) {
			t.Fatalf("Expected an equivalent preHooks map")
		}
		for k, v := range l.preHooks {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.preHooks[k]).Pointer() {
				t.Fatalf("Expected an equivalent preHooks map")
			}
		}
		if reflect.ValueOf(l.postHooks).Pointer() == reflect.ValueOf(newLog.postHooks).Pointer() {
			t.Fatalf("Expected to be a different postHooks map")
		}
		if len(l.postHooks) != len(newLog.postHooks) {
			t.Fatalf("Expected an equivalent postHooks map")
		}
		for k, v := range l.postHooks {
			if reflect.ValueOf(v).Pointer() != reflect.ValueOf(newLog.postHooks[k]).Pointer() {
				t.Fatalf("Expected an equivalent postHooks map")
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
