package logger

// LogFields is just an alias to a map that
// represents the fields of the log
type LogFields = map[string]interface{}

func (l *logger) Field(key string) interface{} {
	return l.fields[key]
}

func (l *logger) Fields(fields LogFields) Logger {
	newLogger := cloneLogger(l)
	mergeOverriding(newLogger.fields, fields)
	return newLogger
}

func (l *logger) RawFields(fields LogFields) Logger {
	newLogger := cloneLogger(l)
	newLogger.fields = fields
	return newLogger
}
