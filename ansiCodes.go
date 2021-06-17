package loxeLog

// The codes that change the terminal color
const (
	terminalColorReset      = "\033[0m"
	terminalColorRed        = "\033[31m"
	terminalBoldColorRed    = "\033[31;1m"
	terminalColorYellow     = "\033[33m"
	terminalColorCyan       = "\033[36m"
	terminalColorWhite      = "\033[37m"
	terminalColorPurple     = "\033[35m"
	terminalBoldColorPurple = "\033[35;1m"
)

func RedString(msg string) string        { return colorizeString(msg, terminalColorRed) }
func BoldRedString(msg string) string    { return colorizeString(msg, terminalBoldColorRed) }
func YellowString(msg string) string     { return colorizeString(msg, terminalColorYellow) }
func CyanString(msg string) string       { return colorizeString(msg, terminalColorCyan) }
func WhiteString(msg string) string      { return colorizeString(msg, terminalColorWhite) }
func PurpleString(msg string) string     { return colorizeString(msg, terminalColorPurple) }
func BoldPurpleString(msg string) string { return colorizeString(msg, terminalBoldColorPurple) }

// colorizeString will take some string and wrap it with
// the selected terminal color code, colorizing it.
func colorizeString(msg, color string) string { return color + msg + terminalColorReset }
