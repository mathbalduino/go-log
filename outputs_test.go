package loxeLog

import (
	"reflect"
	"testing"
)

func TestLogger_Outputs(t *testing.T) {
	t.Run("Should return a new equal instance of logger", func(t *testing.T) {
		l := &Logger{}
		newLog := l.Outputs(nil)
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(newLog).Pointer() {
			t.Fatalf("Expected a different instance")
		}
	})
	t.Run("The new instance should point to the same configuration", func(t *testing.T) {
		c := &Configuration{}
		l := &Logger{configuration: c}
		newLog := l.Outputs(nil)
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
		newLog := l.Outputs(nil)
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
	t.Run("Should ignore the nil outputs", func(t *testing.T) {
		fnA := func(lvl uint64, msg string, fields LogFields) {}
		fnB := func(lvl uint64, msg string, fields LogFields) {}
		fnC := func(lvl uint64, msg string, fields LogFields) {}
		fnD := func(lvl uint64, msg string, fields LogFields) {}
		fnE := func(lvl uint64, msg string, fields LogFields) {}
		l := &Logger{outputs: []Output{fnA, fnB, fnC}}
		newLog := l.Outputs(nil, nil, fnD, nil, nil, fnE)
		if len(newLog.outputs) != 5 ||
			reflect.ValueOf(newLog.outputs[0]).Pointer() != reflect.ValueOf(fnA).Pointer() ||
			reflect.ValueOf(newLog.outputs[1]).Pointer() != reflect.ValueOf(fnB).Pointer() ||
			reflect.ValueOf(newLog.outputs[2]).Pointer() != reflect.ValueOf(fnC).Pointer() ||
			reflect.ValueOf(newLog.outputs[3]).Pointer() != reflect.ValueOf(fnD).Pointer() ||
			reflect.ValueOf(newLog.outputs[4]).Pointer() != reflect.ValueOf(fnE).Pointer() {
			t.Fatalf("Expected the merge of the logger outputs + argument")
		}
	})
	t.Run("Merge the given Outputs with the logger outputs", func(t *testing.T) {
		fnA := func(lvl uint64, msg string, fields LogFields) {}
		fnB := func(lvl uint64, msg string, fields LogFields) {}
		fnC := func(lvl uint64, msg string, fields LogFields) {}
		fnD := func(lvl uint64, msg string, fields LogFields) {}
		fnE := func(lvl uint64, msg string, fields LogFields) {}
		l := &Logger{outputs: []Output{fnA, fnB, fnC}}
		newLog := l.Outputs(fnD, fnE)
		if len(newLog.outputs) != 5 ||
			reflect.ValueOf(newLog.outputs[0]).Pointer() != reflect.ValueOf(fnA).Pointer() ||
			reflect.ValueOf(newLog.outputs[1]).Pointer() != reflect.ValueOf(fnB).Pointer() ||
			reflect.ValueOf(newLog.outputs[2]).Pointer() != reflect.ValueOf(fnC).Pointer() ||
			reflect.ValueOf(newLog.outputs[3]).Pointer() != reflect.ValueOf(fnD).Pointer() ||
			reflect.ValueOf(newLog.outputs[4]).Pointer() != reflect.ValueOf(fnE).Pointer() {
			t.Fatalf("Expected the merge of the logger outputs + argument")
		}
	})
}

func TestLogger_RawOutputs(t *testing.T) {
	t.Run("Should return a new equal instance of logger", func(t *testing.T) {
		l := &Logger{}
		newLog := l.RawOutputs(nil)
		if reflect.ValueOf(l).Pointer() == reflect.ValueOf(newLog).Pointer() {
			t.Fatalf("Expected a different instance")
		}
	})
	t.Run("The new instance should point to the same configuration", func(t *testing.T) {
		c := &Configuration{}
		l := &Logger{configuration: c}
		newLog := l.RawOutputs(nil)
		if reflect.ValueOf(l.configuration).Pointer() != reflect.ValueOf(newLog.configuration).Pointer() {
			t.Fatalf("Expected the same configuration")
		}
	})
	t.Run("The new instance should be equivalent, but with new slices/maps (except outputs)", func(t *testing.T) {
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
		newLog := l.RawOutputs(nil)
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
	})
	t.Run("If it's given only nil outputs, set an empty outputs slice", func(t *testing.T) {
		fnA := func(lvl uint64, msg string, fields LogFields) {}
		fnB := func(lvl uint64, msg string, fields LogFields) {}
		fnC := func(lvl uint64, msg string, fields LogFields) {}
		l := &Logger{outputs: []Output{fnA, fnB, fnC}}
		newLog := l.RawOutputs(nil, nil, nil, nil)
		if len(newLog.outputs) != 0 {
			t.Fatalf("Expected set the logger outputs to an empty slice")
		}
	})
	t.Run("Should ignore the nil outputs", func(t *testing.T) {
		fnA := func(lvl uint64, msg string, fields LogFields) {}
		fnB := func(lvl uint64, msg string, fields LogFields) {}
		fnC := func(lvl uint64, msg string, fields LogFields) {}
		fnD := func(lvl uint64, msg string, fields LogFields) {}
		fnE := func(lvl uint64, msg string, fields LogFields) {}
		l := &Logger{outputs: []Output{fnA, fnB, fnC}}
		newLog := l.RawOutputs(nil, nil, fnD, nil, nil, fnE)
		if len(newLog.outputs) != 2 ||
			reflect.ValueOf(newLog.outputs[0]).Pointer() != reflect.ValueOf(fnD).Pointer() ||
			reflect.ValueOf(newLog.outputs[1]).Pointer() != reflect.ValueOf(fnE).Pointer() {
			t.Fatalf("Expected set the logger outputs to the given argument")
		}
	})
	t.Run("Set the given Outputs to the logger outputs", func(t *testing.T) {
		fnA := func(lvl uint64, msg string, fields LogFields) {}
		fnB := func(lvl uint64, msg string, fields LogFields) {}
		fnC := func(lvl uint64, msg string, fields LogFields) {}
		fnD := func(lvl uint64, msg string, fields LogFields) {}
		fnE := func(lvl uint64, msg string, fields LogFields) {}
		l := &Logger{outputs: []Output{fnA, fnB, fnC}}
		newLog := l.RawOutputs(fnD, fnE)
		if len(newLog.outputs) != 2 ||
			reflect.ValueOf(newLog.outputs[0]).Pointer() != reflect.ValueOf(fnD).Pointer() ||
			reflect.ValueOf(newLog.outputs[1]).Pointer() != reflect.ValueOf(fnE).Pointer() {
			t.Fatalf("Expected to set the logger outputs to the given arguments")
		}
	})
}
