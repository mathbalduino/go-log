package loggerCLI

import (
	"fmt"
	logger "gitlab.com/loxe-tools/go-log"
	"os"
)

func outputANSI(lvl uint64, msg string, fields logger.LogFields) {
	logger.OutputToAnsiStdout(lvl, msg, fields)
	handleFatal(lvl, msg, fields)
}

func outputJson(lvl uint64, msg string, fields logger.LogFields) {
	logger.OutputJsonToWriter(os.Stdout, nil)(lvl, msg, fields)
	handleFatal(lvl, msg, fields)
}

func handleFatal(lvl uint64, msg string, fields logger.LogFields) {
	if lvl == logger.LvlFatal {
		err, ok := fields[logger.DefaultErrorParserKey].(error)
		if ok {
			panic(err)
		}
		panic(fmt.Errorf(msg))
	}
}
