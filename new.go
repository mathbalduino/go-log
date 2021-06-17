package loxeLog

type Logger struct {
	configuration *Configuration
	fields        LogFields
	syncHooks     Hooks
	asyncHooks    Hooks
	outputs       []Output
}

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

func NewDefault() *Logger {
	return New(DefaultConfig()).
		RawOutputs(OutputToAnsiTerm)
}
