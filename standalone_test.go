// Matheus Leonel Balduino
// Everywhere, under @mathbalduino
//   @mathbalduino on GitHub
//   @mathbalduino on Instagram
//   @mathbalduino on Twitter
// Live at mathbalduino.com.br
// 2021-11-16 11:02 AM

package logger

import (
	"fmt"
	"testing"
)

func TestTrace(t *testing.T) {
	t.Run("Should not panic", raceFreeTest(func(t *testing.T) {
		Trace("Any msg")
	}, rHandleLog))
}

func TestDebug(t *testing.T) {
	t.Run("Should not panic", raceFreeTest(func(t *testing.T) {
		Debug("Any msg")
	}, rHandleLog))
}

func TestInfo(t *testing.T) {
	t.Run("Should not panic", raceFreeTest(func(t *testing.T) {
		Info("Any msg")
	}, rHandleLog))
}

func TestWarn(t *testing.T) {
	t.Run("Should not panic", raceFreeTest(func(t *testing.T) {
		Warn("Any msg")
	}, rHandleLog))
}

func TestError(t *testing.T) {
	t.Run("Should not panic", raceFreeTest(func(t *testing.T) {
		Error("Any msg")
	}, rHandleLog))
}

func TestFatal(t *testing.T) {
	t.Run("Should panic", raceFreeTest(func(t *testing.T) {
		defer func() {
			e := recover()
			if e == nil {
				t.Fatal("Expected to panic")
			}
			_, isErr := e.(error)
			if !isErr {
				t.Fatal("Expected to panic")
			}
		}()

		Fatal("any msg")
	}, rHandleLog))
}

func TestErrorFrom(t *testing.T) {
	t.Run("Should not panic", raceFreeTest(func(t *testing.T) {
		ErrorFrom(fmt.Errorf("any error msg"))
	}, rHandleLog))
}

func TestFatalFrom(t *testing.T) {
	t.Run("Should panic", raceFreeTest(func(t *testing.T) {
		defer func() {
			e := recover()
			if e == nil {
				t.Fatal("Expected to panic")
			}
			_, isErr := e.(error)
			if !isErr {
				t.Fatal("Expected to panic")
			}
		}()

		FatalFrom(fmt.Errorf("any msg"))
	}, rHandleLog))
}
