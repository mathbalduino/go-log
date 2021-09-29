---
sidebar_position: 2
---

# Log creation

Every time a new log is created, it's necessary to alloc a new map to represent it's fields. This new map will be filled with the `Logger` `Base fields`, `PreHooks`, `AdHoc fields` and `PostHooks`. This process can be expensive, depending on the `Logger` size.

If you need performance and your logs are heavy, you should set the `Logger` to be `async` (see [Async Logger](../basic-concepts/async_logger.md) for details), and all this process will be performed into a dedicated go routine. Just remember that the `PreHooks` are always executed by the same go routine that created the `Log`.

:::note
The final `LogFields` will be an entire new map. You can, of course, change the value of some of it's keys directly inside the `Output` functions, but it's **not recommended**. The `LogFields` map is intended to be used as a `read-only` map.
:::

:::info
The life time of these maps are very short. They're alive just while the `Output` functions are still being executed.
:::

If you want to see details about the implementation, see the `Logger.Log` and `handleLog` method/function, at the root package. The entire process is done inside them.
