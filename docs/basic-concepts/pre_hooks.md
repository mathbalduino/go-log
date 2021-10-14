---
sidebar_position: 3
---

# PreHooks

You can customize the `PreHooks` of the `Logger` instance using the following methods:

```go
// hooks.go
func PreHooks(hooks Hooks) Logger { ... }
func RawPreHooks(hooks Hooks) Logger { ... }
```

:::note
The returned `Logger` instance will be a copy of the previous one, sharing the same `Configuration` struct. The only difference will be the `PreHooks`. 

More info at the [Loggers clonage](../advanced/loggers_clonage.md) page.
:::

As you may have seen in the introduction, `PreHooks` are executed at the [Second phase](life_cycle.md#sync-phase-2-pre-handling) of the life cycle, right after the log creation.

`Hooks` are, essentially, just a function that will be called at log creation time and will return the value for some field, allowing you to create dynamic fields:

```go
// hooks.go
// Just an alias
type Hooks = map[string]func(Log) interface{}
```

You may have noticed that `Hooks` functions receive one argument of the type `Log`. This type is a `struct` used internally by the library to handle the log creation:

```go
// handleLog.go
type Log struct {
  ...
}
func (l Log) Field(key string) interface{} { ... }
```

Using this `struct`, you can query the fields of the created log using the `Field` method.

:::info
It's important to distinguish between this method and the `Logger.Field()` method. They're **not** the same.
:::

This method will return the current value for some field, if it's evaluated at the time of the call. Don't expect to call this method inside `PreHooks` and get a value that will be evaluated just by some `PostHook`, because they're not ready.

:::caution
It's not recommended that your `Hooks` depend on the `Logger` `Configuration`, because it can cause trouble when trying to change the `Configuration` dynamically. For details, see [Dynamic Configuration](../advanced/dynamic_config.md).
:::

:::danger
You have to be very carefull when writing `Hooks`, because the library is not prepared to handle any kind of `panic` that can occur inside them, and it can cause issues.
:::

## PreHooks method

This method, as the `Fields`/`RawFields`, will return a new copy of the `Logger` instance, with the given `PreHooks` applied. So, lets say you want to timestamp your logs every time they're created, you can create the following `PreHook`:

```go
someLogger := logger.New(logger.DefaultConfig()).
  PreHooks(logger.Hooks{
    "timestamp": func(_ logger.Log) interface{} {
		  return time.Now().Second() 
  }}).
  Outputs(logger.OutputJsonToWriter(os.Stdout, nil))
someLogger.Debug("some log")
/*
  {
    "msg": "some log",
    "lvl": 2,
    "timestamp": 23464356
  }
*/
```

This method will override any `PreHook` of the previous `Logger` instance with a `key` that clashes with some of the new ones. Example:

```go
firstLogger := logger.New(logger.DefaultConfig()).
  PreHooks(logger.Hooks{
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
  PreHooks(logger.Hooks{
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

## RawPreHooks method

If you want to reset the `Logger` `PreHooks` you can use the `RawPreHooks` method, that will set the `Logger` `PreHooks` right away, ignoring any previous values (returning a new copy of the `Logger` instance, just like `PreHooks`). Example:

```go
firstLogger := logger.New(logger.DefaultConfig()).
  PreHooks(logger.Hooks{
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
  RawPreHooks(logger.Hooks{
    "field-B": func(_ logger.Log) interface{} { return "different" },
  })
secondLogger.Info("second log")
/*
  {
    "msg": "second log",
    "lvl": 4,
    "field-B": "different",
  }
*/
```

## PreHooks querying PreHooks

Don't use the `Log.Field()` method to query `PreHook` fields INSIDE a `PreHook`. `PreHooks` are applied using an iteration over the `PreHooks` map, and the order is not guaranteed ([read more](https://golangdocs.com/golang-iterate-over-a-map)). Don't do this ~~or do it~~:

```go
someLogger := logger.NewDefault().
  PreHooks(logger.Hooks{
    // Even if the "field-A" is defined BEFORE "field-A-plus5",
    "field-A": func(l logger.Log) interface{} { return 10 },
    
    "field-A-plus5": func(l logger.Log) interface{} {
      v := l.Field("field-A")
      // "v" may be nil (maybe not), so this line may (maybe not) panic
      return v.(int) + 5
    },
  })

// when iterating/executing the PreHooks inside the following log creation
someLogger.Info("some log")
```
