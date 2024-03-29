package golog

import (
	"reflect"
	"strings"
	"testing"
)

func TestLvlToString(t *testing.T) {
	t.Run("LvlTrace should be converted to 'TRACE'", func(t *testing.T) {
		if LvlToString(LvlTrace) != "TRACE" {
			t.Fatalf("Expected LvlTrace to be converted to 'TRACE'")
		}
	})
	t.Run("LvlDebug should be converted to 'DEBUG'", func(t *testing.T) {
		if LvlToString(LvlDebug) != "DEBUG" {
			t.Fatalf("Expected LvlDebug to be converted to 'DEBUG'")
		}
	})
	t.Run("LvlInfo should be converted to 'INFO'", func(t *testing.T) {
		if LvlToString(LvlInfo) != "INFO" {
			t.Fatalf("Expected LvlInfo to be converted to 'INFO'")
		}
	})
	t.Run("LvlWarn should be converted to 'WARN'", func(t *testing.T) {
		if LvlToString(LvlWarn) != "WARN" {
			t.Fatalf("Expected LvlWarn to be converted to 'WARN'")
		}
	})
	t.Run("LvlError should be converted to 'ERROR'", func(t *testing.T) {
		if LvlToString(LvlError) != "ERROR" {
			t.Fatalf("Expected LvlError to be converted to 'ERROR'")
		}
	})
	t.Run("LvlFatal should be converted to 'FATAL'", func(t *testing.T) {
		if LvlToString(LvlFatal) != "FATAL" {
			t.Fatalf("Expected LvlFatal to be converted to 'FATAL'")
		}
	})
	t.Run("Unrecognized levels should return '????'", func(t *testing.T) {
		if LvlToString(9856) != "????" {
			t.Fatalf("Expected unrecognized levels to be converted to '????'")
		}
	})
}

func TestColorizeStrByLvl(t *testing.T) {
	t.Run("The msg should be wrapped by DarkGrey and Reset, when lvl is LvlTrace", func(t *testing.T) {
		msg := "some msg"
		s := ColorizeStrByLvl(LvlTrace, msg)
		if !strings.HasPrefix(s, ansiCodeDarkGrey) {
			t.Fatalf("Expected to prefix the msg with the correct ANSI color code")
		}
		if !strings.HasSuffix(s, ansiCodeReset) {
			t.Fatalf("Expected to be suffixed with the reset ANSI code")
		}
		if !strings.Contains(s, msg) {
			t.Fatalf("Expected to preserve the original msg content")
		}
	})
	t.Run("The msg should be wrapped by LightGrey and Reset, when lvl is LvlDebug", func(t *testing.T) {
		msg := "some msg"
		s := ColorizeStrByLvl(LvlDebug, msg)
		if !strings.HasPrefix(s, ansiCodeLightGrey) {
			t.Fatalf("Expected to prefix the msg with the correct ANSI color code")
		}
		if !strings.HasSuffix(s, ansiCodeReset) {
			t.Fatalf("Expected to be suffixed with the reset ANSI code")
		}
		if !strings.Contains(s, msg) {
			t.Fatalf("Expected to preserve the original msg content")
		}
	})
	t.Run("The msg should be wrapped by Cyan and Reset, when lvl is LvlInfo", func(t *testing.T) {
		msg := "some msg"
		s := ColorizeStrByLvl(LvlInfo, msg)
		if !strings.HasPrefix(s, ansiCodeCyan) {
			t.Fatalf("Expected to prefix the msg with the correct ANSI color code")
		}
		if !strings.HasSuffix(s, ansiCodeReset) {
			t.Fatalf("Expected to be suffixed with the reset ANSI code")
		}
		if !strings.Contains(s, msg) {
			t.Fatalf("Expected to preserve the original msg content")
		}
	})
	t.Run("The msg should be wrapped by Yellow and Reset, when lvl is LvlWarn", func(t *testing.T) {
		msg := "some msg"
		s := ColorizeStrByLvl(LvlWarn, msg)
		if !strings.HasPrefix(s, ansiCodeYellow) {
			t.Fatalf("Expected to prefix the msg with the correct ANSI color code")
		}
		if !strings.HasSuffix(s, ansiCodeReset) {
			t.Fatalf("Expected to be suffixed with the reset ANSI code")
		}
		if !strings.Contains(s, msg) {
			t.Fatalf("Expected to preserve the original msg content")
		}
	})
	t.Run("The msg should be wrapped by Red and Reset, when lvl is LvlError", func(t *testing.T) {
		msg := "some msg"
		s := ColorizeStrByLvl(LvlError, msg)
		if !strings.HasPrefix(s, ansiCodeRed) {
			t.Fatalf("Expected to prefix the msg with the correct ANSI color code")
		}
		if !strings.HasSuffix(s, ansiCodeReset) {
			t.Fatalf("Expected to be suffixed with the reset ANSI code")
		}
		if !strings.Contains(s, msg) {
			t.Fatalf("Expected to preserve the original msg content")
		}
	})
	t.Run("The msg  should be wrapped by BoldRed and Reset, when lvl is LvlFatal", func(t *testing.T) {
		msg := "some msg"
		s := ColorizeStrByLvl(LvlFatal, msg)
		if !strings.HasPrefix(s, ansiCodeBoldRed) {
			t.Fatalf("Expected to prefix the msg with the correct ANSI color code")
		}
		if !strings.HasSuffix(s, ansiCodeReset) {
			t.Fatalf("Expected to be suffixed with the reset ANSI code")
		}
		if !strings.Contains(s, msg) {
			t.Fatalf("Expected to preserve the original msg content")
		}
	})
	t.Run("Unrecognized levels should just return the msg", func(t *testing.T) {
		msg := "some msg"
		s := ColorizeStrByLvl(685168, msg)
		if s != msg {
			t.Fatalf("Expected to return the original msg")
		}
	})
}

