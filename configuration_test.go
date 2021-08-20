package loxeLog

import (
	"fmt"
	"testing"
)

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
	t.Run("Should set an error parser that returns the error message and nil fields", func(t *testing.T) {
		c := DefaultConfig()
		if c.ErrorParser == nil {
			t.Fatalf("Expected a configuration with a non-nil error parser")
		}
		errMsg := "some error message"
		str, fields := c.ErrorParser(fmt.Errorf(errMsg))
		if str != errMsg {
			t.Fatalf("Expected a configuration with an error parser that returns the error msg")
		}
		if fields != nil {
			t.Fatalf("Expected a configuration with an error parser that always returns nil fields")
		}
	})
}

func TestValidateConfig(t *testing.T) {
	t.Run("Should return the right error when the given configuration has the same LvlFieldName and MsgFieldName", func(t *testing.T) {
		if validateConfig(Configuration{MsgFieldName: "a", LvlFieldName: "a"}) != ErrLvlMsgSameKey {
			t.Fatalf("Expected a different error token")
		}
	})
	t.Run("Should return nil error if the LvlFieldName and MsgFieldName are different", func(t *testing.T) {
		if validateConfig(Configuration{MsgFieldName: "a", LvlFieldName: "b"}) != nil {
			t.Fatalf("Expected a nil error")
		}
	})
}
