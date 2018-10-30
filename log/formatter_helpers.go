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
		"Sprint":     fmt.Sprint,
		"Sprintf":    fmt.Sprintf,
		"ToUpper":    strings.ToUpper,
		"ToLower":    strings.ToLower,
		"Replace":    strings.Replace,
		"TimeFormat": iso8601.TimeToString,
		"level":      levelFn,
		"timestamp":  timestampFn,
		"message":    messageFn,
		"colorCode":  colorCodeFn,
		"colorName":  colorNameFn,
		"colored":    colorFieldFn,
		"scheme":     colorSchemeFn,
	}
}

func levelFn(entry Entry) string {
	return colorFieldFn(FieldLevel, entry, " %s ")
}

func timestampFn(format string, entry Entry) string {
	timeStr := iso8601.TimeToString(entry.Timestamp, format)
	return colorSchemeFn(FieldTimestamp, timeStr, entry)
}

func messageFn(entry Entry, newLinePadding ...string) string {
	if len(newLinePadding) > 0 && newLinePadding[0] != "" {
		return colorSchemeFn(FieldMessage, strings.Replace(entry.Message, "\n", "\n"+newLinePadding[0], -1), entry)
	}
	return colorFieldFn(FieldMessage, entry)
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
