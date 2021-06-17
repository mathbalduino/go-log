package loxeLog

func (l *Logger) Trace(msg string, adHocFields ...LogFields) { l.Log(LvlTrace, msg, adHocFields) }
func (l *Logger) Debug(msg string, adHocFields ...LogFields) { l.Log(LvlDebug, msg, adHocFields) }
func (l *Logger) Info(msg string, adHocFields ...LogFields)  { l.Log(LvlInfo, msg, adHocFields) }
func (l *Logger) Warn(msg string, adHocFields ...LogFields)  { l.Log(LvlWarn, msg, adHocFields) }
func (l *Logger) Error(msg string, adHocFields ...LogFields) { l.Log(LvlError, msg, adHocFields) }
func (l *Logger) Fatal(msg string, adHocFields ...LogFields) {
	l.Log(LvlFatal, msg, adHocFields)
	if !notEnabled(l.configuration.LvlsEnabled, LvlFatal) {
		panic(msg)
	}
}

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
	LvlTrace uint64 = 1 << iota
	LvlDebug
	LvlInfo
	LvlWarn
	LvlError
	LvlFatal
)

const (
	LvlProduction = LvlInfo | LvlWarn | LvlError | LvlFatal
	LvlDefaults   = LvlDebug | LvlProduction
	LvlAll        = LvlTrace | LvlDefaults
)
