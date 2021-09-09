package loggerCLI

import (
	"fmt"
	"gitlab.com/loxe-tools/go-log"
	"os"
)

type LoggerCLI logger.Logger

func New(json, debug, trace bool) *LoggerCLI {
	output := logger.OutputAnsiToStdout
	if json {
		output = logger.OutputJsonToWriter(os.Stdout, func(err error) {  panic(fmt.Errorf("loggerCLI: %w", err)) })
	}
	conf := logger.DefaultConfig()
	conf.LvlsEnabled = logger.LvlProduction
	if debug {
		conf.LvlsEnabled = conf.LvlsEnabled | logger.LvlDebug
	}
	if trace {
		conf.LvlsEnabled = conf.LvlsEnabled | logger.LvlDebug | logger.LvlTrace
	}

	return (*LoggerCLI)(logger.New(conf).
		RawOutputs(output, logger.OutputPanicOnFatal))
}
