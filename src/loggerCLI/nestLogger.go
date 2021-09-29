package loggerCLI

import (
	"fmt"
	"go-log/src"
	"time"
)

func nestLogger(l_ *LoggerCLI) *src.Logger {
	l := (*src.Logger)(l_)
	fields := src.LogFields{TimestampFieldName: fmt.Sprintf("%d", time.Now().UnixNano())}
	parent := l.Field(TimestampFieldName)
	if parent != nil {
		fields[ParentFieldName] = parent
	}
	newLogger := l.Fields(fields)
	return newLogger
}

const ParentFieldName = "parent"
const TimestampFieldName = "timestamp"
