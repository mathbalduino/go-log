---
sidebar_position: 12
---

# Configuration

You can customize `Logger` instances using the `Configuration` struct:

```go
// configuration.go
type Configuration struct {
	AsyncScheduler AsyncScheduler
	LvlFieldName   string
	MsgFieldName   string
	LvlsEnabled    uint64
	ErrorParser    func(error) (string, LogFields)
}
```

- [AsyncScheduler](async_logger.md#asyncscheduler): If nil, the Logger will be set to `sync`. If not nil, the `Logger` will be [async](async_logger.md)
- `LvlFieldName`: The `lvl` field is always required to be present. This configuration is used to control the `key` that represents the `level` of the log inside the `LogFields`. **If it's equal to the `MsgFieldName`, an `error` will be thrown**
- `MsgFieldName`: The `msg` field is always required to be present. This configuration is used to control the `key` that represents the `message` of the log inside the `LogFields`. **If it's equal to the `LvlFieldName`, an `error` will be thrown**
- `LvlsEnabled`: `uint64` used to check the created logs to see if they're enabled. The log levels are expected to be used as integer flags, using values that are equal to the power of two: 2^0 (1), 2^1 (2), 2^2 (4), 2^3 (8), ...
- [ErrorParser](log_levels.md#errorparser): function used to extract information from the errors given to `ErrorFrom` and `FatalFrom` methods. **If it's nil, an `error` will be thrown**

## LvlsEnabled usage

For every created log, the log `level` and the `Configuration.LvlsEnabled` will be compared using the bitwise `and` operator. If the result of the operation is `true`, the log is allowed to continue it's life cycle. If `false`, the log is not created and nothing happens (`noop`).

This is the real function that checks to see if the log `level` is enabled:

```go
// util.go
// Real production code
func notEnabled(flags uint64, logLvl uint64) bool {
	return (flags & logLvl) == 0
}
```

The library comes with some builtin constant values, that can be used to calculate new values, for example, or be used to configure the `Logger` in a different way:

```go
// logLevels.go
const (
	LvlTrace uint64 = 1 << iota
	LvlDebug
	LvlInfo
	LvlWarn
	LvlError
	LvlFatal
)
const (
	LvlProduction = LvlInfo | LvlWarn | LvlError | LvlFatal
	LvlDefaults = LvlDebug | LvlProduction
	LvlAll = LvlTrace | LvlDefaults
)
```

:::tip
If you need to convert the log level `uint64` to some string representation, you can use the `LvlToString` function, exported by the root package.
:::

## Default configuration

The library comes with a builtin `Default Configuration`, accessible by calling the `DefaultConfig()` root package function, that will handle the most basic scenarios:

```go
// configuration.go
func DefaultConfig() Configuration {
	return Configuration{
		AsyncScheduler: nil,
		LvlFieldName: "lvl",
		MsgFieldName: "msg",
		LvlsEnabled:  LvlDefaults,
		ErrorParser:  DefaultErrorParser,
  }
}
```

This configuration will set the `Logger` to be `sync` (nil `AsyncScheduler`), use the default values for the `level` and `message` required fields, disable only the `Trace` log levels and use the [DefaultErrorParser](log_levels.md#default-errorparser) to extract the errors given to the `ErrorFrom` and `FatalFrom` methods.

It is, in fact, the configuration used to create [Default Loggers](logger_creation.md).
