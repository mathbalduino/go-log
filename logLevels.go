package logger

func (l *logger) Trace(msg string, adHocFields ...LogFields) { l.Log(LvlTrace, msg, adHocFields) }
func (l *logger) Debug(msg string, adHocFields ...LogFields) { l.Log(LvlDebug, msg, adHocFields) }
func (l *logger) Info(msg string, adHocFields ...LogFields)  { l.Log(LvlInfo, msg, adHocFields) }
func (l *logger) Warn(msg string, adHocFields ...LogFields)  { l.Log(LvlWarn, msg, adHocFields) }
func (l *logger) Error(msg string, adHocFields ...LogFields) { l.Log(LvlError, msg, adHocFields) }
func (l *logger) Fatal(msg string, adHocFields ...LogFields) { l.Log(LvlFatal, msg, adHocFields) }

func (l *logger) ErrorFrom(e error, adHocFields ...LogFields) {
	msg, f := l.configuration.ErrorParser(e)
	if f != nil {
		l.Error(msg, append([]LogFields{f}, adHocFields...)...)
	} else {
		l.Error(msg, adHocFields...)
	}
}

func (l *logger) FatalFrom(e error, adHocFields ...LogFields) {
	msg, f := l.configuration.ErrorParser(e)
	if f != nil {
		l.Fatal(msg, append([]LogFields{f}, adHocFields...)...)
	} else {
		l.Fatal(msg, adHocFields...)
	}
}

func (l *logger) Log(lvl uint64, msg string, adHocFields []LogFields) {
	if notEnabled(l.configuration.LvlsEnabled, lvl) {
		return
	}

	log := Log{lvl, msg, l, nil, adHocFields, nil}
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
	// (excluding only Debug and Trace)
	LvlProduction = LvlInfo | LvlWarn | LvlError | LvlFatal

	// LvlDefaults will enable Debug and all LvlProduction logs
	// (excluding only Trace)
	LvlDefaults = LvlDebug | LvlProduction

	// LvlAll will enable all logs
	LvlAll = LvlTrace | LvlDefaults
)
