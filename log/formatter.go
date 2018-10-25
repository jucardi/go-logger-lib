package log

import (
	"io"
	"os"
	"strings"
	"text/template"
)

const (
	TemplateDefault = `{{ .Level }} {{ .Timestamp }} | {{ .Message }}`
)

// The Formatter interface is used to implement a custom Formatter. It takes an
// `Entry`. It exposes all the fields, including the default ones:
//
// * `entry.Data["msg"]`. The message passed from Info, Warn, Error ..
// * `entry.Data["time"]`. The timestamp.
// * `entry.Data["level"]. The level the entry was logged at.
//
// Any additional fields added with `WithField` or `WithFields` are also in
// `entry.Data`. Format is expected to return an array of bytes which are then
// logged to `logger.Out`.
type IFormatter interface {
	Format(io.Writer, *Entry) error
	SetTemplate(string) error
}

// BaseFormatter base structure for formatters
type BaseFormatter struct {
	templateHandler *template.Template
	helpers         template.FuncMap
}

func (f *BaseFormatter) SetTemplate(tmpl string) error {
	t, err := template.New("formatter").Funcs(f.helpers).Parse(tmpl)
	if err != nil {
		return err
	}
	f.templateHandler = t
	return nil
}

// BaseTerminalFormatter base structure to create formatters for a terminal
type BaseTerminalFormatter struct {
	BaseFormatter
	// Set to true to bypass checking for a TTY before outputting colors.
	ForceColors bool

	// Force disabling colors.
	DisableColors bool

	// Override coloring based on CLICOLOR and CLICOLOR_FORCE. - https://bixense.com/clicolors/
	EnvironmentOverrideColors bool
	supportsColor             *bool
	colorScheme               LevelColorScheme
}

func (f *BaseTerminalFormatter) isColored() bool {
	if f.supportsColor == nil {
		supportsColor := f.ForceColors

		if force, ok := os.LookupEnv("CLICOLOR_FORCE"); ok && force != "0" {
			supportsColor = true
		} else if ok && force == "0" {
			supportsColor = false
		} else if os.Getenv("CLICOLOR") == "0" {
			supportsColor = false
		} else if strings.Contains(os.Getenv("TERM"), "color") {
			supportsColor = true
		}
		f.supportsColor = &supportsColor
	}

	return *f.supportsColor && !f.DisableColors
}

func (f *BaseTerminalFormatter) SetColorScheme(scheme LevelColorScheme) {
	f.colorScheme = scheme
}
