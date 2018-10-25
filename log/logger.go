package log

// ILogger defines the contract for a logger interface to be used by the mgo and mongo packages.
// This interface matches most commonly used loggers which should make it simple to assign any
// logger implementation being used. By default it uses the sirupsen/logrus standard logger
// implementation.
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

	// SetFormatter sets a custom formatter to display the logs
	SetFormatter(formatter IFormatter)
}
