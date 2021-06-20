package loxeLog

// Logger is a struct that represents
// a base log template, that contains
// basic information to create/persist
// new logs
type Logger struct {
	configuration *Configuration
	fields        LogFields
	syncHooks     Hooks
	asyncHooks    Hooks
	outputs       []Output
}

// New creates a new Logger instance, validating
// the given configuration
//
// Note that if the given configuration is invalid,
// panic will be called
func New(config Configuration) *Logger {
	// TODO: validate config: 1) lvl and msg field names cannot be equal. 2) accept only ascii chars

	return &Logger{
		&config,
		nil,
		nil,
		nil,
		nil,
	}
}

// NewDefault will create a basic sync Logger instance
// that outputs newly created logs to the terminal (using
// ANSI codes), using 'lvl' and 'msg' as level and message
// keys and enabling only the default log levels
func NewDefault() *Logger {
	return New(DefaultConfig()).
		RawOutputs(OutputToAnsiTerm)
}
