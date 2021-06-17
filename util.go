package loxeLog

import (
	"gitlab.com/loxe-tools/go-base-library/util"
)

func tryRead(key string, f ...LogFields) interface{} {
	if len(f) == 0 || len(f[0]) == 0 {
		return nil
	}

	value, exists := f[0][key]
	if !exists {
		return nil
	}
	return value
}

func cloneOrNew(f LogFields) LogFields {
	n := LogFields{}
	for key, value := range f {
		n[key] = value
	}
	return n
}

func mergeOverriding(dest LogFields, srcs ...LogFields) {
	if len(srcs) == 0 {
		return
	}

	for key, value := range srcs[0] {
		dest[key] = value
	}
}

func mergeOverriding_(dest Hooks, srcs ...Hooks) {
	if len(srcs) == 0 {
		return
	}

	for key, value := range srcs[0] {
		dest[key] = value
	}
}

func applyHooks(log Log, fields LogFields, hooks map[string]func(Log) interface{}) {
	for key, hook := range hooks {
		fields[key] = hook(log)
	}
}

func notEnabled(flags uint64, logLvl uint64) bool { return (flags & logLvl) == 0 }

func LvlToString(lvl uint64) string {
	switch lvl {
	case LvlTrace:
		return "TRACE"
	case LvlDebug:
		return "DEBUG"
	case LvlInfo:
		return "INFO"
	case LvlWarn:
		return "WARN"
	case LvlError:
		return "ERROR"
	case LvlFatal:
		return "FATAL"
	default:
		return "____"
	}
}

func ColorizeStrByLvl(lvl uint64, msg string) string {
	switch lvl {
	case LvlTrace:
		return msg
	case LvlDebug:
		return util.WhiteString(msg)
	case LvlInfo:
		return util.CyanString(msg)
	case LvlWarn:
		return util.YellowString(msg)
	case LvlError:
		return util.RedString(msg)
	case LvlFatal:
		return util.BoldRedString(msg)
	default:
		return msg
	}
}

func cloneLogger(l *Logger) *Logger {
	return &Logger{
		configuration: l.configuration,
		fields:        l.fields,
		syncHooks:     l.syncHooks,
		asyncHooks:    l.asyncHooks,
		outputs:       l.outputs,
	}
}
