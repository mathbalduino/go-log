---
sidebar_position: 6
---

# Fields override order

As you may have noticed, the log fields are constantly being overwritten. It's important that you understand the order in which all this operation is executed.

Everything starts at the `Base fields`, go through `PreHooks`, `AdHoc fields` and end at the `PostHooks`:

1. `Base fields`: gets overriden by everyone
2. `PreHooks`: override `Base fields`, gets overriden by `AdHocs` and `PostHooks`
3. `AdHoc fields`: override `Base fields` and `PreHooks`, gets overriden by `PostHooks`
4. `PostHooks`: overrides everyone

It means that if you define some field using the `Fields` method, it will be overriden by fields defined by the `PreHooks` method, if they have the same key. Example:

```go
someLogger := logger.NewDefault().
  Fields(logger.LogFields{
    "lvl": "Trying to override the 'lvl' field",
    "msg": "Trying to override the 'msg' field",
    "base-field": 1,
    "preHook": 1,
    "adhoc-field": 1,
    "postHook": 1,
  }).
  PreHooks(logger.Hooks{
    "lvl": func(l logger.Log) interface{} { return "Trying to override the 'lvl' field (2)" }
    "msg": func(l logger.Log) interface{} { return "Trying to override the 'msg' field (2)" }
    "preHook": func(l logger.Log) interface{} { return 2 },
    "adhoc-field": func(l logger.Log) interface{} { return 2 },
    "postHook": func(l logger.Log) interface{} { return 2 },
  }).
  PostHooks(logger.Hooks{
    "lvl": func(l logger.Log) interface{} { return "Trying to override the 'lvl' field (4)" }
    "msg": func(l logger.Log) interface{} { return "Trying to override the 'msg' field (4)" }
    "postHook": func(l logger.Log) interface{} { return 4 },
  })
someLogger.Debug("some log", logger.LogFields{
  "lvl": "Trying to override the 'lvl' field (3)",
  "msg": "Trying to override the 'msg' field (3)",
  "adHoc-field": 3,
  "postHook": 3,
})
/*
  {
    "msg": "some log",
    "lvl": 2,
    "base-field": 1,
    "preHook": 2,
    "adHoc-field": 3,
    "postHook": 4,
  }
*/
```

:::info
The `lvl` and `msg` fields will be the last to be applied, overriding everyone else. You cannot change this behaviour
:::
