---
sidebar_position: 7
---

# Outputs

Every log needs to be written to somewhere, otherwise why would we create it? In this library, `Output` functions represent this final destiny:

```go
// outputs.go
// Just an alias to a simple function
type Output = func(lvl uint64, msg string, fields LogFields)
```

Inside these functions, you can do whatever you want. In general, it will be a write operation to the filesystem, database, or some request to the cloud.

:::note
Even with the `lvl` and `msg` being directly accessible via the function parameters, they're inside the `fields` map param too, so you can just parse the `LogFields` map right away.
:::

:::danger
You have to be very carefull when writing `Outputs`, because the library is not prepared to handle any kind of `panic` that can occur inside them, and it can cause issues. 

_**Except if it's the last `Output` (last thing to be executed), avoid `panic` calls inside `Outputs`**_
:::

## Defining Outputs

Just like `Hooks`, you can set a new `Output` using the `Outputs`/`RawOutputs` methods:

```go
// outputs.go
func Outputs(output Output, outputs ...Output) *Logger { ... }
func RawOutputs(output Output, outputs ...Output) *Logger { ... }
```

Note that it's a variadic function, so you can pass as many outputs as you want. The order will be preserved.

The `Outputs` will **append** the new outputs to the old ones, while the `RawOutputs` will ignore the old ones and use just the new outputs.

## Ordering

Since these functions are stored as an slice inside the `Logger`, the order can be preserved and you can use it at you benefit. You can save it first to the database A, and, in the next `Output`, read from database A, process, and save to database B, for example.

At the end of the life cycle of every created log, there's a for loop that will iterate over the `Outputs` slice:

```go
// Not real production-code (just to illustrate)
for _, output := range logger.outputs {
  output(lvl, msg, logFields)
}
```

You must handle possible `panic` calls that may occur inside the outputs, because it will be not handled by the library.

## Builtin outputs

There's 4 builtin outputs, ready to be used:

### OutputToWriter

Writes the log to some `io.Writer` (usually, a file), after being parsed using the `OutputParser`:

```go
// outputs.go
type OutputParser = func(LogFields) ([]byte, error) // just an alias
func OutputToWriter(w io.Writer, parser OutputParser, onError func(error)) Output {
  ...
}
```

Note that there's a thirty argument: `onError`. It is used to handle possible errors when trying to parse the log using the `OutputParser` or trying to write it to the `io.Writer`. It is intended to be used as a last fallback.

This is, in fact, a function that will return another function. Note that if you pass it directly to the `Logger`, the compiler will stop you. The returned function is the real `Output`, pay attention. Example:

```go
// compiler error
logger.NewDefault().
  Outputs(logger.OutputToWriter)

// ok
logger.NewDefault().
  Outputs(logger.OutputToWriter(w, p, func(error) {}))
```

### OutputJsonToWriter

Writes the log to some `io.Writer` (usually, a file), after being parsed to `json`:

```go
// outputs.go
func OutputJsonToWriter(w io.Writer, onError func(error)) Output {
  ...
}
```

Note that the `onError` argument has the same purpose as the one in [OutputToWriter](#outputtowriter).

This is, in fact, a function that will return another function. Note that if you pass it directly to the `Logger`, the compiler will stop you. The returned function is the real `Output`, pay attention.

### OutputAnsiToStdout

Writes the log to the `stdout`, displaying just the `level` and the `message`, using `ANSI` codes to colorize it accordingly to it's `level`. If your `stdout` don't have support for `ANSI` codes, don't use this `Output` (not common, since in general it will be some terminal).

```go
// outputs.go
func OutputAnsiToStdout(lvl uint64, msg string, _ LogFields) {
  ...
}
```

### OutputPanicOnFatal

As you will see, the `Fatal` log level doesn't do anything special. In order to unlock it's ability to `panic`, you will need to use this special `Output`:

```go
// outputs.go
func OutputPanicOnFatal(lvl uint64, msg string, fields LogFields) {
  ...
}
```

Just set it to be the last `Output` and it will call `panic` if the received log is a Fatal one. 

If there's some `error` value inside the `LogFields`, it will be given to the `panic` call, otherwise, a new `error` will be created using the `msg` argument and `fmt.Errorf()`.

This `Output` will search for the error value inside the log fields using the `DefaultErrorKey` key:

```go
// configuration.go
func DefaultErrorParser(err error) (string, LogFields) {
    return err.Error(), LogFields{DefaultErrorKey: err}
}
const DefaultErrorKey = "error"
```

```go
// outputs.go
func OutputPanicOnFatal(lvl uint64, msg string, fields LogFields) {
  ...
  err := fields[DefaultErrorKey]
  ...
}
```

If you're using a different `ErrorParser`, other than the [DefaultErrorParser](log_levels.md#default-errorparser), make sure that the error value is stored inside the `LogFields` under the `DefaultErrorKey` key. ~~Or not, it's up to you~~

## Writing your own outputs

If you're going to write your own `Outputs`, there's two functions that you need to know about: 

```go
func LvlToString(lvl uint64) string { ... }
func ColorizeStrByLvl(lvl uint64, msg string) string { ... }
```

These functions are used internally by some builtin outputs, but are exported, so you can use too. The first one (`LvlToString`) will take the log level `uint64` and return the string that represents it. The second one (`ColorizeStrByLvl`) will take the log level `uint64` and an arbitrary `string`, returning a new `string` wrapped with the `ANSI` code that is used to colorize that log level.

:::note
Remember that these functions will ignore any custom log levels, written outside the library
:::
