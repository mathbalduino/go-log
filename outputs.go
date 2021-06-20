package loxeLog

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Output is just an alias to a function that is intended to
// handle newly created logs and persist/print it to somewhere
type Output = func(lvl uint64, msg string, fields LogFields)

// Outputs will append the given output params (including the variadic ones)
// to the Logger outputs and return a new Logger instance
func (l *Logger) Outputs(output Output, outputs ...Output) *Logger {
	newLogger := cloneLogger(l)
	newLogger.outputs = append(newLogger.outputs, output)
	newLogger.outputs = append(newLogger.outputs, outputs...)
	return newLogger
}

// RawOutputs will set the given output params (including the variadic ones)
// to the Logger outputs, discarding the old value and returning a new Logger
// instance
func (l *Logger) RawOutputs(output Output, outputs ...Output) *Logger {
	newLogger := cloneLogger(l)
	newLogger.outputs = []Output{output}
	newLogger.outputs = append(newLogger.outputs, outputs...)
	return newLogger
}

// OutputToAnsiTerm will take some log and print it to the
// terminal, using ANSI codes to colorize it accordingly to
// the log level
func OutputToAnsiTerm(lvl uint64, msg string, _ LogFields) {
	msg = fmt.Sprintf("[ %s ] %s", LvlToString(lvl), strings.ReplaceAll(msg, "\n", "\n\t"))
	fmt.Println(ColorizeStrByLvl(lvl, msg))
}

// OutputJsonToFile will parse the log fields to JSON and write
// it to the given Writer interface, calling onError with any errors
// that occur between the Write() call
func OutputJsonToFile(w io.Writer, onError func(error)) Output {
	return func(_ uint64, _ string, fields LogFields) {
		j, e := json.Marshal(fields)
		if e != nil {
			onError(e)
			return
		}

		n, e := w.Write(append(j, "\n"...))
		if e != nil {
			onError(e)
			return
		}
		if n != len(j)+1 {
			onError(ErrIncompleteFileWrite)
		}
	}
}
