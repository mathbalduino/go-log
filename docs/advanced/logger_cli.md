---
sidebar_position: 5
---

# Logger CLI - Experimental

Since this library was built to compose my personal stack, I extended the original `Logger` into the `LoggerCLI`, that's intended to be used by my CLI tools. You can find this at the `loggerCLI` package, inside the root package. This is just an experimental feature, so it can be removed in future versions.

`LoggerCLI` logs can be nested, creating a tree-like structure, easing visualization. You can create an instance of the `LoggerCLI` using the `New` function, inside the `loggerCLI` package.

Inside the `loggerCLI` package there's another package: `beautify`. You can use this package to pretty-print a set of logs.

When using some CLI tool that uses the `LoggerCLI`, you can just pipe the output to the `beautify` package:

```sh
some-tool --json --some-flag some-file | go run PATH/TO/go-log/loggerCLI/beautify
```

:::note
The `beautify` package expects to receive one log per line, parsed as a JSON object. Just remember to create and set a `--json` flag, forwarding it to the `New` function that creates the `LoggerCLI` instance.
:::

This package is far from optimized, and expects that the output knows how to print ANSI codes (virtually any modern terminal knows how to do it). For details, see the [source code](https://github.com/mathbalduino/go-log).

## Creating a new LoggerCLI instance

Using the `New` function, at the package `loggerCLI`, you can get a new `LoggerCLI` instance. This constructor receives 3 booleans: 

1. `json`: Controls the output. If true, it will print one log per line, parsed as `json` object
2. `debug`: Enables the `Debug` log level
3. `trace`: Enabled the `Debug` and `Trace` log levels. If true, the `debug` boolean will be ignored.

## Creating a CLI tool with LoggerCLI

Logs created by the `LoggerCLI` are intended to be nested, so, the API is built in a way that makes it possible. Every log level method will return another instance of `LoggerCLI`, that will nest it's own logs inside the `LoggerCLI` that created it. Example:

```go
flags := cliFlags()
lvl0 := loggerCLI.New(flags.json, flags.debug, flags.trace)
lvl1 := lvl0.Info("Lvl 0")
lvl2 := lvl1.Warn("Lvl 1")
lvl0.Info("Lvl 1 - again")
lvl2.Error("Lvl 2")
lvl1.Warn("Lvl 1 - again")

// Will output:
//    [ INFO ] Lvl 0
//    [ WARN ] Lvl 1
//    [ INFO ] Lvl 0 - again
//    [ ERROR ] Lvl 2
//    [ WARN ] Lvl 1 - again

// If forwarded to "beautify" (with "json" = true):
//    [ INFO ] Lvl 0
//    |--[ WARN ] Lvl 1
//    |  '--[ ERROR ] Lvl 2
//    '--[ WARN ] Lvl 1 - again
//    [ INFO ] Lvl 0 - again
```

:::caution
Don't forget that it's recommended to create, at least, one CLI Flag: `--json`. The value of this flag should be forwarded to the `New` function, easing the pipe of the output to the `beautify` package. You can, of course, create flags to the `debug` and `trace` flags too.
:::
