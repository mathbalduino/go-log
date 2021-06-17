package loxeLog

type Hooks = map[string]func(Log) interface{}

func (l *Logger) SyncHooks(hooks Hooks) *Logger {
	newLogger := cloneLogger(l)
	mergeOverriding_(newLogger.syncHooks, hooks)
	return newLogger
}

func (l *Logger) RawSyncHooks(hooks Hooks) *Logger {
	newLogger := cloneLogger(l)
	newLogger.syncHooks = hooks
	return newLogger
}

func (l *Logger) AsyncHooks(hooks Hooks) *Logger {
	newLogger := cloneLogger(l)
	mergeOverriding_(newLogger.asyncHooks, hooks)
	return newLogger
}

func (l *Logger) RawAsyncHooks(hooks Hooks) *Logger {
	newLogger := cloneLogger(l)
	newLogger.asyncHooks = hooks
	return newLogger
}
