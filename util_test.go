package loxeLog

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
	t.Run("Return nil if the len of the FIRST variadic param 'f' is zero", func(t *testing.T) {
		if tryRead("", nil) != nil {
			t.Fatalf("Expected to return nil")
		}
	})
	t.Run("Return nil if the FIRST variadic param doesn't contains the given key", func(t *testing.T) {
		// Note that the second variadic param should be ignored even if it contains the key
		key := "someKey"
		if tryRead(key, map[string]interface{}{"key": "value"}, map[string]interface{}{key: "a"}) != nil {
			t.Fatalf("Expected to return nil when the key is not present")
		}
	})
	t.Run("Return the value of the given key inside the FIRST variadic param", func(t *testing.T) {
		// Note that the second variadic param should be always ignored
		key := "someKey"
		v := tryRead(key, map[string]interface{}{key: "value"}, map[string]interface{}{key: "b"})
		if v == nil {
			t.Fatalf("Expected to return the actual value")
		}
		if v.(string) != "value" {
			t.Fatalf("Expected to return the value of the first variadic param")
		}
	})
}

func TestCloneOrNew(t *testing.T) {
	t.Run("If the given fields is empty/nil, return a new, non-nil, empty fields", func(t *testing.T) {
		f := cloneOrNew(nil)
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

func TestMergeOverriding(t *testing.T) {
	t.Run("If the len of the variadic param is zero, just exit, ignoring the dest param", func(t *testing.T) {
		dest := LogFields{"a": "aaa"}
		mergeOverriding(dest)
		if len(dest) != 1 || dest["a"] != "aaa" {
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

func TestMergeOverriding_(t *testing.T) {
	t.Run("If the len of the variadic param is zero, just exit, ignoring the dest param", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		dest := Hooks{"a": fnA}
		mergeOverriding_(dest)
		if len(dest) != 1 || reflect.ValueOf(dest["a"]).Pointer() != reflect.ValueOf(fnA).Pointer() {
			t.Fatalf("Not expected to apply side effects to the dest param")
		}
	})
	t.Run("Should copy just the FIRST variadic param to the dest param, overriding", func(t *testing.T) {
		fnA := func(log Log) interface{} { return nil }
		fnAA := func(log Log) interface{} { return nil }
		fnB := func(log Log) interface{} { return nil }
		fnC := func(log Log) interface{} { return nil }
		fnD := func(log Log) interface{} { return nil }
		fnE := func(log Log) interface{} { return nil }
		dest := Hooks{"a": fnA}
		mergeOverriding_(dest,
			Hooks{"b": fnB, "a": fnAA},
			Hooks{"d": fnD},
			Hooks{"b": fnC, "e": fnE},
			Hooks{"a": fnAA, "d": nil})
		if len(dest) != 2 {
			t.Fatalf("The dest arg doesn't have the correct len")
		}
		if reflect.ValueOf(dest["a"]).Pointer() != reflect.ValueOf(fnAA).Pointer() ||
			reflect.ValueOf(dest["b"]).Pointer() != reflect.ValueOf(fnB).Pointer() {
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

