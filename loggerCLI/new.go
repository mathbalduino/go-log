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
	baseLogger logger.Logger
}

func New(json, debug, trace bool) LoggerCLI {
	output := logger.OutputAnsiToStdout
	if json {
		output = logger.OutputJsonToWriter(os.Stdout, func(err error) { panic(fmt.Errorf("loggerCLI: %w", err)) })
	}
	conf := logger.DefaultConfig()
	conf.LvlsEnabled = logger.LvlProduction
	if debug {
		conf.LvlsEnabled = conf.LvlsEnabled | logger.LvlDebug
	}
	if trace {
		conf.LvlsEnabled = conf.LvlsEnabled | logger.LvlDebug | logger.LvlTrace
	}

	return &loggerCLI{
		logger.New(conf).
			RawOutputs(output, logger.OutputPanicOnFatal),
	}
}