func TestTryRead(t *testing.T) {
	t.Run("Return nil if the len of the variadic param 'f' is zero", func(t *testing.T) {
		if tryRead("") != nil {
			t.Fatalf("Expected to return nil")
		}
	})
	t.Run("Return nil if there's no variadic param that contains the given key", func(t *testing.T) {
		// Note that the second variadic param should be ignored even if it contains the key
		key := "someKey"
		if tryRead(key, map[string]interface{}{"key": "value"}, map[string]interface{}{"key2": "a"}) != nil {
			t.Fatalf("Expected to return nil when the key is not present")
		}
	})
	t.Run("Return the value of the given key from the latter variadic param", func(t *testing.T) {
		// Note that the second variadic param should be always ignored
		key := "someKey"
		v := tryRead(key, map[string]interface{}{key: "value"}, map[string]interface{}{key: "b"})
		if v == nil {
			t.Fatalf("Expected to return the actual value")
		}
		if v.(string) != "b" {
			t.Fatalf("Expected to return the value of the latter variadic param")
		}
	})
}

func TestCloneOrNew(t *testing.T) {
	t.Run("If the given fields is empty/nil, return a new, non-nil, empty fields", func(t *testing.T) {
		f := cloneOrNew((LogFields)(nil))
		if f == nil {
			t.Fatalf("Not expected to be nil")
		}
		if len(f) != 0 {
			t.Fatalf("Expected to return an empty fields")
		}

		f = cloneOrNew(LogFields{})
		if f == nil {
			t.Fatalf("Not expected to be nil")
		}
		if len(f) != 0 {
			t.Fatalf("Expected to return an empty fields")
		}
	})
	t.Run("Should always return a new, different LogFields", func(t *testing.T) {
		p := LogFields{}
		f := cloneOrNew(p)
		if reflect.ValueOf(p).Pointer() == reflect.ValueOf(f).Pointer() {
			t.Fatalf("Expected to return a different map, not the same")
		}
	})
	t.Run("Should copy all the given LogFields entries and values, returning a new LogFields", func(t *testing.T) {
		p := LogFields{"a": "aa", "b": "bb", "c": "cc"}
		f := cloneOrNew(p)
		if reflect.ValueOf(p).Pointer() == reflect.ValueOf(f).Pointer() {
			t.Fatalf("Expected to return a different map, not the same")
		}
		if !reflect.DeepEqual(p, f) {
			t.Fatalf("Expected the returned map to be equivalent to the given one")
		}
	})
}

func TestCloneOrNew_(t *testing.T) {
	t.Run("If the given hooks param is empty/nil, return a new, non-nil, empty hooks", func(t *testing.T) {
		f := cloneOrNew((Hooks)(nil))
		if f == nil {
			t.Fatalf("Not expected to be nil")
		}
		if len(f) != 0 {
			t.Fatalf("Expected to return an empty Hooks")
		}

		f = cloneOrNew(Hooks{})
		if f == nil {
			t.Fatalf("Not expected to be nil")
		}
		if len(f) != 0 {
			t.Fatalf("Expected to return an empty Hooks")
		}
	})
	t.Run("Should always return a new, different Hooks", func(t *testing.T) {
		p := Hooks{}
		f := cloneOrNew(p)
		if reflect.ValueOf(p).Pointer() == reflect.ValueOf(f).Pointer() {
			t.Fatalf("Expected to return a different map, not the same")
		}
	})
	t.Run("Should copy all the given Hooks entries and values, returning a new Hooks", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		fnB := func(log Log) interface{} { return nil }
		fnC := func(log Log) interface{} { return nil }
		p := Hooks{"a": fnA, "b": fnB, "c": fnC}
		f := cloneOrNew(p)
		if reflect.ValueOf(p).Pointer() == reflect.ValueOf(f).Pointer() {
			t.Fatalf("Expected to return a different map, not the same")
		}
		if reflect.ValueOf(f["a"]).Pointer() != reflect.ValueOf(fnA).Pointer() ||
			reflect.ValueOf(f["b"]).Pointer() != reflect.ValueOf(fnB).Pointer() ||
			reflect.ValueOf(f["c"]).Pointer() != reflect.ValueOf(fnC).Pointer() {
			t.Fatalf("Expected the returned map to be equivalent to the given one")
		}
	})
}

