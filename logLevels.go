package logger

// Trace will create a new log with the Trace level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
func (l *Logger) Trace(msg string, adHocFields ...LogFields) { l.Log(LvlTrace, msg, adHocFields) }

// Debug will create a new log with the Debug level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
func (l *Logger) Debug(msg string, adHocFields ...LogFields) { l.Log(LvlDebug, msg, adHocFields) }

// Info will create a new log with the Info level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
func (l *Logger) Info(msg string, adHocFields ...LogFields) { l.Log(LvlInfo, msg, adHocFields) }

// Warn will create a new log with the Warn level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
func (l *Logger) Warn(msg string, adHocFields ...LogFields) { l.Log(LvlWarn, msg, adHocFields) }

// Error will create a new log with the Error level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
func (l *Logger) Error(msg string, adHocFields ...LogFields) { l.Log(LvlError, msg, adHocFields) }

// Fatal will create a new log with the Fatal level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
func (l *Logger) Fatal(msg string, adHocFields ...LogFields) { l.Log(LvlFatal, msg, adHocFields) }

// ErrorFrom will create a new log with the Error level, using the
// given error and adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
//
// Note that this method will call the 'ErrorParser' configured
// function to extract the log message and custom fields from the
// given error. These custom fields will be overridden by the
// 'adHocFields' param
func (l *Logger) ErrorFrom(e error, adHocFields ...LogFields) {
	msg, f := l.configuration.ErrorParser(e)
	if f != nil {
		l.Error(msg, append([]LogFields{f}, adHocFields...)...)
	} else {
		l.Error(msg, adHocFields...)
	}
}

// FatalFrom will create a new log with the Fatal level, using the
// given error and adHocFields, if enabled.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
//
// Note that this method will call the 'ErrorParser' configured
// function to extract the log message and custom fields from the
// given error. These custom fields will be overridden by the
// 'adHocFields' param
func (l *Logger) FatalFrom(e error, adHocFields ...LogFields) {
	msg, f := l.configuration.ErrorParser(e)
	if f != nil {
		l.Fatal(msg, append([]LogFields{f}, adHocFields...)...)
	} else {
		l.Fatal(msg, adHocFields...)
	}
}

// Log is the base method that creates a new log, being used by all
// other log methods (Trace, Debug, Warn, ...). If there's the need
// to create custom log levels, use this method.
//
// Note that the 'adHocFields' param is variadic just to simulate
// optional params. Latter values will override former ones
func (l *Logger) Log(lvl uint64, msg string, adHocFields []LogFields) {
	if notEnabled(l.configuration.LvlsEnabled, lvl) {
		return
	}

	log := Log{
		lvl,
		msg,
		l,
		nil,
		adHocFields,
		nil,
	}
	if len(l.preHooks) > 0 {
		log.preFields = LogFields{}
		applyHooks(log, log.preFields, l.preHooks)
	}

	if l.configuration.AsyncScheduler != nil {
		l.configuration.AsyncScheduler.NextChannel() <- log
	} else {
		handleLog(log)
	}
}

const (
	// LvlTrace is a flag that if used will enable
	// Trace logs (via Logger Trace method)
	LvlTrace uint64 = 1 << iota

	// LvlDebug is a flag that if used will enable
	// Debug logs (via Logger Debug method)
	LvlDebug

	// LvlInfo is a flag that if used will enable
	// Info logs (via Logger Info method)
	LvlInfo

	// LvlWarn is a flag that if used will enable
	// Warn logs (via Logger Warn method)
	LvlWarn

	// LvlError is a flag that if used will enable
	// Error logs (via Logger Error method)
	LvlError

	// LvlFatal is a flag that if used will enable
	// Fatal logs (via Logger Fatal method)
	LvlFatal
)

const (
	// LvlProduction will enable Info, Warn, Error and Fatal logs
	LvlProduction = LvlInfo | LvlWarn | LvlError | LvlFatal

	// LvlDefaults will enable Debug and all LvlProduction logs
	LvlDefaults = LvlDebug | LvlProduction

	// LvlAll will enable all logs
	LvlAll = LvlTrace | LvlDefaults
)
