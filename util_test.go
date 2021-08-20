package loxeLog

import (
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
