package log

var defaultLogger ILogger

func init() {
	RegisterBuilder(LoggerLogrus, func(...interface{}) ILogger {
		return NewLoggrus()
	})
	RegisterBuilder(LoggerNil, func(...interface{}) ILogger {
		return NewNilLogger()
	})

	defaultLogger = Get(LoggerLogrus)
}

// SetDefault sets a logger instance as the default logger.
func SetDefault(logger ILogger) {
	if logger == nil {
		defaultLogger = NewNilLogger()
		return
	}

	defaultLogger = logger
}

// GetDefault returns the the instance currently set as the default logger
func GetDefault() ILogger {
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

// Debugln logs a message at level Debug on the default logger.
func Debugln(args ...interface{}) {
	defaultLogger.Debugln(args...)
}

// Info logs a message at level Info on the default logger.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Infof logs a message at level Info on the default logger.
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Infoln logs a message at level Info on the default logger.
func Infoln(args ...interface{}) {
	defaultLogger.Infoln(args...)
}

// Warn logs a message at level Warn on the default logger.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Warnf logs a message at level Warn on the default logger.
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Warnln logs a message at level Warn on the default logger.
func Warnln(args ...interface{}) {
	defaultLogger.Warnln(args...)
}

// Error logs a message at level Error on the default logger.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Errorf logs a message at level Error on the default logger.
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Errorln logs a message at level Error on the default logger.
func Errorln(args ...interface{}) {
	defaultLogger.Errorln(args...)
}

// Fatal logs a message at level Fatal on the default logger.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Fatalf logs a message at level Fatal on the default logger.
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// Fatalln logs a message at level Fatal on the default logger.
func Fatalln(args ...interface{}) {
	defaultLogger.Fatalln(args...)
}

// Panic logs a message at level Panic on the default logger.
func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

// Panicf logs a message at level Panic on the default logger.
func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

// Panicln logs a message at level Panic on the default logger.
func Panicln(args ...interface{}) {
	defaultLogger.Panicln(args...)
}
