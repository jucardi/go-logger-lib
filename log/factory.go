package log

var loggersRepo = map[string]func(...interface{}) ILogger{}

// Register registers an instance of ILogger to be returned as the singleton
// instance by the given name.
//
//   {name}   - The logger implementation name.
//   {logger} - The logger instance.
//
func Register(name string, logger ILogger) {
	loggersRepo[name] = func(...interface{}) ILogger {
		return logger
	}
}

// RegisterBuilder registers an ILogger constructor function which will be used
// to create a new instance of the logger when requested instance by the given name.
//
// The constructor allows a variadic interface{} array that can be used for optional constructor
// variables, such as the instance name of a logger, the package name where it is used, etc.
// It is up to the custom implementation of a logger to use these values.
//
//   {name} - The logger implementation name.
//   {ctor} - The constructor function used to create the ILogger instance.
//
func RegisterBuilder(name string, ctor func(...interface{}) ILogger) {
	loggersRepo[name] = ctor
}

// Get returns an instance of the requested logger by its name. Returns nil if a logger
// by that name has not been previously registered.
//
//   {name} - The implementation name of the instance to be retrieved.
//   {args} - Variadic interface{} array as optional arguments for a registered logger constructor.
//
func Get(name string, args ...interface{}) ILogger {
	if ctor, ok := loggersRepo[name]; ok {
		return ctor(args...)
	}

	return nil
}

// List returns the list of loggers that have been registered to the factory.
func List() []string {
	var ret []string
	for k := range loggersRepo {
		ret = append(ret, k)
	}
	return ret
}

// Contains indicates if a logger by the given name is contained by the factory
func Contains(name string) bool {
	for k := range loggersRepo {
		if name == k {
			return true
		}
	}

	return false
}
