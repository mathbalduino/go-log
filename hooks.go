package loxeLog

// Hooks is just an alias to a map that contains
// information about fields keys/value (via function
// call)
type Hooks = map[string]func(Log) interface{}

// SyncHooks will append the given Hooks to the Logger
// sync Hooks, overriding any already existing hooks
// and returning a new Logger instance
func (l *Logger) SyncHooks(hooks Hooks) *Logger {
	newLogger := cloneLogger(l)
	mergeOverriding_(newLogger.syncHooks, hooks)
	return newLogger
}

// RawSyncHooks will set the given Hooks to the Logger
// sync Hooks, discarding the previous value and returning
// a new Logger instance
func (l *Logger) RawSyncHooks(hooks Hooks) *Logger {
	newLogger := cloneLogger(l)
	newLogger.syncHooks = hooks
	return newLogger
}

// AsyncHooks will append the given Hooks to the Logger
// async Hooks, overriding any already existing hooks
// and returning a new Logger instance
func (l *Logger) AsyncHooks(hooks Hooks) *Logger {
	newLogger := cloneLogger(l)
	mergeOverriding_(newLogger.asyncHooks, hooks)
	return newLogger
}

// RawAsyncHooks will set the given Hooks to the Logger
// async Hooks, discarding the previous value and returning
// a new Logger instance
func (l *Logger) RawAsyncHooks(hooks Hooks) *Logger {
	newLogger := cloneLogger(l)
	newLogger.asyncHooks = hooks
	return newLogger
}
