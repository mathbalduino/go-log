package loxeLog

// LvlToString will take the given log level and
// return the string that represents it
//
// Note that this function can only translate default
// log levels
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
		return "????"
	}
}

// ColorizeStrByLvl will colorize the log msg using the
// ANSI color code associated with the log level
func ColorizeStrByLvl(lvl uint64, msg string) string {
	switch lvl {
	case LvlTrace:
		return DarkGreyString(msg)
	case LvlDebug:
		return LightGreyString(msg)
	case LvlInfo:
		return CyanString(msg)
	case LvlWarn:
		return YellowString(msg)
	case LvlError:
		return RedString(msg)
	case LvlFatal:
		return BoldRedString(msg)
	default:
		return msg
	}
}

// tryRead will read the given key from some LogFields,
// returning nil if it's not present
//
// Note that the 'f' param is a variadic just to sugar the
// syntax (only the first index is used)
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

// cloneOrNew will create a new LogFields and
// copy all values from the given param. If nil/empty,
// returns a new empty LogFields
func cloneOrNew(f LogFields) LogFields {
	n := LogFields{}
	for key, value := range f {
		n[key] = value
	}
	return n
}

// mergeOverriding will copy the values from srcs[0] to dest,
// overriding any existing values
func mergeOverriding(dest LogFields, srcs ...LogFields) {
	if len(srcs) == 0 {
		return
	}

	for key, value := range srcs[0] {
		dest[key] = value
	}
}

// mergeOverriding_ will copy the values from srcs[0] to dest,
// overriding any existing values
func mergeOverriding_(dest Hooks, srcs ...Hooks) {
	if len(srcs) == 0 {
		return
	}

	for key, value := range srcs[0] {
		dest[key] = value
	}
}

// applyHooks will call the given hooks and set the returned
// values to the given fields (using the given log as a parameter)
func applyHooks(log Log, fields LogFields, hooks Hooks) {
	for key, hook := range hooks {
		fields[key] = hook(log)
	}
}

// notEnabled will return true if the given log level is not enabled by the
// given flags
func notEnabled(flags uint64, logLvl uint64) bool { return (flags & logLvl) == 0 }

// cloneLogger will create a new identical Logger
// instance from the given one
func cloneLogger(l *Logger) *Logger {
	return &Logger{
		configuration: l.configuration,
		fields:        l.fields,
		syncHooks:     l.syncHooks,
		asyncHooks:    l.asyncHooks,
		outputs:       l.outputs,
	}
}
