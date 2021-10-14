---
sidebar_position: 11
---

# Logger creation

To create a new `Logger`, you can use these two functions exported by the root package:

```go
// new.go
func New(config Configuration) Logger { ... }
func NewDefault() Logger { ... }
```

The first one will create an empty `Logger` (with the given `Configuration`), while the latter one will create a `Logger` using the [Default Configuration](configuration.md#default-configuration) and setting two outputs (in this order): 
1. [OutputAnsiToStdout](outputs.md#outputansitostdout)
2. [OutputPanicOnFatal](outputs.md#outputpaniconfatal)

More information about why the order of the `Outputs` is important [here](outputs.md#ordering).

:::caution
Before using the `New` function, check the created `Configuration`. If there's some error inside it (`lvl` and `msg` using the same string, for example), it will `panic`. More information about how to structure your `Configuration` struct [here](configuration.md)
:::
