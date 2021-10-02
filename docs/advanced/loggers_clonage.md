---
sidebar_position: 1
---

# Loggers clonage

Every time you modify some information inside the `Logger`, a new instance will be created by cloning the old version into a new one.

This process can be expensive, depending on the size of the `Logger` being cloned. Because of this, it's recommended that you create all your `Loggers` at startup time, not on-demand.

A `Logger` is represented using the following struct:

```go
// new.go
type Logger struct {
	configuration *Configuration
	fields        LogFields
	preHooks      Hooks
	postHooks     Hooks
	outputs       []Output
}
```

To clone the `Logger` instance, it's necessary to clone the `LogFields` into another map, clone the `PreHooks` and `PostHooks` into another map, and clone the `Outputs` into another slice.

Note that the values are not cloned, just the `containers`. If you have a field called `pointer` that is a pointer to some `int`, for example, the cloned `LogFields` will point to the same `int` (the same to the `Hooks` and `Outputs` methods):

```go
myInt := 10

// Field "pointer" points to "myInt"
someLogger := logger.NewDefault().
	Fields(logger.LogFields{"pointer": &myInt})

// Clone "someLogger" using any method
clone := someLogger.RawPreHooks(nil)

// Now, both "someLogger" and "clone" have a field
// called "pointer" that points to "myInt", but in
// different maps
fmt.Println(*someLogger.Field("pointer").(*int)) // 10
fmt.Println(*clone.Field("pointer").(*int))      // 10

otherLogger := someLogger.Fields(logger.LogFields{"pointer": 5})
fmt.Println(otherLogger.Field("pointer").(int)) // 5
fmt.Println(*clone.Field("pointer").(*int))     // 10
```

:::info
The `Configuration` pointer is just copied to the new `Logger`, without changing it, so all the following versions of some `Logger` will share the same `Configuration`. This way, you can configure your logging strategy in a more centralized way.

If you still want new `Loggers` to point to a different `Configuration`, you can just create another instance using `logger.New()`/`logger.NewDefault()`
:::
