package loxeLog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

// OutputToWriter will call the given parser and write the results to the given io.Writer, calling
// 'onError' with any errors that occur
func OutputToWriter(w io.Writer, parser func(LogFields) ([]byte, error), onError func(error)) Output {
	return func(_ uint64, _ string, fields LogFields) {
		data, e := parser(fields)
		if e != nil {
			onError(e)
			return
		}

		_, e = w.Write(data)
		if e != nil {
			onError(e)
		}
	}
}

// OutputJsonToFile will parse the log fields to JSON and write
// it to the given Writer interface, calling onError with any errors
// that occur between the Marshal()/Write() call
func OutputJsonToFile(w io.Writer, onError func(error)) Output {
	parser := func(fields LogFields) ([]byte, error) {
		data, e := json.Marshal(fields)
		if e != nil {
			return nil, e
		}
		return append(data, '\n'), nil
	}
	return OutputToWriter(w, parser, onError)
}

// OutputToAnsiStdout will take some log and write it to the
// os.Stdout, using ANSI codes to colorize it accordingly to
// the log level
func OutputToAnsiStdout(lvlFieldName, msgFieldName string) Output {
	return OutputToWriter(os.Stdout, func(f LogFields) ([]byte, error) {
		lvl, msg := f[lvlFieldName].(uint64), f[msgFieldName].(string)
		msg = fmt.Sprintf("[ %s ] %s\n", LvlToString(lvl), strings.ReplaceAll(msg, "\n", "\n\t"))
		msg = ColorizeStrByLvl(lvl, msg)
		return []byte(msg), nil
	}, func(err error) {})
}
