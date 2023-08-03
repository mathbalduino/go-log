package logger

// Configuration is a struct that holds
// global src configurations
type Configuration struct {
	// AsyncScheduler will be used to distinguish
	// between a sync (if nil) or async (if not nil)
	// approach
	AsyncScheduler AsyncScheduler

	// LvlFieldName is used to customize the name of
	// the key that represents the level of the Log
	//
	// Note that it cannot be equal to the MsgFieldName
	LvlFieldName string

	// MsgFieldName is used to customize the name of
	// the key that represents the message of the Log
	//
	// Note that it cannot be equal to the LvlFieldName
	MsgFieldName string

	// LvlsEnabled is an integer that represents which
	// Log levels are enabled.
	//
	// Note that it is intended to be used as a combination
	// ("or" bitwise operation) of log levels
	LvlsEnabled uint64

	// ErrorParser is a function that takes an error and
	// return a msg string and an optional collection of fields
	// (used by ErrorFrom and FatalFrom methods)
	ErrorParser func(error) (string, LogFields)
}

func (l *logger) Configuration(c Configuration) { l.configuration = &c }

// DefaultConfig creates a default Logger configuration,
// with a synchronous approach (nil AsyncScheduler),
// omitting only Trace logs, using "lvl" and "msg"
// as LvlFieldName and MsgFieldName, respectively, and
// extracting only the error message via 'ErrorParser'
func DefaultConfig() Configuration {
	return Configuration{
		nil,
		"lvl",
		"msg",
		LvlDefaults,
		DefaultErrorParser}
}

// DefaultErrorParser will return a tuple containing the error string
// and the following map: { "error": err }
func DefaultErrorParser(err error) (string, LogFields) {
	return err.Error(), LogFields{DefaultErrorKey: err}
}

// DefaultErrorKey is the value of the key used
// to store the errors returned by the DefaultErrorParser
const DefaultErrorKey = "error"

// validateConfig will return a non-nil error
// if the given Configuration contains errors
func validateConfig(c Configuration) error {
	if c.LvlFieldName == c.MsgFieldName {
		return ErrLvlMsgSameKey
	}
	if c.ErrorParser == nil {
		return ErrNilErrorParser
	}
	return nil
}
