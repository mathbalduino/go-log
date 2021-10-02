---
sidebar_position: 2
---

# Base fields

You can customize the `Base Fields` of the `Logger` instance using the following methods:

```go
// fields.go
func (l *Logger) Fields(fields LogFields) *Logger { ... }
func (l *Logger) RawFields(fields LogFields) *Logger { ... }
```

:::note
The returned `Logger` instance will be a copy of the previous one, sharing the same `Configuration` struct. The only difference will be the `Base Fields`. 

More info at the [Loggers clonage](../advanced/loggers_clonage.md) page.
:::

There's a third method, used to query `Base Fields` of the `Logger` instance:

```go
// fields.go
func (l *Logger) Field(key string) interface{} { ... }
```

## Fields method

This method will return a new copy of the `Logger` instance, with the given `LogFields` applied. Note that if the previous `Logger` had a `Base Field` with a key that clashes with one of the new ones, it will be overriden:

```go
firstLogger := logger.NewDefault().
  Fields(logger.LogFields{
    "field-A": "value-A",
    "field-B": "value-B",
    "field-C": "value-C",
  })
firstLogger.Trace("first log")
/*
  {
    "msg": "first log",
    "lvl": 1,
    "field-A": "value-A",
    "field-B": "value-B",
    "field-C": "value-C"
  }
*/

secondLogger := firstLogger.
  Fields(logger.LogFields{
    "field-B": "new value",
  })
secondLogger.Info("second log")
/*
  {
    "msg": "second log",
    "lvl": 4,
    "field-A": "value-A",
    "field-B": "new value",
    "field-C": "value-C"
  }
*/
```

## RawFields method

This method is almost equal to the previous one (returns a new copy of the `Logger` instance too), with one difference: the given `LogFields` will be set right away, ignoring any previous `Base Fields`.

```go
firstLogger := logger.NewDefault().
  Fields(logger.LogFields{
    "field-A": "value-A",
    "field-B": "value-B",
    "field-C": "value-C",
  })
firstLogger.Trace("first log")
/*
  {
    "msg": "first log",
    "lvl": 1,
    "field-A": "value-A",
    "field-B": "value-B",
    "field-C": "value-C"
  }
*/

secondLogger := firstLogger.
  RawFields(logger.LogFields{
    "field-B": "new value",
  })
secondLogger.Info("second log")
/*
  {
    "msg": "second log",
    "lvl": 4,
    "field-B": "new value",
  }
*/
```

## Field method

This helper method can be used to retrieve the values of the `Base fields` (and just `Base fields`):

```go
someLogger := logger.NewDefault().
  Fields(logger.LogFields{
    "field": "value",
  })
v := someLogger.Field("field")
fmt.Println(v)
// "value"
```

:::info
It's impossible to query `PreHooks`, `AdHoc fields` or `PostHooks` directly from the `Logger`, because they're not ready yet.
:::

:::tip
You can use this method to create more complex rules when cloning `Loggers`. See the implementation of [LoggerCLI](https://github.com/mathbalduino/go-log/blob/79263810d94dd1f2d112727824d1c5256b27951b/loggerCLI/nestLogger.go#L12) for a concrete example.
:::


## Dynamic Fields

You may have noticed that if you need to calculate the value of some field every time a new log is created (think about calculating the timestamp of your logs, for example), the `Base Fields` are useless, since the `Fields`/`RawFields` methods can only set fields with constant values.

To set dynamic fields, you will need to use `Hooks`. They will be discussed in the following pages.
