package loxeLog

type Log struct {
	lvl         uint64
	msg         string
	logger      *Logger
	syncFields  LogFields
	adHocFields []LogFields
	fields      LogFields
}

func (l Log) Field(key string) interface{} {
	v := tryRead(key, l.fields)
	if v != nil {
		return v
	}

	v = tryRead(key, l.adHocFields...)
	if v != nil {
		return v
	}

	v = tryRead(key, l.syncFields)
	if v != nil {
		return v
	}

	return tryRead(key, l.logger.fields)
}

// -----

func handleLog(log Log) {
	log.fields = cloneOrNew(log.logger.fields)
	mergeOverriding(log.fields, log.syncFields)
	mergeOverriding(log.fields, log.adHocFields...)
	applyHooks(log, log.fields, log.logger.asyncHooks)
	log.fields[log.logger.configuration.LvlFieldName] = log.lvl
	log.fields[log.logger.configuration.MsgFieldName] = log.msg

	for _, output := range log.logger.outputs {
		output(log.lvl, log.msg, log.fields)
	}
}
