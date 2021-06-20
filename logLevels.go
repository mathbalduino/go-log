package loxeLog

// Trace will create a new log with the Trace level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is a variadic to simulate
// optional params. If it's needed, set just the first variadic
// value (all the subsequent ones will be ignored)
func (l *Logger) Trace(msg string, adHocFields ...LogFields) { l.Log(LvlTrace, msg, adHocFields) }

// Debug will create a new log with the Debug level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is a variadic to simulate
// optional params. If it's needed, set just the first variadic
// value (all the subsequent ones will be ignored)
func (l *Logger) Debug(msg string, adHocFields ...LogFields) { l.Log(LvlDebug, msg, adHocFields) }

// Info will create a new log with the Info level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is a variadic to simulate
// optional params. If it's needed, set just the first variadic
// value (all the subsequent ones will be ignored)
func (l *Logger) Info(msg string, adHocFields ...LogFields) { l.Log(LvlInfo, msg, adHocFields) }

// Warn will create a new log with the Warn level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is a variadic to simulate
// optional params. If it's needed, set just the first variadic
// value (all the subsequent ones will be ignored)
func (l *Logger) Warn(msg string, adHocFields ...LogFields) { l.Log(LvlWarn, msg, adHocFields) }

// Error will create a new log with the Error level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is a variadic to simulate
// optional params. If it's needed, set just the first variadic
// value (all the subsequent ones will be ignored)
func (l *Logger) Error(msg string, adHocFields ...LogFields) { l.Log(LvlError, msg, adHocFields) }

// Fatal will create a new log with the Fatal level, using the
// given adHocFields, if enabled.
//
// Note that the 'adHocFields' param is a variadic to simulate
// optional params. If it's needed, set just the first variadic
// value (all the subsequent ones will be ignored)
func (l *Logger) Fatal(msg string, adHocFields ...LogFields) {
	l.Log(LvlFatal, msg, adHocFields)
	if !notEnabled(l.configuration.LvlsEnabled, LvlFatal) {
		panic(msg)
	}
}

// Log is the base method that creates a new log, being used by all
// other log methods (Trace, Debug, Warn, ...). If there's the need
// to create custom log levels, use this method.
//
// Note that the 'adHocFields' param is implemented as a variadic by
// all the other log methods (Trace, Debug, Warn, ...), so the 'adHocFields'
// param is an slice. Note that this method will use just the first index of
// this slice, ignoring the subsequent ones
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
	if len(l.syncHooks) > 0 {
		log.syncFields = LogFields{}
		applyHooks(log, log.syncFields, l.syncHooks)
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
