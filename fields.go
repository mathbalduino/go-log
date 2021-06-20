package loxeLog

// LogFields is just an alias to a map that
// represents the fields of the log
type LogFields = map[string]interface{}

// Fields will append the given LogFields to the Logger
// fields, overriding any already existing fields and
// returning a new Logger instance
func (l *Logger) Fields(fields LogFields) *Logger {
	newLogger := cloneLogger(l)
	mergeOverriding(newLogger.fields, fields)
	return newLogger
}

// RawFields will set the given LogFields directly to the
// Logger fields, discarding the previous value and
// returning a new Logger instance
func (l *Logger) RawFields(fields LogFields) *Logger {
	newLogger := cloneLogger(l)
	newLogger.fields = fields
	return newLogger
}
