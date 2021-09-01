package loggerCLI

import (
	"gitlab.com/loxe-tools/go-log"
	"time"
)

type LoggerCLI logger.Logger

func New(json, debug, trace bool) *LoggerCLI {
	output := outputANSI
	if json {
		output = outputJson
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
		RawOutputs(output).
		Fields(logger.LogFields{"timestamp": time.Now().UnixNano()}))
}
