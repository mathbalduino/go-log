package loxeLog

type Logger struct {
	configuration *Configuration
	fields        LogFields
	syncHooks     Hooks
	asyncHooks    Hooks
	outputs       []Output
}

func New(config Configuration) *Logger {
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
