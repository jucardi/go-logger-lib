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

	// Debug logs a message at level Debug on the standard logger.
	Debug(args ...interface{})
	// Debugf logs a message at level Debug on the standard logger.
	Debugf(format string, args ...interface{})
	// Debugln logs a message at level Debug on the standard logger.
	Debugln(args ...interface{})

	// Info logs a message at level Info on the standard logger.
	Info(args ...interface{})
	// Infof logs a message at level Info on the standard logger.
	Infof(format string, args ...interface{})
	// Infoln logs a message at level Info on the standard logger.
	Infoln(args ...interface{})

	// Warn logs a message at level Warn on the standard logger.
	Warn(args ...interface{})
	// Warnf logs a message at level Warn on the standard logger.
	Warnf(format string, args ...interface{})
	// Warnln logs a message at level Warn on the standard logger.
	Warnln(args ...interface{})

	// Error logs a message at level Error on the standard logger.
	Error(args ...interface{})
	// Errorf logs a message at level Error on the standard logger.
	Errorf(format string, args ...interface{})
	// Errorln logs a message at level Error on the standard logger.
	Errorln(args ...interface{})

	// Fatal logs a message at level Fatal on the standard logger.
	Fatal(args ...interface{})
	// Fatalf logs a message at level Fatal on the standard logger.
	Fatalf(format string, args ...interface{})
	// Fatalln logs a message at level Fatal on the standard logger.
	Fatalln(args ...interface{})

	// Panic logs a message at level Panic on the standard logger.
	Panic(args ...interface{})
	// Panicf logs a message at level Panic on the standard logger.
	Panicf(format string, args ...interface{})
	// Panicln logs a message at level Panic on the standard logger.
	Panicln(args ...interface{})
}
