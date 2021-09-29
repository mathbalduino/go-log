---
sidebar_position: 10
---

# Async Logger

Until now, everything is `sync`. The log is created and handled by the same go routine.

What do you do if you want to log something that needs to calculate some expensive field value? It can lead to performance issues in your webserver, since the log operation is being executed by the same go routine that runs the core logic.

Even if your `Hooks` are cheap, the `Outputs` may be not. Writing logs to the filesystem/database can lead to performance issues too.

To avoid it, you can set the `Logger` to be `async`. `Async` `Loggers` should be used when you don't want to waste any time doing log operations inside your precious go routine.

To do it, we create another go routine (possibly, more than just one), dedicated to handling logs. That simple.

## The sync of the async

Even with `async` `Loggers`, some time will be wasted creating the log and handling the `PreHooks`, that will always be evaluated by the same go routine that created the log.

So, as you may have noticed, it's always better to use `PostHooks` instead of `PreHooks`, because if you decide to set the `Logger` to be `async` in the future, you don't need to worry about refactoring. Yes and no. Consider the following scenarios:

1. You have some `PreHook` that's a dependency for some `PostHook`. Here, you cannot just transform it into a `PostHook` because of [this](post_hooks.md#posthooks-querying-posthooks).
2. You need to execute some Hook inside the same go routine that created the log. In this case, you will need to use a `PreHook`, by definition.

If your hook don't fall into any of the above scenarios, set it to be a `PostHook`. Virtually _everything_ that is not a `PreHook` will be handled asynchronously.

## AsyncScheduler

The mechanism used to communicate the creation of new logs to the dedicated go routines are ~~surprise, surprise~~: channels!

When you create some `Logger` instance, you can set the `Configuration.AsyncScheduler` field, that is the following interface:

```go
// asyncSheduler.go
type AsyncScheduler interface {
	NextChannel() chan<- Log
	Shutdown()
}
```

The first method, `NextChannel`, should return the receiver channel that was selected to handle the next log.

The second method, `Shutdown`, should be used to signal that the dedicated go routines need to exit. This method is meant to be used to implement some kind of gracefull shutdown in webservers, for example.

You can use your own scheduler, just implement the interface and pass it to the `Logger` `Configuration`.

### Implementing your own AsyncScheduler

When implementing your own `AsyncScheduler`, you will need to call the `AsyncHandleLog` function, from the root package, somewhere inside your dedicated go routines. This is the function that actually handle the created logs:

```go
// asyncSheduler.go
func AsyncHandleLog(ctx context.Context, c <-chan Log, wg WaitGroup) error {
  ...
}
```

The function will wait on the given `channel` argument, receiving new created logs and handling it.

The `context.Context` and `WaitGroup` function arguments are intended to be used to control the exit of the go routines. The `context.Context` is used internally to exit the `AsyncHandleLog` function, using it's `Done()` method (inside a `select` statement), while the `WaitGroup` will notify that some go routine exited.

:::note
Note that the `WaitGroup` is an interface from this library root package. It's not a big deal, since I've used an interface just to ease tests. It's intended to be used as an `sync.WaitGroup`
:::

### DefaultAsyncScheduler

If you don't want to implement your own scheduler, you can use the library builtin. The lib comes with a round-robin-like scheduler, that can be configured to throw `N` go routines, with custom channels `cap` and with support to gracefull-shutdown.

To create a new instance of the builtin scheduler, you will use the `DefaultAsyncScheduler` function, at the root package:

```go
// asyncSheduler.go
func DefaultAsyncScheduler(nGoRoutines uint64, chanCap uint64) AsyncScheduler {
  // Note that the chanCap, if zero, will create blocking channels
  ...
}
```

The returned `AsyncScheduler` is ready to be used by the `Logger`.

The number of go routines or the correct capacity for the channels is a matter of testing/benchmarking, but it's not recommended to create channels with zero capacity, to avoid blocking on every log creation.

If you want some complex scheduling schema, something like creating new go routines on-demand, etc, you will need to create your own implementation.

If you want to enable gracefull shutdown, hold a reference to the returned `AsyncScheduler`, right after calling `DefaultAsyncScheduler`, to be able to call `Shutdown()`.

### Gracefull shutdown

If you want to enable gracefull shutdown in your webserver, you will need to call the `Shutdown` method, of the `AsyncScheduler`. You cannot get a reference to it from the `Logger` itself, since there's no `getter` implemented.

The recommendation is to hold a reference to it, at creation time (or startup time), to centralize all the `Logger` shutdown/configuration logic.

Before calling `Shutdown`, you will need to stop the creation of new logs. You can do this by setting the `LvlsEnabled` configuration to zero (disabling all logs), and them calling `Shutdown`, or just closing the webserver API, for example. For more details, see [Dynamic Configuration](../advanced/dynamic_config.md)

## The real async part

So, what's done after the log is created, `PreHooks` applied and the log sent to the channel? The dedicated go routine will receive the created log, copy the [Base fields](base_fields.md) to a new `LogFields`, apply the [PreHooks](pre_hooks.md), apply the [AdHoc fields](adhoc_fields.md), evaluate and apply every [PostHooks](post_hooks.md), and call every configured [Output](outputs.md).
