package log

import (
	"errors"
	"fmt"
	"github.com/jucardi/go-terminal-colors"
	"io"
)

// TextFormatter formats logs into text
type TerminalFormatter struct {
	BaseTerminalFormatter
}

type LevelColorScheme map[Level][]fmtc.Color

const (
	TemplateTerminalFormatter = `{{ level . }} {{ color (timestamp . "HH:mm:ss") "Cyan" }} {{ .Message }}`
)

var DefaultColorScheme = LevelColorScheme{
	DebugLevel: []fmtc.Color{fmtc.Bold, fmtc.DarkGray},
	InfoLevel:  []fmtc.Color{fmtc.Bold, fmtc.White, fmtc.BgBlue},
	WarnLevel:  []fmtc.Color{fmtc.Black, fmtc.BgYellow},
	ErrorLevel: []fmtc.Color{fmtc.Bold, fmtc.White, fmtc.BgRed},
	FatalLevel: []fmtc.Color{fmtc.Bold, fmtc.White, fmtc.BgRed},
	PanicLevel: []fmtc.Color{fmtc.Bold, fmtc.White, fmtc.BgRed},
}

func NewTerminalFormatter() IFormatter {
	ret := &TerminalFormatter{}
	ret.helpers = getDefaultHelpers()
	ret.SetTemplate(TemplateTerminalFormatter)
	ret.SetColorScheme(DefaultColorScheme)
	return ret
}

// Format renders a single log entry
func (f *TerminalFormatter) Format(writer io.Writer, entry *Entry) error {
	if f.templateHandler == nil {
		return errors.New("no template parser found")
	}

	if writer == nil {
		return errors.New("writer cannot be nil")
	}

	entry.AddMetadata(metadataColorEnabled, f.isColored())
	entry.AddMetadata(metadataColorScheme, f.colorScheme)
	if err := f.templateHandler.Execute(writer, entry); err != nil {
		return fmt.Errorf("unable to write log to io writer, %s", err.Error())
	}

	fmt.Fprintln(writer)
	return nil
}
