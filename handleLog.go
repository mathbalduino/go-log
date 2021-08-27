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
}

// Field will return the value associated with the
// given key
//
// Note that if this method is called in sync hooks,
// the async fields aren't ready yet, so the returned
// value may be overridden by them. Pay attention
func (l Log) Field(key string) interface{} {
	v := tryRead(key, l.adHocFields...)
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
// that will be overridden by async fields and latter by
// adHoc fields
//
// Note that there are two reserved fields (lvl and msg),
// that will override any existing fields with the same
// configured keys
//
// Using var just to ease tests
var handleLog = func(log Log) {
	logFields := cloneOrNew(log.logger.fields)
	mergeOverriding(logFields, log.syncFields)
	applyHooks(log, logFields, log.logger.asyncHooks)
	mergeOverriding(logFields, log.adHocFields...)
	logFields[log.logger.configuration.LvlFieldName] = log.lvl
	logFields[log.logger.configuration.MsgFieldName] = log.msg

	for _, output := range log.logger.outputs {
		output(log.lvl, log.msg, logFields)
	}
	if log.lvl == LvlFatal {
		panic(log.msg)
	}
}
