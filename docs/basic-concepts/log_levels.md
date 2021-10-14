---
sidebar_position: 9
---

# Log levels

## Basic log levels

There are 6 basic log levels implemented by default:

```go
// logLevels.go
func Trace(msg string, adHocFields ...LogFields) { ... }
func Debug(msg string, adHocFields ...LogFields) { ... }
func Info(msg string, adHocFields ...LogFields) { ... }
func Warn(msg string, adHocFields ...LogFields) { ... }
func Error(msg string, adHocFields ...LogFields) { ... }
func Fatal(msg string, adHocFields ...LogFields) { ... }
```

In order to call these methods, you will need a valid `Logger` instance. See [Logger creation](logger_creation.md) for details.

The only difference between them is the value of the `lvl` log field. Note that even if it's common to `panic` when calling `Fatal`, you will have to implement some `Output` to simulate this behaviour (see [OutputPanicOnFatal](outputs.md#outputpaniconfatal) for details).

:::note
The `lvl` of the log is represented using an `uint64` and values that correspond to the power of two (1, 2, 4, 8, 16, ...). This way, it's possible to apply the `and` bitwise operation and easily check to see if some log level X is enabled. For details, see [LevelsEnabled usage](configuration.md#lvlsenabled-usage)
:::

## Extra Log Levels

If you've read the [Introduction](../intro.md), you may have noticed these two log levels called `ErrorFrom` and `FatalFrom`.

These two methods are another way to call `Error` and `Fatal`, respectively, but using an `error` interface directly.

### ErrorParser

If you're using a concise `error` handling strategy, it should be easy to manipulate errors and
extract more information from them, beyond the error string.

The `Configuration` struct accepts a function called `ErrorParser`, that will take some `error` interface and return a tuple containing the string that better represents the error, and some log fields extracted from the error itself:

```go
// configuration.go
type Configuration struct {
  ...
  ErrorParser func(error) (string, LogFields)
  ...
}
```

This function will be called from both `ErrorFrom` and `FatalFrom` methods, in order to extract more information about the error, to finally call the correspondent log method (`Error` and`Fatal`, respectively).

Instead of doing something like this:

```go
type SomeErrorStruct struct{
  Value int
}
func (s *SomeErrorStruct) Error() string { return "error msg" }

func maybeReturnsError() error { return &SomeErrorStruct{Value: 10} }

func main() {
  someLogger := logger.NewDefault()

  e := maybeReturnsError()
  if e != nil {
    err := e.(*SomeErrorStruct)
    someLogger.Fatal(e.Error(), logger.LogFields{ "errorValue": err.Value })
  }
}
```

You can do this:
```go
type SomeErrorStruct struct{
  Value int
}
func (s *SomeErrorStruct) Error() string { return "error msg" }

func maybeReturnsError() error { return &SomeErrorStruct{Value: 10} }

func MyErrorParser(e error) (string, logger.LogFields) {
  return e.Error(), logger.LogFields{
	  "errorValue": e.(*SomeErrorStruct).Value,
  }
}

func main() {
  config := logger.DefaultConfig()
  config.ErrorParser = MyErrorParser
  someLogger := logger.New(config).
    Outputs(logger.OutputJsonToWriter(os.Stdout, nil))

  e := maybeReturnsError()
  if e != nil {
    someLogger.FatalFrom(e)
    // { "errorValue": 10,"lvl": 32, "msg": "error msg" }
  }
}
```

Note that the fields returned by the `ErrorParser` will be placed **before** the `AdHoc fields`, causing them to be overwritten by the `AdHoc fields` if there's a clash of keys. Example:

```go
// Suppose that the ErrorParser is set to:
func MyErrorParser(_ error) (string, logger.LogFields) {
  return "some example error msg", logger.LogFields{ 
    "someKey": "value", 
    "anotherKey": "another value",
  }
}

func main() {
  config := logger.DefaultConfig()
  config.ErrorParser = MyErrorParser
  someLogger := logger.New(config).
    Outputs(logger.OutputJsonToWriter(os.Stdout, nil))

  // And you create the following log:
  someLogger.ErrorFrom(fmt.Errorf("any error"), logger.LogFields{
    "someKey": "newValue",
    "thirtyKey": "another another value",
  })
  /*
    {
      "msg": "some example error msg",
      "lvl": 16,
      "someKey": "newValue",
      "anotherKey": "another value",
      "thirtyKey": "another another value"
    }
  */
}
```

