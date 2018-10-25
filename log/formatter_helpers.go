package log

import (
	"fmt"
	"github.com/jucardi/go-iso8601"
	"github.com/jucardi/go-streams/streams"
	"github.com/jucardi/go-terminal-colors"
	"strings"
	"text/template"
)

const (
	metadataColorEnabled = "colored"
	metadataColorScheme  = "color_scheme"
)

func getDefaultHelpers() template.FuncMap {
	return template.FuncMap{
		"string":    stringFn,
		"upper":     uppercaseFn,
		"lower":     lowercaseFn,
		"fmt":       formatFn,
		"level":     coloredLevelFn,
		"timestamp": timestampFn,
		"colorCode": colorCodeFn,
		"color":     colorFn,
	}
}

func stringFn(arg interface{}) string {
	return fmt.Sprintf("%+v", arg)
}

func uppercaseFn(arg string) string {
	return strings.ToUpper(stringFn(arg))
}

func lowercaseFn(arg interface{}) string {
	return strings.ToLower(stringFn(arg))
}

func formatFn(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func coloredLevelFn(entry Entry) string {
	if v, ok := entry.metadata[metadataColorEnabled]; ok && !v.(bool) {
		return entry.Level.String()
	}
	if _, ok := entry.metadata[metadataColorScheme]; !ok {
		return entry.Level.String()
	}

	scheme := entry.metadata[metadataColorScheme].(LevelColorScheme)
	colors := scheme[entry.Level]

	return fmtc.New().Print(fmt.Sprintf(" %s ", strings.ToUpper(entry.Level.String())), colors...).String()
}

func timestampFn(entry Entry, format string) string {
	return iso8601.TimeToString(entry.Timestamp, format)
}

func colorCodeFn(arg interface{}, colors ...fmtc.Color) string {
	return fmtc.New().Print(arg, colors...).String()
}

func colorFn(arg interface{}, colors ...string) string {
	return fmtc.New().Print(arg, streams.From(colors).Map(func(i interface{}) interface{} {
		ret, _ := fmtc.Parse(i.(string))
		return ret
	}).ToArray().([]fmtc.Color)...).String()
}
