package loxeLog

type LogFields = map[string]interface{}

func (l *Logger) Fields(fields LogFields) *Logger {
	newLogger := cloneLogger(l)
	mergeOverriding(newLogger.fields, fields)
	return newLogger
}

func (l *Logger) RawFields(fields LogFields) *Logger {
	newLogger := cloneLogger(l)
	newLogger.fields = fields
	return newLogger
}
