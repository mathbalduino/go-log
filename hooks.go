package golog

// Hooks is just an alias to a map that represents
// dynamic fields of the log (values are evaluated
// by calling the function)
type Hooks = map[string]func(Log) interface{}

func (l *logger) PreHooks(hooks Hooks) Logger {
	newLogger := cloneLogger(l)
	mergeOverriding(newLogger.preHooks, hooks)
	return newLogger
}

func (l *logger) RawPreHooks(hooks Hooks) Logger {
	newLogger := cloneLogger(l)
	newLogger.preHooks = hooks
	return newLogger
}

func (l *logger) PostHooks(hooks Hooks) Logger {
	newLogger := cloneLogger(l)
	mergeOverriding(newLogger.postHooks, hooks)
	return newLogger
}

func (l *logger) RawPostHooks(hooks Hooks) Logger {
	newLogger := cloneLogger(l)
	newLogger.postHooks = hooks
	return newLogger
}
