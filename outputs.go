package logger

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
// it to a file, etc, idk...
type OutputParser = func(LogFields) ([]byte, error)

func (l *logger) Outputs(output Output, outputs ...Output) Logger {
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

func (l *logger) RawOutputs(output Output, outputs ...Output) Logger {
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

// OutputToWriter will call the given parser and write the returned values to the
// given io.Writer, calling 'onError' with any errors that occur
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

// OutputAnsiToStdout will take some log and write it to the
// os.Stdout, using ANSI codes to colorize it accordingly to
// the log level
func OutputAnsiToStdout(lvl uint64, msg string, _ LogFields) {
	fmt.Printf(ColorizeStrByLvl(lvl, "[ %s ] %s\n"),
		LvlToString(lvl), strings.ReplaceAll(msg, "\n", "\n\t"))
}

// OutputPanicOnFatal will call "panic" if the lvl of the given log
// is LvlFatal, using the "error" interface inside the fields as
// argument, if present. Otherwise, the log msg string will be used
// to create a new error (using fmt.Errorf)
func OutputPanicOnFatal(lvl uint64, msg string, fields LogFields) {
	if lvl == LvlFatal {
		err, ok := fields[DefaultErrorKey].(error)
		if ok {
			panic(err)
		}
		panic(fmt.Errorf(msg))
	}
}
