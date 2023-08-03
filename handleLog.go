package golog

// Log is a struct that represents
// a newly created log
//
// It is intended to be internal only,
// exported just to allow outside type
// reference
type Log struct {
	lvl         uint64
	msg         string
	logger      *logger
	preFields   LogFields
	adHocFields []LogFields
	postFields  LogFields
}

// Field will return the value associated with the
// given key
//
// Note that if this method is called in pre hooks,
// the adHoc and post fields aren't ready yet, so the
// returned value may be overridden by them.
// If called in post hook N, it can only see the result
// of the application of the N-1 previous post hooks. The
// returned value can be overridden by the N+1 next post
// hooks
//
// Pay attention
func (l Log) Field(key string) interface{} {
	if key == l.logger.configuration.MsgFieldName {
		return l.msg
	}
	if key == l.logger.configuration.LvlFieldName {
		return l.lvl
	}

	v := tryRead(key, l.postFields)
	if v != nil {
		return v
	}

	v = tryRead(key, l.adHocFields...)
	if v != nil {
		return v
	}

	v = tryRead(key, l.preFields)
	if v != nil {
		return v
	}

	return tryRead(key, l.logger.fields)
}

// -----

// handleLog will compile all the log fields and call
// the registered output functions.
//
// The src fields will be overridden by pre fields,
// that will be overridden by adHoc fields and latter by
// post fields
//
// Note that there are two reserved fields (lvl and msg),
// that will override any existing fields with the same
// configured keys
//
// Using var just to ease tests
var handleLog = func(log Log) {
	defer func() {
		if log.logger.configuration.AsyncScheduler == nil {
			return
		}
		// Ignore any panics that may happen here, if it is async
		recover()
	}()

	logFields := cloneOrNew(log.logger.fields)
	mergeOverriding(logFields, log.preFields)
	mergeOverriding(logFields, log.adHocFields...)
	if len(log.logger.postHooks) != 0 {
		log.postFields = LogFields{}
		applyHooks(log, log.postFields, log.logger.postHooks)
		mergeOverriding(logFields, log.postFields)
	}
	logFields[log.logger.configuration.LvlFieldName] = log.lvl
	logFields[log.logger.configuration.MsgFieldName] = log.msg

	for _, output := range log.logger.outputs {
		output(log.lvl, log.msg, logFields)
	}
}
