---
sidebar_position: 1
---

# Introduction

`go-log` is an opinionated, modularized, structured and production-ready GO logging library. 

It was created, at first, to be used in `webservers` and `CLI` tools. ~~But feel free to explore the possibilities.~~

## Getting Started

You will need to install `go-log` before starting. To do it, execute the following command:

```sh
go get github.com/mathbalduino/go-log
```

The library is built around the `Logger` struct type, that contains all the main methods:

```go
// new.go
type Logger struct {
  ...
}
```

Once you get a `Logger` instance, you're ready to throw `logs` by calling the correspondent methods:

```go
// logLevels.go
func (l *Logger) Trace(msg string, adHocFields ...LogFields)  { ... }
func (l *Logger) Debug(msg string, adHocFields ...LogFields)  { ... }
func (l *Logger) Info(msg string, adHocFields ...LogFields)   { ... }
func (l *Logger) Warn(msg string, adHocFields ...LogFields)   { ... }
func (l *Logger) Error(msg string, adHocFields ...LogFields)  { ... }
func (l *Logger) Fatal(msg string, adHocFields ...LogFields)  { ... }
func (l *Logger) ErrorFrom(e error, adHocFields ...LogFields) { ... }
func (l *Logger) FatalFrom(e error, adHocFields ...LogFields) { ... }
```

As you may have seen, this library contains some builtin [Log levels](basic-concepts/log_levels.md), but you can [customize it](basic-concepts/log_levels.md#extending-the-log-levels).

To [create a new Logger](basic-concepts/logger_creation.md) instance, you have two options:
1. Function `New`, that will create an empty `Logger` instance (takes a [Configuration](basic-concepts/configuration.md) struct as argument)
2. Function `NewDefault`, that will create a `Logger` instance with the [Default configuration](basic-concepts/configuration.md#default-configuration)

## Basic concepts

To understand this library, you'll have to understand 4 different concepts: 
1. [Log fields](#log-fields)
2. [Hooks](#hooks)
3. [Log life cycle](#log-life-cycle)
4. [Outputs](#outputs)

### Log fields

`Logs` are represented as a set of fields, just like a JSON object:

```json
{
  "msg": "some log message",
  "lvl": 10,
  ...
  "custom": "field",
  ...
}
```

Looking at the GO type definition for log fields, it becomes clear that it's just a `map` from string to anything:

```go
// fields.go
// Just an alias
type LogFields = map[string]interface{}
```

Fields can be statically or dynamically added or overrided, so you have freedom to structure your log as you please.

You can set log fields in 3 different ways, using:

1. [Base fields](basic-concepts/base_fields.md)
2. [PreHooks](basic-concepts/pre_hooks.md)
3. [Ad hoc fields](basic-concepts/adhoc_fields.md)
4. [PostHooks](basic-concepts/post_hooks.md)

### Hooks

Hooks are used to define dynamic fields. They're divided in two categories:
1. [PreHooks](basic-concepts/pre_hooks.md)
2. [PostHooks](basic-concepts/post_hooks.md)

A set of `Hooks` is defined as a map from field name to a function that evaluates to the field value:
```go
// hooks.go
// Just an alias
type Hooks = map[string]func(Log) interface{}
```

:::note
`Hooks` are intended to be used to calculate the value of non-constant fields, just like the log timestamp, for example
:::

### Log life cycle

Every created log will go through four different phases, sending the final log fields to the configured outputs:

1. [Creation](basic-concepts/life_cycle.md#sync-phase-1-creation)
2. [Pre handling](basic-concepts/life_cycle.md#sync-phase-2-pre-handling)
3. [Post handling](basic-concepts/life_cycle.md#async-phase-3-post-handling)
4. [Output](basic-concepts/life_cycle.md#async-phase-4-output)

For more details, see the [Log life cycle](basic-concepts/life_cycle.md) page.

### Outputs

When some log is created it needs to be sent to somewhere, be it the stdout, some database, or anything. `Outputs` are functions that handle these scenarios, forwarding the created logs to it's final destiny.

The library contains some basic `Output` functions already defined, so you can play around with them ~~or create your own, from scratch~~.

At the end of the day, `outputs` are just functions that will receive some log fields:
```go
// outputs.go
// Just an alias
type Output = func(lvl uint64, msg string, fields LogFields)
```

For more details, see the [Outputs](basic-concepts/outputs.md) page.

## Extras

### Sync vs Async

Logs can be handled in two ways: `sync` or `async`. When set to `async`, the expensive part of the logging operation will run in a different go routine.

In order to configure the `Logger` to be `async`, you will need to read more about [Async Loggers](basic-concepts/async_logger.md).

### Advanced customization

Virtually _everything_ is customizable. You can change the name of the required fields `msg` and `lvl`, create new log levels, new error parsers, etc.

For more details, see the [Configuration](basic-concepts/configuration.md), [Log levels](basic-concepts/log_levels.md) and [Outputs](basic-concepts/outputs.md) pages.
