package loxeLog

// Log is a struct that represents
// a newly created log
//
// It is intended to be internal only,
// exported just to allow outside type
// reference
type Log struct {
	lvl         uint64
	msg         string
	logger      *Logger
	syncFields  LogFields
	adHocFields []LogFields
	fields      LogFields
}

// Field will return the value associated with the
// given key
//
// Note that if this method is called in sync hooks,
// the async fields aren't ready yet. Pay attention
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

// handleLog will compile all the log fields and call
// the registered output functions.
//
// The logger fields will be overridden by sync fields,
// that will be overridden by adhoc fields and latter by
// async fields
//
// Note that there are two reserved fields (lvl and msg),
// that will override any existing fields with the same
// configured keys
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
	if log.lvl == LvlFatal {
		panic(log.msg)
	}
}
