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
		"colorName": colorNameFn,
		"colored":   colorFieldFn,
		"scheme":    colorSchemeFn,
	}
}

func stringFn(arg ...interface{}) string {
	return fmt.Sprint(arg...)
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
	return colorFieldFn(FieldLevel, entry, " %s ")
}

func timestampFn(format string, entry Entry, disableColor ...bool) string {
	timeStr := iso8601.TimeToString(entry.Timestamp, format)
	if len(disableColor) > 0 && disableColor[0] {
		return timeStr
	}
	return colorSchemeFn(FieldTimestamp, timeStr, entry)
}

func colorFieldFn(field string, entry Entry, format ...string) string {
	if len(format) > 0 {
		return colorSchemeFn(field, fmt.Sprintf(format[0], entry.getField(field)), entry)
	}
	return colorSchemeFn(field, fmt.Sprint(entry.getField(field)), entry)
}

func colorSchemeFn(schemeName, value string, entry Entry) string {
	if v, ok := entry.metadata[metadataColorEnabled]; ok && !v.(bool) {
		return value
	}
	if v, ok := entry.metadata[metadataColorScheme]; !ok || v == nil {
		return value
	}

	scheme := entry.metadata[metadataColorScheme].(TerminalColorScheme)
	var colors []fmtc.Color

	if v, ok := scheme[schemeName]; !ok || v == nil {
		return value
	} else {
		colors = scheme[schemeName][entry.Level]
	}

	return fmtc.WithColors(colors...).Sprint(value)
}

func colorCodeFn(arg interface{}, colors ...fmtc.Color) string {
	return fmtc.WithColors(colors...).Sprint(arg)
}

func colorNameFn(arg interface{}, colors ...string) string {
	return fmtc.WithColors(streams.From(colors).Map(func(i interface{}) interface{} {
		ret, _ := fmtc.Parse(i.(string))
		return ret
	}).ToArray().([]fmtc.Color)...).Sprint(arg)
}
