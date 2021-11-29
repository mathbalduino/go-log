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
//
// This function accepts an extra string representing all
// the log levels: "ALL". If the string contains "ALL", or
// is equal to "ALL", all the log levels will be enabled
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
		case "ALL":
			return logger.LvlAll
		}
	}

	return logLevels
}

// ValidateLogLevels will take some string and check to see
// if it's a valid LogLevels string.
//
// If the string contains anything not recognizable as a LogLevel,
// the returned value will be false. Otherwise, returns true
// and the string is safe to be used.
//
// Note that you can use this function to validate some CLI
// flag value.
func ValidateLogLevels(s string) bool {
	strLvls := strings.Split(s, "|")
	for _, strLvl := range strLvls {
		if strLvl == "TRACE" || strLvl == "DEBUG" || strLvl == "INFO" || strLvl == "WARN" || strLvl == "ERROR" || strLvl == "FATAL" || strLvl == "ALL" {
			continue
		}
		return false
	}
	return true
}

// LogLevelsValues is a constant that is intended to be used to describe
// which values are accepted inside the LogLevels string.
//
// When building CLI tools, use this constant to provide "--help" information
// about the flag that controls the log levels
const LogLevelsValues = `"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "ALL"`
