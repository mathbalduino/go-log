---
sidebar_position: 8
---

# Life cycle

Any created log will pass between four phases:

1. `sync` [Creation](#sync-phase-1-creation): The log is created and `AdHoc fields` are collected
2. `sync` [Pre handling](#sync-phase-2-pre-handling): `PreHooks` are evaluated
3. `sync`/`async` [Post handling](#async-phase-3-post-handling): `Base fields`, `PreHooks`, `AdHoc fields` are applied and `PostHooks` are evaluated and applied
4. `sync`/`async` [Output](#async-phase-4-output): The final log fields are forwarded to every configured `Output`

## Sync Phase 1: Creation

This phase is characterized by the call to the `Logger` log level method itself:

```go
someLogger.Trace("some msg", logger.LogFields{ ...someFields... })
```

At this point, the `AdHoc fields` are collected and a referece to them is stored, in order to be able to process them by the following phases.

## Sync Phase 2: Pre handling

Right after the `AdHoc fields` are collected, the `PreHooks` are evaluated. It means that all the configured `PreHooks` functions will be called at this stage, in a synchronous way. The returned `PreHooks` function values will be stored inside the created log, in order to be applied later.

:::note
At this stage, the log fields aren't ready yet. We're just calculating the _possible_ values of the fields. These values can be overriden by `AdHoc fields` or `PostHooks`, at later phases, but we don't know for sure.
:::

## (A)Sync Phase 3: Post handling

This is the **most expensive phase**, since it's responsible to let the log fields ready to be used by the `Outputs`.

Here, the `Base fields` are copied from the `Logger` instance that created the log into a new `LogFields` map, followed by the copy of the evaluated `PreHooks` (from phase 2), followed by the copy of the `AdHoc fields` (from phase 1). The `PostHooks` values will be last to be evaluated and copied to the final log fields. More information, see the [Fields override order](override_order.md) page.

At the end of the phase 3, the `message` and `level` log fields will be applied to the `LogFields` map, overriding any previous value that used the keys respectively configured. Something like this:

```go
logFields[lvlFieldKey] = log.lvl
logFields[msgFieldKey] = log.msg
```

:::info
If the `Logger` is set to be `async`, this phase will be executed by a different go routine than the one that created the log
:::

## (A)Sync Phase 4: Output

Just a `for` loop over the `Logger` configured `Outputs` slice, forwarding the final log fields, created at the phase 3:

```go
// Not real production-code (just to illustrate)
for _, output := range logger.outputs {
  output(lvl, msg, logFields)
}
```

:::info
If the `Logger` is set to be `async`, this phase will be executed by a different go routine than the one that created the log
:::
