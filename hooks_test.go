package logger

import (
	"reflect"
	"testing"
)

func TestPreHooks(t *testing.T) {
	t.Run("Should return a new equal instance of src", func(t *testing.T) {
		l := &logger{}
		newLog := l.PreHooks(nil).(*logger)
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(newLog).Pointer() {
			t.Fatalf("Expected a different instance")
		}
	})
	t.Run("The new instance should point to the same configuration", func(t *testing.T) {
		c := &Configuration{}
		l := &logger{configuration: c}
		newLog := l.PreHooks(nil).(*logger)
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
		l := &logger{
			fields:    LogFields{"a": "aaa", "b": "bbb", "c": "ccc"},
			preHooks:  Hooks{"d": fnD, "e": fnE},
			postHooks: Hooks{"f": fnF, "g": fnG},
			outputs:   []Output{outH, outI},
		}
		newLog := l.PreHooks(nil).(*logger)
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
	t.Run("Merge the given PreHooks with the src pre hooks, overriding duplicates", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		fnB := func(log Log) interface{} { return nil }
		fnC := func(log Log) interface{} { return nil }
		fnD := func(log Log) interface{} { return nil }
		fnE := func(log Log) interface{} { return nil }
		preHooks := Hooks{"b": fnD, "d": fnE}
		l := &logger{preHooks: Hooks{"a": fnA, "b": fnB, "c": fnC}}
		newLog := l.PreHooks(preHooks).(*logger)
		if reflect.ValueOf(newLog.preHooks["a"]).Pointer() != reflect.ValueOf(fnA).Pointer() ||
			reflect.ValueOf(newLog.preHooks["b"]).Pointer() != reflect.ValueOf(fnD).Pointer() ||
			reflect.ValueOf(newLog.preHooks["c"]).Pointer() != reflect.ValueOf(fnC).Pointer() ||
			reflect.ValueOf(newLog.preHooks["d"]).Pointer() != reflect.ValueOf(fnE).Pointer() {
			t.Fatalf("Expected the merge of the src preHooks + argument")
		}
	})
}

func TestRawPreHooks(t *testing.T) {
	t.Run("Should return a new equal instance of src", func(t *testing.T) {
		l := &logger{}
		newLog := l.RawPreHooks(nil).(*logger)
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(newLog).Pointer() {
			t.Fatalf("Expected a different instance")
		}
	})
	t.Run("The new instance should point to the same configuration", func(t *testing.T) {
		c := &Configuration{}
		l := &logger{configuration: c}
		newLog := l.RawPreHooks(nil).(*logger)
		if reflect.ValueOf(l.configuration).Pointer() != reflect.ValueOf(newLog.configuration).Pointer() {
			t.Fatalf("Expected the same configuration")
		}
	})
	t.Run("The new instance should be equivalent, but with new slices/maps (except the preHooks)", func(t *testing.T) {
		fnF := func(log Log) interface{} { return nil }
		fnG := func(log Log) interface{} { return nil }
		outH := func(lvl uint64, msg string, fields LogFields) {}
		outI := func(lvl uint64, msg string, fields LogFields) {}
		l := &logger{
			fields:    LogFields{"a": "aaa", "b": "bbb", "c": "ccc"},
			postHooks: Hooks{"f": fnF, "g": fnG},
			outputs:   []Output{outH, outI},
		}
		newLog := l.RawPreHooks(nil).(*logger)
		if reflect.ValueOf(l.fields).Pointer() == reflect.ValueOf(newLog.fields).Pointer() {
			t.Fatalf("Expected to be a different fields map")
		}
		if !reflect.DeepEqual(l.fields, newLog.fields) {
			t.Fatalf("Expected an equivalent fields map")
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
	t.Run("Should set the Logger preHooks, ignoring the previous one", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		fnB := func(log Log) interface{} { return nil }
		fnC := func(log Log) interface{} { return nil }
		fnD := func(log Log) interface{} { return nil }
		fnE := func(log Log) interface{} { return nil }
		preHooks := Hooks{"b": fnD, "d": fnE}
		l := &logger{preHooks: Hooks{"a": fnA, "b": fnB, "c": fnC}}
		newLog := l.RawPreHooks(preHooks).(*logger)
		if len(newLog.preHooks) != 2 ||
			reflect.ValueOf(newLog.preHooks["b"]).Pointer() != reflect.ValueOf(fnD).Pointer() ||
			reflect.ValueOf(newLog.preHooks["d"]).Pointer() != reflect.ValueOf(fnE).Pointer() {
			t.Fatalf("Expected the override the value of the src preHooks completelly")
		}
	})
}

func TestPostHooks(t *testing.T) {
	t.Run("Should return a new equal instance of src", func(t *testing.T) {
		l := &logger{}
		newLog := l.PostHooks(nil).(*logger)
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(newLog).Pointer() {
			t.Fatalf("Expected a different instance")
		}
	})
	t.Run("The new instance should point to the same configuration", func(t *testing.T) {
		c := &Configuration{}
		l := &logger{configuration: c}
		newLog := l.PostHooks(nil).(*logger)
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
		l := &logger{
			fields:    LogFields{"a": "aaa", "b": "bbb", "c": "ccc"},
			preHooks:  Hooks{"d": fnD, "e": fnE},
			postHooks: Hooks{"f": fnF, "g": fnG},
			outputs:   []Output{outH, outI},
		}
		newLog := l.PostHooks(nil).(*logger)
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
	t.Run("Merge the given PostHooks with the src post hooks, overriding duplicates", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		fnB := func(log Log) interface{} { return nil }
		fnC := func(log Log) interface{} { return nil }
		fnD := func(log Log) interface{} { return nil }
		fnE := func(log Log) interface{} { return nil }
		postHooks := Hooks{"b": fnD, "d": fnE}
		l := &logger{postHooks: Hooks{"a": fnA, "b": fnB, "c": fnC}}
		newLog := l.PostHooks(postHooks).(*logger)
		if reflect.ValueOf(newLog.postHooks["a"]).Pointer() != reflect.ValueOf(fnA).Pointer() ||
			reflect.ValueOf(newLog.postHooks["b"]).Pointer() != reflect.ValueOf(fnD).Pointer() ||
			reflect.ValueOf(newLog.postHooks["c"]).Pointer() != reflect.ValueOf(fnC).Pointer() ||
			reflect.ValueOf(newLog.postHooks["d"]).Pointer() != reflect.ValueOf(fnE).Pointer() {
			t.Fatalf("Expected the merge of the src postHooks + argument")
		}
	})
}

func TestRawPostHooks(t *testing.T) {
	t.Run("Should return a new equal instance of src", func(t *testing.T) {
		l := &logger{}
		newLog := l.RawPostHooks(nil).(*logger)
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(newLog).Pointer() {
			t.Fatalf("Expected a different instance")
		}
	})
	t.Run("The new instance should point to the same configuration", func(t *testing.T) {
		c := &Configuration{}
		l := &logger{configuration: c}
		newLog := l.RawPostHooks(nil).(*logger)
		if reflect.ValueOf(l.configuration).Pointer() != reflect.ValueOf(newLog.configuration).Pointer() {
			t.Fatalf("Expected the same configuration")
		}
	})
	t.Run("The new instance should be equivalent, but with new slices/maps (except the postHooks)", func(t *testing.T) {
		fnF := func(log Log) interface{} { return nil }
		fnG := func(log Log) interface{} { return nil }
		outH := func(lvl uint64, msg string, fields LogFields) {}
		outI := func(lvl uint64, msg string, fields LogFields) {}
		l := &logger{
			fields:   LogFields{"a": "aaa", "b": "bbb", "c": "ccc"},
			preHooks: Hooks{"f": fnF, "g": fnG},
			outputs:  []Output{outH, outI},
		}
		newLog := l.RawPostHooks(nil).(*logger)
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
	t.Run("Should set the Logger postHooks, ignoring the previous one", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		fnB := func(log Log) interface{} { return nil }
		fnC := func(log Log) interface{} { return nil }
		fnD := func(log Log) interface{} { return nil }
		fnE := func(log Log) interface{} { return nil }
		postHooks := Hooks{"b": fnD, "d": fnE}
		l := &logger{postHooks: Hooks{"a": fnA, "b": fnB, "c": fnC}}
		newLog := l.RawPostHooks(postHooks).(*logger)
		if len(newLog.postHooks) != 2 ||
			reflect.ValueOf(newLog.postHooks["b"]).Pointer() != reflect.ValueOf(fnD).Pointer() ||
			reflect.ValueOf(newLog.postHooks["d"]).Pointer() != reflect.ValueOf(fnE).Pointer() {
			t.Fatalf("Expected the override the value of the src postHooks completelly")
		}
	})
}
