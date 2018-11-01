package log

import (
	"bytes"

	"github.com/sirupsen/logrus"
)

// LoggerLogrus indicates the name of the predefined logrus ILogger implementation
const LoggerLogrus = "logrus"

type logrusImpl struct {
	logrus.Logger
}

func (l *logrusImpl) SetLevel(level Level) {
	l.Logger.Level = logrus.Level(level)
}

func (l *logrusImpl) GetLevel() Level {
	return Level(logrus.GetLevel())
}

func (l *logrusImpl) SetFormatter(formatter IFormatter) {
	l.Logger.SetFormatter(&logrusFormatter{
		f: formatter,
	})
}

// NewLogrus creates a new instance of the logrus implementation of ILogger
func NewLogrus() ILogger {
	ret := &logrusImpl{
		Logger: *logrus.StandardLogger(),
	}
	ret.SetFormatter(NewTerminalFormatter())
	return ret
}

type logrusFormatter struct {
	f IFormatter
}

func (f *logrusFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	buffer := entry.Buffer
	if buffer == nil {
		buffer = &bytes.Buffer{}
	}
	if err := f.f.Format(buffer, &Entry{
		Data:      entry.Data,
		Timestamp: entry.Time,
		Level:     Level(entry.Level),
		Message:   entry.Message,
	}); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
