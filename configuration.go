package loxeLog

// Configuration is a struct that holds
// global logger configurations
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
}

// DefaultConfig creates a default Logger configuration,
// with a synchronous approach (nil AsyncScheduler),
// omitting only Trace logs and using "lvl" and "msg"
// as LvlFieldName and MsgFieldName, respectively
func DefaultConfig() Configuration {
	return Configuration{
		nil,
		"lvl",
		"msg",
		LvlDefaults,
	}
}

// validateConfig will return a non-nil error
// if the given Configuration contains errors
func validateConfig(c Configuration) error {
	if c.LvlFieldName == c.MsgFieldName {
		return ErrLvlMsgSameKey
	}
	return nil
}
