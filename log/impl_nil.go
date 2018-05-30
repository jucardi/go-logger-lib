package log

// LoggerNil indicates the name of the predefined nil ILogger implementation which does nothing when
// the log functions are invoked.
const LoggerNil = "nil"

// ILogger implementation that does nothing on function calls.
type nilLogger struct{}

func (n *nilLogger) GetLevel() Level {
	return DebugLevel
}

func (n *nilLogger) SetLevel(level Level)                      {}
func (n *nilLogger) Debug(args ...interface{})                 {}
func (n *nilLogger) Debugf(format string, args ...interface{}) {}
func (n *nilLogger) Debugln(args ...interface{})               {}
func (n *nilLogger) Info(args ...interface{})                  {}
func (n *nilLogger) Infof(format string, args ...interface{})  {}
func (n *nilLogger) Infoln(args ...interface{})                {}
func (n *nilLogger) Warn(args ...interface{})                  {}
func (n *nilLogger) Warnf(format string, args ...interface{})  {}
func (n *nilLogger) Warnln(args ...interface{})                {}
func (n *nilLogger) Error(args ...interface{})                 {}
func (n *nilLogger) Errorf(format string, args ...interface{}) {}
func (n *nilLogger) Errorln(args ...interface{})               {}
func (n *nilLogger) Fatal(args ...interface{})                 {}
func (n *nilLogger) Fatalf(format string, args ...interface{}) {}
func (n *nilLogger) Fatalln(args ...interface{})               {}
func (n *nilLogger) Panic(args ...interface{})                 {}
func (n *nilLogger) Panicf(format string, args ...interface{}) {}
func (n *nilLogger) Panicln(args ...interface{})               {}

// NewNil creates a new instance of the Nil logger
func NewNil() ILogger {
	return &nilLogger{}
}
