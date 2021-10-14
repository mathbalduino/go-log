package loggerCLI

import (
	"fmt"
	"github.com/mathbalduino/go-log"
	"time"
)

func nestLogger(l *loggerCLI) *loggerCLI {
	fields := logger.LogFields{timestampFieldName: fmt.Sprintf("%d", time.Now().UnixNano())}
	parent := l.baseLogger.Field(timestampFieldName)
	if parent != nil {
		fields[parentFieldName] = parent
	}
	return &loggerCLI{
		l.baseLogger.Fields(fields),
	}
}

const parentFieldName = "parent"
const timestampFieldName = "timestamp"
