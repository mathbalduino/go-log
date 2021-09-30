---
sidebar_position: 1
---

# Log fields

A log is represented by this library as a collection of fields, represented by the `LogFields` `go` type, that translates to a simple map:

```go
// fields.go
// Just an alias
type LogFields = map[string]interface{}
```

With this definition, this `LogFields` map:

```go
m := LogFields{
  "msg": "some log message",
  "lvl": 10,
  ...
  "custom": "field",
  ...
}
```

Can be seen, for example, as the following `json` object:

```json
{
  "msg": "some log message",
  "lvl": 10,
  ...
  "custom": "field",
  ...
}
```

In the following chapters, you will see how you can transform the `LogFields` `go` type into whatever you want (including `json`). If you don't want to wait, you can go directly to the [Outputs](outputs.md) page.

Inside the log fields, there are two fields that will always be present:

- One to identify the log `string` message (default key: `msg`)
- One to identify the log `uint64` level (default key: `lvl`)

A log is always created using the information from the `Logger` instance that created it. The `Logger` instance can be seen as a "template" of how the created logs should look like. This "template" (the `Logger` instance) can have 4 different types of fields:

1. [Base fields](base_fields.md)
2. [PreHooks](pre_hooks.md)
3. [Ad hoc fields](adhoc_fields.md)
2. [PostHooks](post_hooks.md)

Each one of them will be discussed in the following chapters. 

:::note
Since the `message` and `level` fields will always be present, if you try to set another field with the same `key` used by them, it will be **overridden**. Don't do it.
:::

:::tip
You can customize the `key` used to represent the required fields using the [Configuration](configuration.md).
:::
