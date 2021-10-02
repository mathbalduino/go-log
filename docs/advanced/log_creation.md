---
sidebar_position: 2
---

# Log creation

Every time a new log is created, it's necessary to alloc a new map to represent it's fields. 
This new map will be filled with the `Logger` `Base fields`, `PreHooks`, `AdHoc fields` and 
`PostHooks` (The order in which the fields are applied is discussed at 
[Fields override order](../basic-concepts/override_order.md)). This process can be expensive, 
depending on the `Logger` size.

:::caution
The final `LogFields` will be an entire new map. You can, of course, change the value of some 
of it's keys directly inside the `Output` functions, but it's **not recommended**. The `LogFields`
map is intended to be used as a `read-only` map. 
:::

:::note
The life time of these maps are very short. They're useless after the last `Output` function
is finished.
:::

If you need performance and your logs are heavy, you should set the `Logger` to be `async` 
(see [Async Logger](../basic-concepts/async_logger.md) for details), and all this process will 
be performed into a dedicated go routine. Just remember that the `PreHooks` are always executed 
by the same go routine that created the `Log`.

If you want to see details about the implementation, see the 
[Logger.Log source-code](https://github.com/mathbalduino/go-log/blob/79263810d94dd1f2d112727824d1c5256b27951b/logLevels.go#L89)
and [handleLog source-code](https://github.com/mathbalduino/go-log/blob/79263810d94dd1f2d112727824d1c5256b27951b/handleLog.go#L70), 
at the root package. The **_entire process_** is done inside them.
