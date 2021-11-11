---
sidebar_position: 5
---

# Logger CLI - Experimental

Since this library was built to compose my personal stack, I extended the original `Logger` into the `LoggerCLI`, that is
intended to be used by my CLI tools. You can find this at the `loggerCLI` package, inside the root package. 

:::caution
This is just an experimental feature, so it can be removed in future versions. Use at your own risk
:::

`LoggerCLI` logs can be nested, creating a tree-like structure, easing visualization. You can create an instance of the `LoggerCLI` using the `New` function, inside the `loggerCLI` package.

Inside the `loggerCLI` package there's another package: `beautify`. You can use this package to pretty-print a set of logs.
When using some CLI tool that uses the `LoggerCLI`, you can just pipe the output to the `beautify` package:

```sh
some-tool --json --some-flag some-file | go run PATH/TO/go-log/loggerCLI/beautify
```

:::note
The `beautify` package expects to receive one log per line, parsed as a JSON object. If you're using some CLI tool, you
can create and set a `--json` CLI flag, and forward it to the `New` function that creates the `LoggerCLI` instance (as a
`bool` value)
:::

:::note
If the nesting is too deep, or if there's some log with a long message, the terminal can break the visualization (by putting
unexpected break lines). To fix it, you can just print the `beautify` output to some file (maybe using the `>` terminal operator?)
and open it using another GUI (maybe some text editor?)
:::

:::caution
This package is far from optimized, and expects that the output knows how to print ANSI codes (virtually any modern terminal
knows how to do it). 

For details, see the [source code](https://github.com/mathbalduino/go-log/tree/main/loggerCLI)
:::

## Creating a new LoggerCLI instance

Using the `New` function, at the package `loggerCLI`, you can get a new `LoggerCLI` instance. This constructor receives 
two arguments: 

1. `json`: Controls the output. If true, it will print one log per line, parsed as `json` object
2. `lvlsEnabled`: An `uint64` value, that represents the [Log Levels](../basic-concepts/log_levels.md)

:::tip
If you're creating some CLI tool and want to provide Log Level customization via its flags, you can use the `ParseLogLevel`
function, exported at the `loggerCLI` package, to parse a human-readable description of the log levels into the `uint64`.

For details, see the [source code](https://github.com/mathbalduino/go-log/blob/3a15937d71e4d2ae6519989ac505fffe80365202/loggerCLI/util.go#L28)
:::

## Creating a CLI tool with LoggerCLI

Logs created by the `LoggerCLI` are intended to be nested, so, the API is built in a way that makes it possible. Every 
log level method will return another instance of `LoggerCLI`, that will nest its own logs inside the `LoggerCLI` that 
created it. Example:

```go
flags := parseCliFlags()
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

:::tip
Don't forget that it's recommended to create, at least, one CLI Flag: `--json`. The value of this flag should be forwarded 
to the `New` function, easing the pipe of the output to the `beautify` package
:::
