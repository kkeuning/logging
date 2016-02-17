# Logging Adapters

[![Build Status](https://travis-ci.org/goadesign/goa.svg?branch=master)](https://travis-ci.org/goadesign/goa)

This repository contains a set of log adapters for use by [goa](http://goa.design). They make it possible to have goa
log informational and error messages to various logger backends. Each adapter exists in its own
package named after the corresponding logger package.

Once instantiated adapters can be used by initializing the goa Log package variable, for example:

```go
  func main() {
    // ...

    // Setup logger adapter
    logger := log15.New()
    goa.Log = goalog15.New(logger)

    // ...
}
```

### Taking Advantage of the Context

New adapters may be written that take advantage of the request context to log additional
information:

```go
func (l *Logger) Info(ctx context.Context, msg string, data ...goa.KV) {
  // extract log info from ctx
  // and log it
}
```
