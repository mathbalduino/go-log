package logger

// Hooks is just an alias to a map that contains
// information about fields keys/value (via function
// call)
type Hooks = map[string]func(Log) interface{}

// PreHooks will append the given Hooks to the Logger
// pre Hooks, overriding any already existing hooks
// and returning a new Logger instance
func (l *Logger) PreHooks(hooks Hooks) *Logger {
	newLogger := cloneLogger(l)
	mergeOverriding_(newLogger.preHooks, hooks)
	return newLogger
}

// RawPreHooks will set the given Hooks to the Logger
// pre Hooks, discarding the previous value and returning
// a new Logger instance
func (l *Logger) RawPreHooks(hooks Hooks) *Logger {
	newLogger := cloneLogger(l)
	newLogger.preHooks = hooks
	return newLogger
}

// PostHooks will append the given Hooks to the Logger
// post Hooks, overriding any already existing hooks
// and returning a new Logger instance
func (l *Logger) PostHooks(hooks Hooks) *Logger {
	newLogger := cloneLogger(l)
	mergeOverriding_(newLogger.postHooks, hooks)
	return newLogger
}

// RawPostHooks will set the given Hooks to the Logger
// post Hooks, discarding the previous value and returning
// a new Logger instance
func (l *Logger) RawPostHooks(hooks Hooks) *Logger {
	newLogger := cloneLogger(l)
	newLogger.postHooks = hooks
	return newLogger
}
