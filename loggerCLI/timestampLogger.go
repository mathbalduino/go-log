package loggerCLI

import (
	logger "gitlab.com/loxe-tools/go-log"
	"time"
)

func timestampLogger(l *logger.Logger) *LoggerCLI {
	newLogger := l.Fields(logger.LogFields{
		"parent":    l.Field("timestamp"),
		"timestamp": time.Now().UnixNano(),
	})
	return (*LoggerCLI)(newLogger)
}
