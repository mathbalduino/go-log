package loggerCLI

import (
	"fmt"
	"go-log/src"
	"os"
)

type LoggerCLI src.logger

func New(json, debug, trace bool) *LoggerCLI {
	output := src.logger.OutputAnsiToStdout
	if json {
		output = src.logger.OutputJsonToWriter(os.Stdout, func(err error) {  panic(fmt.Errorf("loggerCLI: %w", err)) })
	}
	conf := src.logger.DefaultConfig()
	conf.LvlsEnabled = src.logger.LvlProduction
	if debug {
		conf.LvlsEnabled = conf.LvlsEnabled | src.logger.LvlDebug
	}
	if trace {
		conf.LvlsEnabled = conf.LvlsEnabled | src.logger.LvlDebug | src.logger.LvlTrace
	}

	return (*LoggerCLI)(src.logger.New(conf).
		RawOutputs(output, src.logger.OutputPanicOnFatal))
}