func TestMergeOverriding(t *testing.T) {
	t.Run("If the len of the variadic param is zero, just exit, ignoring the dest param", func(t *testing.T) {
		dest := LogFields{"a": "aaa"}
		mergeOverriding(dest)
		if len(dest) != 1 || dest["a"] != "aaa" {
			t.Fatalf("Not expected to apply side effects to the dest param")
		}
	})
	t.Run("If the len of the variadic param is zero, just exit, ignoring the dest param", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		dest := Hooks{"a": fnA}
		mergeOverriding(dest)
		if len(dest) != 1 || reflect.ValueOf(dest["a"]).Pointer() != reflect.ValueOf(fnA).Pointer() {
			t.Fatalf("Not expected to apply side effects to the dest param")
		}
	})
	t.Run("Should copy all the variadic params to the dest param. The later always overrides the previous", func(t *testing.T) {
		dest := LogFields{"a": "aaa"}
		mergeOverriding(dest,
			LogFields{"c": "ccc", "d": "ddd"},
			LogFields{"t": "ttt"},
			LogFields{"c": "CCC", "p": "ppp"},
			LogFields{"a": "AAA", "t": "TTT"})
		if len(dest) != 5 {
			t.Fatalf("The dest arg doesn't have the correct len")
		}
		if dest["a"] != "AAA" || dest["c"] != "CCC" || dest["t"] != "TTT" || dest["p"] != "ppp" || dest["d"] != "ddd" {
			t.Fatalf("The dest arg doesn't contains the correct overrided values")
		}
	})
}

func TestApplyHooks(t *testing.T) {
	t.Run("If the given hooks are empty, noop", func(t *testing.T) {
		f := LogFields{"a": "aaa"}
		applyHooks(Log{}, f, nil)
		if len(f) != 1 || f["a"] != "aaa" {
			t.Fatalf("Not expected to change the given LogFields")
		}
	})
	t.Run("Should call each hook, passing the given log as argument, setting the equivalent LogField key/value", func(t *testing.T) {
		msg := "Some Log Msg"
		callsA, callsB, callsC := 0, 0, 0
		h := Hooks{
			"a": func(log Log) interface{} {
				callsA += 1
				if log.msg != msg {
					t.Fatalf("Expected to pass the correct log to each hook")
				}
				return "aaa"
			},
			"b": func(log Log) interface{} {
				callsB += 1
				if log.msg != msg {
					t.Fatalf("Expected to pass the correct log to each hook")
				}
				return "bbb"
			},
			"c": func(log Log) interface{} {
				callsC += 1
				if log.msg != msg {
					t.Fatalf("Expected to pass the correct log to each hook")
				}
				return "ccc"
			},
		}

		f := LogFields{"a": "aaa"}
		applyHooks(Log{msg: msg}, f, h)
		if len(f) != 3 || f["a"] != "aaa" || f["b"] != "bbb" || f["c"] != "ccc" {
			t.Fatalf("Expected to set the given LogFields entry to the return of each hooks")
		}
		if callsA != 1 || callsB != 1 || callsC != 1 {
			t.Fatalf("Each hook must be called just one time")
		}
	})
}

func TestNotEnabled(t *testing.T) {
	t.Run("If the bitwise AND operation between the integers return zero, it's not enabled, return true", func(t *testing.T) {
		flags := uint64(0b1010)
		logLv := uint64(0b0101)
		if !notEnabled(flags, logLv) {
			t.Fatalf("Expected to return true")
		}
	})
	t.Run("If there's two (or more) equivalent bits set to 1, return false", func(t *testing.T) {
		flags := uint64(0b1011)
		logLv := uint64(0b1101)
		if notEnabled(flags, logLv) {
			t.Fatalf("Expected to return false")
		}
	})
	t.Run("If one of them is zero, return true", func(t *testing.T) {
		flags := uint64(0)
		logLv := uint64(0b1111)
		if !notEnabled(flags, logLv) {
			t.Fatalf("Expected to return true")
		}
	})
}

