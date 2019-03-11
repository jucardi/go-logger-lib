package log

import (
	"bytes"
	"io"
	"os"
)

type ILoggerAsync interface {
	ILogger
	Flush(writer io.Writer, clear ...bool)
	FlushToStdout(clear ...bool)
	Reset()
}

type loggerMemory struct {
	ILogger
	buffer *bytes.Buffer
}

func NewMemory(name string) ILoggerAsync {
	b := &bytes.Buffer{}
	l := NewLogrus(name, b).(ILogger)

	return &loggerMemory{
		ILogger: l,
		buffer:  b,
	}
}

func (l *loggerMemory) Flush(writer io.Writer, clear ...bool) {
	_, _ = writer.Write(l.buffer.Bytes())
	if len(clear) > 1 && clear[0] {
		l.Reset()
	}
}

func (l *loggerMemory) FlushToStdout(clear ...bool) {
	l.Flush(os.Stdout, clear...)
}

func (l *loggerMemory) Reset() {
	l.buffer.Reset()
}
