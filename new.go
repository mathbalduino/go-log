package loxeLog

// Logger is a struct that represents
// a base log template, that contains
// basic information on how to handle
// new logs
type Logger struct {
	configuration *Configuration
	fields        LogFields
	preHooks      Hooks
	postHooks     Hooks
	outputs       []Output
}

// New creates a new Logger instance, validating
// the given configuration
//
// Note that if the given configuration is invalid,
// panic will be called with the proper errors
func New(config Configuration) *Logger {
	e := validateConfig(config)
	if e != nil {
		panic(e)
	}

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
		RawOutputs(OutputToAnsiStdout)
}
