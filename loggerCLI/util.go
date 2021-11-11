// Matheus Leonel Balduino
// Everywhere, under @mathbalduino
//   @mathbalduino on GitHub
//   @mathbalduino on Instagram
//   @mathbalduino on Twitter
// Live at mathbalduino.com.br
// 2021-11-09 4:05 PM

package loggerCLI

import (
	logger "github.com/mathbalduino/go-log"
	"strings"
)

// ParseLogLevel will take a string (usually from the CLI
// flag arguments) and returns the equivalent uint64 that
// enables the selected log levels.
//
// You can set some CLI flag that accepts human-readable
// log level selection, using the "|" separator. The following
// imaginary flag value, for example, will enable the DEBUG,
// TRACE and WARN log levels:
// 		--log-levels="DEBUG|TRACE|WARN"
//
// Note that the string is not expected to have blank spaces
// in between the selected log levels
func ParseLogLevel(str string) uint64 {
	logLevels := uint64(0)
	strLvls := strings.Split(str, "|")
	for _, strLvl := range strLvls {
		switch strings.ToUpper(strLvl) {
		case "TRACE":
			logLevels = logLevels | logger.LvlTrace
			continue
		case "DEBUG":
			logLevels = logLevels | logger.LvlDebug
			continue
		case "INFO":
			logLevels = logLevels | logger.LvlInfo
			continue
		case "WARN":
			logLevels = logLevels | logger.LvlWarn
			continue
		case "ERROR":
			logLevels = logLevels | logger.LvlError
			continue
		case "FATAL":
			logLevels = logLevels | logger.LvlFatal
			continue
		}
	}

	return logLevels
}
