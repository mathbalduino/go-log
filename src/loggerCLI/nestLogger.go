package loggerCLI

import (
	"fmt"
	"github.com/mathbalduino/go-log/src"
	"time"
)

func nestLogger(l_ *LoggerCLI) *logger.Logger {
	l := (*logger.Logger)(l_)
	fields := logger.LogFields{TimestampFieldName: fmt.Sprintf("%d", time.Now().UnixNano())}
	parent := l.Field(TimestampFieldName)
	if parent != nil {
		fields[ParentFieldName] = parent
	}
	newLogger := l.Fields(fields)
	return newLogger
}

const ParentFieldName = "parent"
const TimestampFieldName = "timestamp"
