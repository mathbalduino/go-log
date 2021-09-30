---
sidebar_position: 4
---

# AdHoc fields

`AdHoc fields` will be applied at the [Thirty phase](life_cycle.md#async-phase-3-post-handling) of the life cycle, right after the `PreHooks` are executed.

Every Log level method can accept `AdHoc fields` (even custom ones), because all log levels must call `Logger.Log()` in the end (that accepts `AdHoc fields`):

```go
func (l *Logger) Log(lvl uint64, msg string, adHocFields []LogFields) { ... }
```

The `Logger.Log()` receives the `AdHoc fields` as a `slice` of `LogFields` just to ease the forwarding from the log levels: 

```go
// logLevels.go
func (l *Logger) Trace(msg string, adHocFields ...LogFields) { 
  l.Log(LvlTrace, msg, adHocFields)
}
```

Example of `AdHoc fields` usage:

```go
someLogger := logger.NewDefault()
someLogger.Trace("some log", logger.LogFields{
  "adHoc-A": "value-A",
  "adHoc-B": "value-B",
})
/*
  {
    "msg": "some log",
    "lvl": 1,
    "adHoc-A": "value-A",
    "adHoc-B": "value-B"
  }
*/
```

`AdHoc fields` are defined as variadic arguments just to simulate "optional arguments", that don't officially exist in `go`. Note that if you pass more than one `LogFields` variadic argument, the latter ones will override the previous ones:

```go
someLogger := logger.NewDefault()
someLogger.Trace("some log", 
  logger.LogFields{"adHoc-A": "value-A", "adHoc-B": "value-B"},
  logger.LogFields{"adHoc-A": "new value"},
)
/*
  {
    "msg": "some log",
    "lvl": 1,
    "adHoc-A": "new value",
    "adHoc-B": "value-B"
  }
*/
```

`AdHoc fields` are very suitable to log values that are different for every created log, like the ID of some user, for example:

```go
someLogger := logger.NewDefault()
someLogger.Info("User created", logger.LogFields{"id": userID})
/*
  {
    "msg": "User created",
    "lvl": 4,
    "id": 10
  }
*/
```
