---
sidebar_position: 3
---

# Dynamic Configuration

It's possible to change the `Logger` `Configuration` while it's executing, using the `Logger.Configuration()` method, but you will have to be careful.

It's always recommended that you use just one go routine to modify the values of the `Configuration` struct, otherwise you will need to fix write-concurrency issues between the writers. The library will _never_ modify the values of the `Configuration` by itself, only _read_.

## Step by step

When using only one writer, and knowing that the library will just **read** the `Configuration`, you can safely just call the `Logger.Configuration()` method passing a new `Configuration` struct.

The library doesn't control some concurrency issues that can arise when you do it. If you want to change just one `Configuration` struct field, like `LvlsEnabled`, there's no issue. Note that since the library just _reads_ the `Configuration`, and you changed just one thing, it's an atomic operation.

If you needs to change more than one thing, it is, naturally, not atomic. The library, right now, doesn't have a way to guarantee the atomicity of these operations. It means that if you want to change the `LvlFieldName` and `MsgFieldName` at once, logs will be saved in one of the following states:

1. Old `LvlFieldName` and `MsgFieldName` values (before `Logger.Configuration()` call)
2. Old `LvlFieldName` value with the new `MsgFieldName` value (some log was being created at the same time of the `Logger.Configuration()` call)
3. New `LvlFieldName` value with the old `MsgFieldName` value (some log was being created at the same time of the `Logger.Configuration()` call)
4. New `LvlFieldName` and `MsgFieldName` values (`Logger.Configuration()` call completed)

There's a plan to implement some blocking `Logger.Configuration()` method variation, that solves this issue, in the future. Please, let me know if it's necessary.
