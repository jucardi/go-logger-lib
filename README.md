# go-logger-lib

Is a simple logger library which defines logger interfaces and factory to use a a contract to easily register and swap logger implementations affecting the behavior of a component.

This library is does not actually provide a logger implementation but instead by default uses a `github.com/sirupsen/logrus` standard logger instance as its default.

## Getting started

```bash
go get github.com/jucardi/go-logger-lib
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

* `RegisterLogger(name string, logger ILogger)`: Registers an instance of ILogger to be returned as the singleton instance by the given name.
  * `name`: The logger implementation name.
  * `logger`: The logger instance.
* `RegisterLoggerBuilder(name string, ctor func(...interface{}) ILogger)`: Registers an ILogger constructor function which will be used to create a new instance of the logger when requested instance by the given name. The constructor allows a variadic interface{} array that can be used for optional constructor variables, such as the instance name of a logger, the package name where it is used, etc. It is up to the custom implementation of a logger to use these values.
  * `name`: The logger implementation name.
  * `ctor`: The constructor function used to create the ILogger instance.
* `Get(name string, args ...interface{}) ILogger`: Returns an instance of the requested logger by its name. Returns nil if a logger by that name has not been previously registered.
  * `name`: The implementation name of the instance to be retrieved.
  * `args`: Variadic interface{} array as optional arguments for a registered logger constructor.
* `LoggersList() []string`: Returns the list of loggers that have been registered to the factory.
* `Contains(name string) bool`: Contains indicates if a logger by the given name is contained by the factory.

## Predefined loggers

* **Logrus**: Predefined logger implementation powered by `github.com/sirupsen/logrus`. It is assigned as default to be de default logger which responds to the static functions in the `log` package. Can also be obtained by the using the `"logrus"` name (also defined in the `LoggerLogrus` constant). Eg: `log.Get(log.LoggerLogrus)`
* **Nil Logger**: This implementation, as the name suggests, is a logger that does nothing when its functions are called. Can be obtained by the using the `"nil"` name (also defined in the `LoggerNil` constant). Eg: `log.Get(log.LoggerNil)`. To easily set the Nil Logger as the default logger, simply pass `nil` value to the `log.SetDefault` function. Eg `log.SetDefault(nil)`

## TODO

* Do an implementation for the default Golang `log` package.