package logger

type Logger interface {
	// Field acts as a getter for base fields.
	// Note that this method will ignore preHooks/postHooks
	Field(key string) interface{}

	// Fields will append the given LogFields to the Logger
	// base fields, overriding old base fields (if there's a
	// key clash) and returning a new Logger instance
	Fields(fields LogFields) Logger

	// RawFields will set the given LogFields directly to the
	// Logger base fields, discarding all the previous base
	// fields and returning a new Logger instance
	RawFields(fields LogFields) Logger

	// PreHooks will append the given Hooks to the Logger
	// preHooks, overriding old preHooks (if there's a key
	// clash) and returning a new Logger instance
	PreHooks(hooks Hooks) Logger

	// RawPreHooks will set the given Hooks to the Logger
	// preHooks, discarding all the previous preHooks and
	// returning a new Logger instance
	RawPreHooks(hooks Hooks) Logger

	// PostHooks will append the given Hooks to the Logger
	// postHooks, overriding old postHooks (if there's a key
	// clash) and returning a new Logger instance
	PostHooks(hooks Hooks) Logger

	// RawPostHooks will set the given Hooks to the Logger
	// postHooks, discarding all the previous postHooks and
	// returning a new Logger instance
	RawPostHooks(hooks Hooks) Logger

	// Outputs will append the given Output params (including
	// the variadic ones), in the given order, to the Logger
	// outputs and return a new Logger instance
	Outputs(output Output, outputs ...Output) Logger

	// RawOutputs will set the given Output params (including
	// the variadic ones), in the given order, to the Logger
	// outputs, discarding all the old values and returning a
	// new Logger instance
	RawOutputs(output Output, outputs ...Output) Logger

	// Trace will create a new log with the Trace level, if enabled,
	// applying the given adHocFields.
	Trace(msg string, adHocFields ...LogFields)

	// Debug will create a new log with the Debug level, if enabled,
	// applying the given adHocFields.
	Debug(msg string, adHocFields ...LogFields)

	// Info will create a new log with the Info level, if enabled,
	// applying the given adHocFields.
	Info(msg string, adHocFields ...LogFields)

	// Warn will create a new log with the Warn level, if enabled,
	// applying the given adHocFields.
	Warn(msg string, adHocFields ...LogFields)

	// Error will create a new log with the Error level, if enabled,
	// applying the given adHocFields.
	Error(msg string, adHocFields ...LogFields)

	// Fatal will create a new log with the Fatal level, if enabled,
	// applying the given adHocFields.
	//
	// Remember that this method will not call "panic"
	Fatal(msg string, adHocFields ...LogFields)

	// ErrorFrom will create a new log with the Error level, if enabled,
	// using the given error and adHocFields.
	ErrorFrom(e error, adHocFields ...LogFields)

	// FatalFrom will create a new log with the Fatal level, if enabled,
	// using the given error and adHocFields.
	//
	// Remember that this method will not call "panic"
	FatalFrom(e error, adHocFields ...LogFields)

	// Log is the base method that handles new logs creation, being used
	// by all other log methods (Trace, Debug, Warn, ...). If there's the
	// need to create custom log levels, you must call this method.
	Log(lvl uint64, msg string, adHocFields []LogFields)
}

// logger is a struct that represents
// a base log template, that contains
// basic information on how to handle
// new logs
type logger struct {
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
func New(config Configuration) Logger {
	e := validateConfig(config)
	if e != nil {
		panic(e)
	}

	return &logger{
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
func NewDefault() Logger {
	return New(DefaultConfig()).
		RawOutputs(OutputAnsiToStdout, OutputPanicOnFatal)
}
