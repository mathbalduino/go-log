package loxeLog

type Configuration struct {
	AsyncScheduler AsyncScheduler
	LvlFieldName   string
	MsgFieldName   string
	LvlsEnabled    uint64
}

func DefaultConfig() Configuration {
	return Configuration{
		nil,
		"lvl",
		"msg",
		LvlDefaults,
	}
}