### Default ErrorParser

The library has a builtin `ErrorParser`, that will return the `e.Error()` as the msg and store the `error` interface value inside the log fields, using the `DefaultErrorKey` as the field key:

```go
// configuration.go
// Production code-fragment
func DefaultErrorParser(err error) (string, LogFields) {
	return err.Error(), LogFields{DefaultErrorKey: err}
}
const DefaultErrorKey = "error"
```

This logic is used to build the [OutputPanicOnFatal](outputs.md#outputpaniconfatal), that extracts the error value (under the `DefaultErrorKey` key) to forward it to the `panic` call.

## Standalone log levels

If you don't want to create and maintain a `Logger` instance by yourself, you can just use the standalone version of the log levels:

```go
// standalone.go
func Trace(msg string, adHocFields ...LogFields) { ... }
func Debug(msg string, adHocFields ...LogFields) { ... }
func Info(msg string, adHocFields ...LogFields) { ... }
func Warn(msg string, adHocFields ...LogFields) { ... }
func Error(msg string, adHocFields ...LogFields) { ... }
func Fatal(msg string, adHocFields ...LogFields) { ... }
```

Under the hood, they will use a default `Logger` instance. Something like this:

```go
func Trace(msg string, adHocFields ...LogFields) {
  NewDefault().Trace(msg, adHocFields...)
}
```

For more information, see [Default Logger](logger_creation.md).

## Extending the log levels

If you want to create new custom log levels, you will need to create a new type that wrap the original `Logger`.

The most straightway to do it, is to use embedded struct fields. Something like this:

```go
type NewLogger struct {
  *logger.Logger
}
```

This way, the `NewLogger` will preserve the original log level methods and you can add new ones. Every log level should call the base log level method, that will actually handle the log. This method is exported by the `Logger` api:

```go
// logLevels.go
func (l *Logger) Log(lvl uint64, msg string, adHocFields []LogFields) { ... }
```

You can notice that the `adHoc` fields are represented as an slice, while the log level methods (`Fatal`, `Info`, etc) implement it as a variadic. Well, variadic arguments are just a slice, so if you want to create custom log levels using variadics too, you can just forward it to the `Log` method:

```go
type NewLogger struct {
  *logger.Logger
}

func (l *NewLogger) NewLogLevel(msg string, adHocFields ...logger.LogFields) {
  l.Logger.Log(<newLogLevelUint64>, msg, adHocFields)
}
```

Remember that it's not recommended to overlap new, custom log levels values (`uint64`), with the builtin ones. They're all exported, so you can just check them or do something like this:

```go
const (
  newLogLevel uint64 = 1 << (iota + 6)
  anotherLogLevel
  anotherOne
)
```

Note that it will continue the sequence after the last builtin log level, `LvlFatal`. There's six builtin log levels (Trace, Debug, Info, Warn, Error and Fatal), that's why there's a `6` being added to the `iota`.

If you want to omit some builtin log levels, just don't use embedded struct fields:

```go
type NewLogger struct {
  original *logger.Logger
}

func (l *NewLogger) Info(msg string, adHocFields ...logger.LogFields) {
  l.original.Info(msg, adHocFields...)
}
func (l *NewLogger) Error(msg string, adHocFields ...logger.LogFields) {
  l.original.Error(msg, adHocFields...)
}
```

This way, there will be only two methods in the new custom `Logger` api: `Info` and `Error`.

You can see a more concrete example by looking at the `LoggerCLI`, [here](https://github.com/mathbalduino/go-log/blob/main/loggerCLI/logLevels.go) and [here](https://github.com/mathbalduino/go-log/blob/79263810d94dd1f2d112727824d1c5256b27951b/loggerCLI/new.go#L9).
