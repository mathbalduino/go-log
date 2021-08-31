package logger

// The ANSI codes used to change the terminal color
const (
	ansiCodeReset     = "\033[0m"
	ansiCodeRed       = "\033[31m"
	ansiCodeBoldRed   = "\033[31;1m"
	ansiCodeYellow    = "\033[33m"
	ansiCodeCyan      = "\033[36m"
	ansiCodeLightGrey = "\033[37m"
	ansiCodeDarkGrey  = "\033[90m"
)

// RedString will wrap the given string between the red and
// reset ANSI codes.
//
// Terminals with ANSI code support will print the string to
// the screen using red as the font color
func RedString(msg string) string { return colorizeString(msg, ansiCodeRed) }

// BoldRedString will wrap the given string between the bold
// red and reset ANSI codes.
//
// Terminals with ANSI code support will print the string to
// the screen using bold red as the font color
func BoldRedString(msg string) string { return colorizeString(msg, ansiCodeBoldRed) }

// YellowString will wrap the given string between the yellow
// and reset ANSI codes.
//
// Terminals with ANSI code support will print the string to
// the screen using yellow as the font color
func YellowString(msg string) string { return colorizeString(msg, ansiCodeYellow) }

// CyanString will wrap the given string between the cyan and
// reset ANSI codes.
//
// Terminals with ANSI code support will print the string to
// the screen using cyan as the font color
func CyanString(msg string) string { return colorizeString(msg, ansiCodeCyan) }

// LightGreyString will wrap the given string between the light
// grey and reset ANSI codes.
//
// Terminals with ANSI code support will print the string to
// the screen using light grey as the font color
func LightGreyString(msg string) string { return colorizeString(msg, ansiCodeLightGrey) }

// DarkGreyString will wrap the given string between the dark
// grey and reset ANSI codes.
//
// Terminals with ANSI code support will print the string to
// the screen using dark grey as the font color
func DarkGreyString(msg string) string { return colorizeString(msg, ansiCodeDarkGrey) }

// colorizeString will wrap the given string between the given
// color ANSI code and the reset ANSI code.
func colorizeString(msg, color string) string { return color + msg + ansiCodeReset }
