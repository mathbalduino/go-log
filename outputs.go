package loxeLog

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Output = func(lvl uint64, msg string, fields LogFields)

func (l *Logger) Outputs(output Output, outputs ...Output) *Logger {
	l.outputs = append(l.outputs, output)
	l.outputs = append(l.outputs, outputs...)
	return l
}

func (l *Logger) RawOutputs(output Output, outputs ...Output) *Logger {
	l.outputs = []Output{output}
	l.outputs = append(l.outputs, outputs...)
	return l
}

func OutputToAnsiTerm(lvl uint64, msg string, _ LogFields) {
	msg = fmt.Sprintf("[ %s ] %s", LvlToString(lvl), strings.ReplaceAll(msg, "\n", "\n\t"))
	fmt.Println(ColorizeStrByLvl(lvl, msg))
}

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
		if n != len(j) + 1 {
			onError(ErrIncompleteFileWrite)
		}
	}
}
