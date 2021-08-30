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

// OutputParser is just an alias to a function that is used
// to parse the LogFields into a byte array, in order to write
// it to a file, etc...
type OutputParser = func(LogFields) ([]byte, error)

// Outputs will append the given output params (including the variadic ones)
// to the Logger outputs and return a new Logger instance
func (l *Logger) Outputs(output Output, outputs ...Output) *Logger {
	newLogger := cloneLogger(l)
	if output != nil {
		newLogger.outputs = append(newLogger.outputs, output)
	}
	for _, output := range outputs {
		if output != nil {
			newLogger.outputs = append(newLogger.outputs, output)
		}
	}
	return newLogger
}

// RawOutputs will set the given output params (including the variadic ones)
// to the Logger outputs, discarding the old value and returning a new Logger
// instance
func (l *Logger) RawOutputs(output Output, outputs ...Output) *Logger {
	newLogger := cloneLogger(l)
	newLogger.outputs = []Output{}
	if output != nil {
		newLogger.outputs = append(newLogger.outputs, output)
	}
	for _, output := range outputs {
		if output != nil {
			newLogger.outputs = append(newLogger.outputs, output)
		}
	}
	return newLogger
}

// OutputToWriter will call the given parser and write the results to the given io.Writer, calling
// 'onError' with any errors that occur
func OutputToWriter(w io.Writer, parser OutputParser, onError func(error)) Output {
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

// OutputJsonToWriter will parse the log fields to JSON and write
// it to the given Writer interface, calling onError with any errors
// that occur between the Marshal()/Write() call
func OutputJsonToWriter(w io.Writer, onError func(error)) Output {
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
func OutputToAnsiStdout(lvl uint64, msg string, _ LogFields) {
	fmt.Printf(
		ColorizeStrByLvl(lvl, "[ %s ] %s") + "\n",
		LvlToString(lvl), strings.ReplaceAll(msg, "\n", "\n\t"))
}