func TestCloneLogger(t *testing.T) {
	t.Run("Should return a different Logger pointer", func(t *testing.T) {
		l := &logger{}
		l2 := cloneLogger(l)
		if reflect.ValueOf(l2).Pointer() == reflect.ValueOf(l).Pointer() {
			t.Fatalf("Expected to return a new Logger, not the same")
		}
	})
	t.Run("Should just copy the configuration pointer", func(t *testing.T) {
		l := &logger{configuration: &Configuration{}}
		l2 := cloneLogger(l)
		if reflect.ValueOf(l2.configuration).Pointer() != reflect.ValueOf(l.configuration).Pointer() {
			t.Fatalf("Expected to return a new Logger with the same configuration")
		}
	})
	t.Run("Should clone the fields into a new LogFields", func(t *testing.T) {
		l := &logger{fields: LogFields{"a": "aaa", "b": "bbb", "c": "ccc"}}
		l2 := cloneLogger(l)
		if reflect.ValueOf(l2.fields).Pointer() == reflect.ValueOf(l.fields).Pointer() {
			t.Fatalf("Expected to return a Logger with a new LogFields")
		}
		if !reflect.DeepEqual(l2.fields, l.fields) {
			t.Fatalf("Expected the new Logger.fields to be equivalent")
		}
	})
	t.Run("Should clone the preHooks into a new Hooks", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		fnB := func(log Log) interface{} { return nil }
		fnC := func(log Log) interface{} { return nil }
		l := &logger{preHooks: Hooks{"a": fnA, "b": fnB, "c": fnC}}
		l2 := cloneLogger(l)
		if reflect.ValueOf(l2.preHooks).Pointer() == reflect.ValueOf(l.preHooks).Pointer() {
			t.Fatalf("Expected to return a Logger with a new preHooks")
		}
		if reflect.ValueOf(l2.preHooks["a"]).Pointer() != reflect.ValueOf(l.preHooks["a"]).Pointer() ||
			reflect.ValueOf(l2.preHooks["b"]).Pointer() != reflect.ValueOf(l.preHooks["b"]).Pointer() ||
			reflect.ValueOf(l2.preHooks["c"]).Pointer() != reflect.ValueOf(l.preHooks["c"]).Pointer() {
			t.Fatalf("Expected the new Logger.preHooks to be equivalent")
		}
	})
	t.Run("Should clone the postHooks into a new Hooks", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		fnB := func(log Log) interface{} { return nil }
		fnC := func(log Log) interface{} { return nil }
		l := &logger{postHooks: Hooks{"a": fnA, "b": fnB, "c": fnC}}
		l2 := cloneLogger(l)
		if reflect.ValueOf(l2.postHooks).Pointer() == reflect.ValueOf(l.postHooks).Pointer() {
			t.Fatalf("Expected to return a Logger with a new postHooks")
		}
		if reflect.ValueOf(l2.postHooks["a"]).Pointer() != reflect.ValueOf(l.postHooks["a"]).Pointer() ||
			reflect.ValueOf(l2.postHooks["b"]).Pointer() != reflect.ValueOf(l.postHooks["b"]).Pointer() ||
			reflect.ValueOf(l2.postHooks["c"]).Pointer() != reflect.ValueOf(l.postHooks["c"]).Pointer() {
			t.Fatalf("Expected the new Logger.postHooks to be equivalent")
		}
	})
	t.Run("Should clone the outputs slice into a new []Output", func(t *testing.T) {
		outA := func(lvl uint64, msg string, fields LogFields) {}
		outB := func(lvl uint64, msg string, fields LogFields) {}
		outC := func(lvl uint64, msg string, fields LogFields) {}
		l := &logger{outputs: []Output{outA, outB, outC}}
		l2 := cloneLogger(l)
		if reflect.ValueOf(l2.outputs).Pointer() == reflect.ValueOf(l.outputs).Pointer() {
			t.Fatalf("Expected to return a Logger with a new outputs")
		}
		if len(l2.outputs) != 3 ||
			reflect.ValueOf(l2.outputs[0]).Pointer() != reflect.ValueOf(l.outputs[0]).Pointer() ||
			reflect.ValueOf(l2.outputs[1]).Pointer() != reflect.ValueOf(l.outputs[1]).Pointer() ||
			reflect.ValueOf(l2.outputs[2]).Pointer() != reflect.ValueOf(l.outputs[2]).Pointer() {
			t.Fatalf("Expected the new Logger.outputs to be equivalent")
		}
	})
}
