# go-logger-lib

Is a simple logger library which defines logger interfaces and factory to use a a contract to easily register and swap logger implementations affecting the behavior of a component.

This library is does not actually provide a logger implementation but instead by default uses a `github.com/sirupsen/logrus` standard logger instance as its default.

## Getting started

To keep up to date with the most recent version:

```bash
go get github.com/jucardi/go-logger-lib
```

To get a specific version:

```bash
go get gopkg.in/jucardi/go-logger-lib.v1
```

## What inspired me to create this simple library?

Working on a go project that I wanted to make open source, I realized that it would make sense to be able to have a logger and log information based on the results of certain operations. Originally I started using `github.com/sirupsen/logrus` which is a very good Go logging library, but I didn't want the consumers of my project to be forced to use `logrus` if they were already using a different logger.

A generic logger interface enabled me to use it in the project and easily give the freedom to any consumer to simply plug in any logger implementation they feel comfortable using (if any) while still continuing to provide the option for output logs that can be adapted to their stack.

## The `ILogger` interface
The `ILogger` interface was thought of by using the most commonly function names for loggers. It exposes the following functions

```Go
type ILogger interface {
    // SetLevel sets the logging level
    SetLevel(level Level)
    // GetLevel gets the logging level
    GetLevel() Level

    // Debug logs a message at level Debug on the logger.
    Debug(args ...interface{})
    // Debugf logs a message at level Debug on the logger.
    Debugf(format string, args ...interface{})
    // Debugln logs a message at level Debug on the logger.
    Debugln(args ...interface{})

    // Info logs a message at level Info on the logger.
    Info(args ...interface{})
    // Infof logs a message at level Info on the logger.
    Infof(format string, args ...interface{})
    // Infoln logs a message at level Info on the logger.
    Infoln(args ...interface{})

    // Warn logs a message at level Warn on the logger.
    Warn(args ...interface{})
    // Warnf logs a message at level Warn on the logger.
    Warnf(format string, args ...interface{})
    // Warnln logs a message at level Warn on the logger.
    Warnln(args ...interface{})

    // Error logs a message at level Error on the logger.
    Error(args ...interface{})
    // Errorf logs a message at level Error on the logger.
    Errorf(format string, args ...interface{})
    // Errorln logs a message at level Error on the logger.
    Errorln(args ...interface{})

    // Fatal logs a message at level Fatal on the logger.
    Fatal(args ...interface{})
    // Fatalf logs a message at level Fatal on the logger.
    Fatalf(format string, args ...interface{})
    // Fatalln logs a message at level Fatal on the logger.
    Fatalln(args ...interface{})

    // Panic logs a message at level Panic on the logger.
    Panic(args ...interface{})
    // Panicf logs a message at level Panic on the logger.
    Panicf(format string, args ...interface{})
    // Panicln logs a message at level Panic on the logger.
    Panicln(args ...interface{})
}
```

## The factory
Simply provides the following functions

* `New(name string, writer ...io.Writer) ILogger`: Creates a new logger instance using the default builder assigned.
  * `name`: The name of the logger to create.
  * `writer`: (Optional) The io.Writer the logger instance should use. If not provided, it is set to the default writer by the implementation, typically Stdout or Stderr
* `Register(name string, logger ILogger)`: Registers an instance of ILogger to be returned as the singleton instance by the given name.
  * `name`: The logger name.
  * `logger`: The logger instance.
* `Get(name string) ILogger`: Returns an instance of the requested logger by its name. Creates a new logger with the default logger builder if the logger does not exist.
  * `name`: The name of the logger instance to be retrieved.
* `List() []string`: Returns the list of loggers that have been registered.
* `Contains(name string) bool`: Indicates if a logger by the given name exists.
* `SetDefaultBuilder(ctor LoggerBuilder)` Assigns a new constructor function to use as the default logger constructor.

## Predefined logger types

* **Logrus**: Predefined logger implementation powered by `github.com/sirupsen/logrus`. It is assigned as default to be de default logger which responds to the static functions in the `log` package. Can also be obtained by the using the `"logrus"` name (also defined in the `LoggerLogrus` constant). Eg: `log.Get(log.LoggerLogrus)`
* **Nil Logger**: This implementation, as the name suggests, is a logger that does nothing when its functions are called. Can be obtained by the using the `"nil"` name (also defined in the `LoggerNil` constant). Eg: `log.Get(log.LoggerNil)`. To easily set the Nil Logger as the default logger, simply pass `nil` value to the `log.SetDefault` function. Eg `log.SetDefault(nil)`

## TODO

* Do an implementation for the default Golang `log` package.