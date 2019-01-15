package log

import (
	"encoding/json"
	"strings"
)

var (
	defaultLogger  ILogger
	defaultBuilder LoggerBuilder
)

func init() {
	Register(LoggerLogrus, NewLogrus(""))
	Register(LoggerNil, NewNil())

	defaultLogger = Get(LoggerLogrus)
	defaultBuilder = NewLogrus
}

// SetDefault sets a logger instance as the default logger.
func SetDefault(logger ILogger) {
	if logger == nil {
		defaultLogger = NewNil()
		return
	}

	defaultLogger = logger
}

// GetDefault returns the the instance currently set as the default logger
func Default() ILogger {
	return defaultLogger
}

// SetLevel sets the default logger level.
func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}

// GetLevel returns the default logger level.
func GetLevel() Level {
	return defaultLogger.GetLevel()
}

// Debug logs a message at level Debug on the default logger.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Debugf logs a message at level Debug on the default logger.
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Info logs a message at level Info on the default logger.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Infof logs a message at level Info on the default logger.
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Warn logs a message at level Warn on the default logger.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Warnf logs a message at level Warn on the default logger.
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Error logs a message at level Error on the default logger.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Errorf logs a message at level Error on the default logger.
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the default logger.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Fatalf logs a message at level Fatal on the default logger.
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// Panic logs a message at level Panic on the default logger.
func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

// Panicf logs a message at level Panic on the default logger.
func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

// SetFormatter sets a custom formatter to display the logs
func SetFormatter(formatter IFormatter) {
	defaultLogger.SetFormatter(formatter)
}

// WarnErr logs a warning using the provided message and error if the error is not nil. Does nothing if the error is nil
func WarnErr(err error, message ...string) {
	if err != nil {
		Warn(strings.Join(append(message, err.Error()), " > "))
	}
}

// ErrorErr logs an error using the provided message and error if the error is not nil. Does nothing if the error is nil
func ErrorErr(err error, message ...string) {
	if err != nil {
		Error(strings.Join(append(message, err.Error()), " > "))
	}
}

// FatalErr logs a fatal error using the provided message and error if the error is not nil. Does nothing if the error is nil
func FatalErr(err error, message ...string) {
	if err != nil {
		Fatal(strings.Join(append(message, err.Error()), " > "))
	}
}

// PanicErr logs a panic error using the provided message and error if the error is not nil. Does nothing if the error is nil
func PanicErr(err error, message ...string) {
	if err != nil {
		Panic(strings.Join(append(message, err.Error()), " > "))
	}
}

// DebugObj logs a debug message of a json representation of the provided object. Does nothing if the object is nil.
func DebugObj(obj interface{}) {
	if obj == nil {
		return
	}
	data, err := json.Marshal(obj)
	Debug(string(data), err)
}