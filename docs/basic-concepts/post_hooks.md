---
sidebar_position: 5
---

# PostHooks

You can customize the `PostHooks` of the `Logger` instance using the following methods:

```go
// hooks.go
func PostHooks(hooks Hooks) *Logger { ... }
func RawPostHooks(hooks Hooks) *Logger { ... }
```

:::note
The returned `Logger` instance will be a copy of the previous one, sharing the same `Configuration` struct. The only difference will be the `PostHooks`. 

More info at the [Loggers clonage](../advanced/loggers_clonage.md) page.
:::

As you may have seen, `PostHooks` will be executed at the [Thirty phase](life_cycle.md#async-phase-3-post-handling) of the life cycle, right after the `AdHoc fields` are applied.

:::info
`PostHooks` share the same characteristics as the `PreHooks`: just a function that receive a `Log` struct and returns the field value. 

Read the introduction in the [PreHooks](pre_hooks.md) page for details.
:::

The advantage of `PostHooks` is that since they're being applied **after** the `Base fields`, `PreHooks` and `AdHoc fields`, you can query more fields using the `Log.Field` method from the hook argument:

```go
someLogger := logger.New(logger.DefaultConfig()).
  Fields(logger.LogFields{
    "base-field": 10,
  }).
  PreHooks(logger.Hooks{
    "preHook": func(l logger.Log) interface{} {
      return l.Field("base-field").(int) + 10
    },
  }).
  PostHooks(logger.Hooks{
    "postHook": func(l logger.Log) interface{} {
      // Querying the adHoc field that will be defined at log creation time
      return l.Field("preHook").(int) + l.Field("adHoc-field").(int)
    },
  }).
  Outputs(logger.OutputJsonToWriter(os.Stdout, nil))
someLogger.Debug("some log", logger.LogFields{
  "adHoc-field": 5,
})
/*
  {
    "msg": "some log",
    "lvl": 2,
    "base-field": 10,
    "preHook": 20,
    "adHoc-field": 5,
    "postHook": 25,
  }
*/
```

## PostHooks method

Similar to the [PreHooks method](pre_hooks.md#prehooks-method), this method will return a new copy of the `Logger` instance with the given `PostHooks` applied, overriding previous `PostHooks` with the same `key`:

```go
firstLogger := logger.New(logger.DefaultConfig()).
  PostHooks(logger.Hooks{
    "field-A": func(_ logger.Log) interface{} { return "dynamic value-A" },
    "field-B": func(_ logger.Log) interface{} { return "dynamic value-B" },
    "field-C": func(_ logger.Log) interface{} { return "dynamic value-C" },
  }).
  Outputs(logger.OutputJsonToWriter(os.Stdout, nil))
firstLogger.Info("first log")
/*
  {
    "msg": "first log",
    "lvl": 4,
    "field-A": "dynamic value-A",
    "field-B": "dynamic value-B",
    "field-C": "dynamic value-C"
  }
*/

secondLogger := firstLogger.
  PostHooks(logger.Hooks{
    "field-B": func(_ logger.Log) interface{} { return "new value" },
  })
secondLogger.Info("second log")
/*
  {
    "msg": "second log",
    "lvl": 4,
    "field-A": "dynamic value-A",
    "field-B": "new value",
    "field-C": "dynamic value-C"
  }
*/
```

## RawPostHooks method

If you want to reset the `Logger` `PostHooks` you can use the `RawPostHooks` method, that will set the `Logger` `PostHooks` right away, ignoring any previous values (returning a new copy of the `Logger` instance, just like `PostHooks`). Example:

```go
firstLogger := logger.New(logger.DefaultConfig()).
  PostHooks(logger.Hooks{
    "field-A": func(_ logger.Log) interface{} { return "dynamic value-A" },
    "field-B": func(_ logger.Log) interface{} { return "dynamic value-B" },
    "field-C": func(_ logger.Log) interface{} { return "dynamic value-C" },
  }).
  Outputs(logger.OutputJsonToWriter(os.Stdout, nil))
firstLogger.Info("first log")
/*
  {
    "msg": "first log",
    "lvl": 4,
    "field-A": "dynamic value-A",
    "field-B": "dynamic value-B",
    "field-C": "dynamic value-C"
  }
*/

secondLogger := firstLogger.
  RawPostHooks(logger.Hooks{
    "field-B": func(_ logger.Log) interface{} { return "value" },
  })
secondLogger.Info("second log")
/*
  {
    "msg": "second log",
    "lvl": 4,
    "field-B": "value",
  }
*/
```

## PostHooks querying PostHooks

Note that if you try to use the `Log.Field()` method inside a `PostHook` to get query some `PostHook` field, you will fall into the same (possible) bug that was discussed when a `PreHook` query another `PreHook` field ([here](pre_hooks#prehooks-querying-prehooks)). **Pay attention**.
