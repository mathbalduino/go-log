package src

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLogger_Configuration(t *testing.T) {
	t.Run("Should set the Logger Configuration to the received arg", func(t *testing.T) {
		c1 := &Configuration{LvlFieldName: "lvlFieldName_1"}
		c2 := Configuration{LvlFieldName: "lvlFieldName_2"}

		l := &Logger{configuration: c1}
		l.Configuration(c2)
		if l.configuration.LvlFieldName != c2.LvlFieldName {
			t.Fatalf("Expected to change the Logger Configuration to the received arg")
		}
	})
}

func TestDefaultConfig(t *testing.T) {
	t.Run("Should be synchronous (nil AsyncScheduler)", func(t *testing.T) {
		c := DefaultConfig()
		if c.AsyncScheduler != nil {
			t.Fatalf("Expected to return a synchronous configuration")
		}
	})
	t.Run("Level field should be 'lvl'", func(t *testing.T) {
		c := DefaultConfig()
		if c.LvlFieldName != "lvl" {
			t.Fatalf("Expected a configuration with the level field set to 'lvl'")
		}
	})
	t.Run("Message field should be 'msg'", func(t *testing.T) {
		c := DefaultConfig()
		if c.MsgFieldName != "msg" {
			t.Fatalf("Expected a configuration with the message field set to 'msg'")
		}
	})
	t.Run("Should enable only the default log levels", func(t *testing.T) {
		c := DefaultConfig()
		if c.LvlsEnabled != LvlDefaults {
			t.Fatalf("Expected a configuration with only the default log levels enabled")
		}
	})
	t.Run("Should set the DefaultErrorParser as the default one", func(t *testing.T) {
		c := DefaultConfig()
		if reflect.ValueOf(c.ErrorParser).Pointer() != reflect.ValueOf(DefaultErrorParser).Pointer() {
			t.Fatalf("Expected the correct ErrorParser")
		}
	})
}

func TestDefaultErrorParser(t *testing.T) {
	t.Run("Should return the error string and the correct map", func(t *testing.T) {
		err := fmt.Errorf("some error msg")
		str, fields := DefaultErrorParser(err)
		if str != err.Error() {
			t.Fatalf("Expected a configuration with an error parser that returns the error msg")
		}
		if !reflect.DeepEqual(fields, LogFields{DefaultErrorKey: err}) {
			t.Fatalf("Expected a configuration with an error parser that returns the correct map")
		}
	})
}

func TestValidateConfig(t *testing.T) {
	t.Run("Should return the right error when the given configuration has the same LvlFieldName and MsgFieldName", func(t *testing.T) {
		fn := func(err error) (string, LogFields) { return "", nil }
		if validateConfig(Configuration{MsgFieldName: "a", LvlFieldName: "a", ErrorParser: fn}) != ErrLvlMsgSameKey {
			t.Fatalf("Expected a different error token")
		}
	})
	t.Run("Should return the right error when the given configuration has a nil ErrorParser", func(t *testing.T) {
		if validateConfig(Configuration{MsgFieldName: "a", LvlFieldName: "b"}) != ErrNilErrorParser {
			t.Fatalf("Expected a different error token")
		}
	})
	t.Run("Should return nil error if LvlFieldName != MsgFieldName, there's an ErrorParser and XXX", func(t *testing.T) {
		fn := func(err error) (string, LogFields) { return "", nil }
		if validateConfig(Configuration{MsgFieldName: "a", LvlFieldName: "b", ErrorParser: fn}) != nil {
			t.Fatalf("Expected a nil error")
		}
	})
}
