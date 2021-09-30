---
sidebar_position: 4
---

# Error tokens

Every error used by the library is an error token. You can use it to do comparisons like this:

```go
e := AsyncHandleLog(...)
if e == logger.ErrNilWaitGroup {
  ...
}
```

You will find all the error tokens inside the `errors.go` file, under the root package.
