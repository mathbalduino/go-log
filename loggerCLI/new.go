package loggerCLI

import (
	"fmt"
	"github.com/mathbalduino/go-log"
	"os"
)

type LoggerCLI interface {
	Trace(format string, args ...interface{}) LoggerCLI
	Debug(format string, args ...interface{}) LoggerCLI
	Info(format string, args ...interface{}) LoggerCLI
	Warn(format string, args ...interface{}) LoggerCLI
	Error(format string, args ...interface{}) LoggerCLI
	Fatal(format string, args ...interface{})
	ErrorFrom(e error) LoggerCLI
	FatalFrom(e error)
}

type loggerCLI struct {
	baseLogger golog.Logger
}

func New(json bool, lvlsEnabled uint64) LoggerCLI {
	output := golog.OutputAnsiToStdout
	if json {
		output = golog.OutputJsonToWriter(os.Stdout, func(err error) { panic(fmt.Errorf("loggerCLI: %w", err)) })
	}
	conf := golog.DefaultConfig()
	conf.LvlsEnabled = lvlsEnabled

	return &loggerCLI{
		golog.New(conf).
			RawOutputs(output, golog.OutputPanicOnFatal),
	}
}
