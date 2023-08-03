package golog

import "testing"

func TestColorizeString(t *testing.T) {
	t.Run("Should prefix the msg with the given color str, and suffix with the reset ANSI code", func(t *testing.T) {
		msg, color := "some msg", "some color ANSI code"
		if colorizeString(msg, color) != color+msg+ansiCodeReset {
			t.Fatalf("Expected that the msg was wrapped inside <ANSI_CODE>msg<RESET_ANSI_CODE>")
		}
	})
	t.Run("Should prefix with DarkGrey, and suffix with the reset ANSI code", func(t *testing.T) {
		msg := "some msg"
		if DarkGreyString(msg) != ansiCodeDarkGrey+msg+ansiCodeReset {
			t.Fatalf("Expected that the msg was wrapped inside <DarkGrey_ANSI_CODE>msg<RESET_ANSI_CODE>")
		}
	})
	t.Run("Should prefix with LightGrey, and suffix with the reset ANSI code", func(t *testing.T) {
		msg := "some msg"
		if LightGreyString(msg) != ansiCodeLightGrey+msg+ansiCodeReset {
			t.Fatalf("Expected that the msg was wrapped inside <LightGrey_ANSI_CODE>msg<RESET_ANSI_CODE>")
		}
	})
	t.Run("Should prefix with Cyan, and suffix with the reset ANSI code", func(t *testing.T) {
		msg := "some msg"
		if CyanString(msg) != ansiCodeCyan+msg+ansiCodeReset {
			t.Fatalf("Expected that the msg was wrapped inside <Cyan_ANSI_CODE>msg<RESET_ANSI_CODE>")
		}
	})
	t.Run("Should prefix with Yellow, and suffix with the reset ANSI code", func(t *testing.T) {
		msg := "some msg"
		if YellowString(msg) != ansiCodeYellow+msg+ansiCodeReset {
			t.Fatalf("Expected that the msg was wrapped inside <Yellow_ANSI_CODE>msg<RESET_ANSI_CODE>")
		}
	})
	t.Run("Should prefix with BoldRed, and suffix with the reset ANSI code", func(t *testing.T) {
		msg := "some msg"
		if BoldRedString(msg) != ansiCodeBoldRed+msg+ansiCodeReset {
			t.Fatalf("Expected that the msg was wrapped inside <BoldRed_ANSI_CODE>msg<RESET_ANSI_CODE>")
		}
	})
	t.Run("Should prefix with Red, and suffix with the reset ANSI code", func(t *testing.T) {
		msg := "some msg"
		if RedString(msg) != ansiCodeRed+msg+ansiCodeReset {
			t.Fatalf("Expected that the msg was wrapped inside <Red_ANSI_CODE>msg<RESET_ANSI_CODE>")
		}
	})
}
