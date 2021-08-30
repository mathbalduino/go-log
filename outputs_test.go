package loxeLog

import (
	"fmt"
	"io/ioutil"
	"os"
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

func TestOutputToWriter(t *testing.T) {
	t.Run("Should never return a nil Output", func(t *testing.T) {
		o := OutputToWriter(nil, nil, nil)
		if o == nil {
			t.Fatalf("Not expected to return a nil output")
		}
	})
	t.Run("Should call the given parser, forwarding the output fields", func(t *testing.T) {
		expectedFields := LogFields{"a": "aaa", "b": "bbb", "c": "ccc"}
		calls := 0
		parser := func(fields LogFields) ([]byte, error) {
			calls += 1
			if reflect.ValueOf(fields).Pointer() != reflect.ValueOf(expectedFields).Pointer() {
				t.Fatalf("Wrong fields given to the parser")
			}
			return nil, nil
		}
		o := OutputToWriter(&mockWriter{}, parser, nil)
		o(0, "", expectedFields)
		if calls != 1 {
			t.Fatalf("Expected to call the parser one time")
		}
	})
	t.Run("Should call the given onError when the parser returns any error, not calling Write()", func(t *testing.T) {
		parserCalls := 0
		err := fmt.Errorf("some error")
		parser := func(fields LogFields) ([]byte, error) {
			parserCalls += 1
			return nil, err
		}
		writerCalls := 0
		writer := &mockWriter{func(data []byte) (int, error) {
			writerCalls += 1
			return 0, nil
		}}
		onErrorCalls := 0
		onError := func(receivedErr error) {
			onErrorCalls += 1
			if receivedErr != err {
				t.Fatalf("Wrong error given to onError")
			}
		}
		o := OutputToWriter(writer, parser, onError)
		o(0, "", LogFields{"a": "aaa", "b": "bbb", "c": "ccc"})
		if parserCalls != 1 {
			t.Fatalf("Expected to call the parser one time")
		}
		if writerCalls != 0 {
			t.Fatalf("Expected to not call the writer.Write()")
		}
		if onErrorCalls != 1 {
			t.Fatalf("Expected to call onError")
		}
	})
	t.Run("Should call the writer.Write() forwarding the returned data from the parser", func(t *testing.T) {
		data := []byte{10, 11, 12, 13, 14}
		parserCalls := 0
		parser := func(fields LogFields) ([]byte, error) {
			parserCalls += 1
			return data, nil
		}
		writerCalls := 0
		writer := &mockWriter{func(receivedData []byte) (int, error) {
			writerCalls += 1
			if reflect.ValueOf(receivedData).Pointer() != reflect.ValueOf(data).Pointer() {
				t.Fatalf("Wrong data given to the writer.Write()")
			}
			return 0, nil
		}}
		o := OutputToWriter(writer, parser, nil)
		o(0, "", LogFields{"a": "aaa", "b": "bbb", "c": "ccc"})
		if parserCalls != 1 {
			t.Fatalf("Expected to call the parser one time")
		}
		if writerCalls != 1 {
			t.Fatalf("Expected to call the writer.Write()")
		}
	})
	t.Run("Should call the given onError forwarding any error returned by the writer.Write()", func(t *testing.T) {
		writerCalls := 0
		err := fmt.Errorf("some error")
		writer := &mockWriter{func(receivedData []byte) (int, error) {
			writerCalls += 1
			return 0, err
		}}
		onErrorCalls := 0
		onError := func(receivedErr error) {
			onErrorCalls += 1
			if receivedErr != err {
				t.Fatalf("Wrong error given to onError")
			}
		}
		parser := func(fields LogFields) ([]byte, error) { return []byte{1}, nil }
		o := OutputToWriter(writer, parser, onError)
		o(0, "", LogFields{"a": "aaa", "b": "bbb", "c": "ccc"})
		if writerCalls != 1 {
			t.Fatalf("Expected to call the writer.Write()")
		}
		if onErrorCalls != 1 {
			t.Fatalf("Expected to call onError")
		}
	})
}

func TestOutputJsonToFile(t *testing.T) {
	t.Run("Should never return a nil Output", func(t *testing.T) {
		o := OutputJsonToWriter(nil, nil)
		if o == nil {
			t.Fatalf("Not expected to return a nil output")
		}
	})
	t.Run("Should call onError if there's errors with the json.Marshal", func(t *testing.T) {
		onErrorCalls := 0
		onError := func(receivedErr error) {onErrorCalls += 1		}
		o := OutputJsonToWriter(nil, onError)
		o(0, "", LogFields{"a": func(){}})
		if onErrorCalls != 1 {
			t.Fatalf("Expeted to call onError")
		}
	})
	t.Run("Should call Writer() with the json representation of the fields", func(t *testing.T) {
		fields := LogFields{"a": "aaa", "b": "bbb", "c": "ccc"}
		calls := 0
		writer := &mockWriter{func(data []byte) (int, error) {
			calls += 1
			if string(data) != "{\"a\":\"aaa\",\"b\":\"bbb\",\"c\":\"ccc\"}\n" {
				t.Fatalf("Wrong json given to Write()")
			}
			return 0, nil
		}}
		o := OutputJsonToWriter(writer, nil)
		o(0, "", fields)
		if calls != 1 {
			t.Fatalf("Expected to call Writer() one time")
		}
	})
}

func TestOutputToAnsiStdout(t *testing.T) {
	t.Run("Should correctly write to the StdOut", raceFreeTest(func(t *testing.T) {
		oldStdOut := os.Stdout
		defer func() { os.Stdout = oldStdOut }()

		tmpFile, e := ioutil.TempFile(os.TempDir(), "test-")
		defer os.Remove(tmpFile.Name())
		if e != nil {
			t.Fatalf("Error not expected")
		}
		os.Stdout = tmpFile

		OutputToAnsiStdout(LvlError, "some msg", nil)

		b := make([]byte, 28)
		_, e = os.Stdout.ReadAt(b, 0)
		if e != nil {
			t.Fatalf("Error not expected")
		}
		if string(b) != ColorizeStrByLvl(LvlError, "[ ERROR ] some msg") + "\n" {
			t.Fatalf("Expected a different msg")
		}
	}, wStdOut))
}

type mockWriter struct {
	mockWrite func(data []byte) (int, error)
}

func (m *mockWriter) Write(data []byte) (int, error) {
	if m.mockWrite != nil {
		return m.mockWrite(data)
	}
	return 0, nil
}
