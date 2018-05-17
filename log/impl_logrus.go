package log

import "github.com/sirupsen/logrus"

// LoggerLogrus indicates the name of the predefined logrus ILogger implementation
const LoggerLogrus = "logrus"

type logrusImpl struct {
	logrus.Logger
}

func (l *logrusImpl) SetLevel(level Level) {
	logrus.SetLevel(logrus.Level(level))
}

func (l *logrusImpl) GetLevel() Level {
	return Level(logrus.GetLevel())
}

func NewLoggrus() ILogger {
	return &logrusImpl{
		Logger: *logrus.StandardLogger(),
	}
}
