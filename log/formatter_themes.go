package log

import "gopkg.in/jucardi/go-terminal-colors.v1"

// LevelColorScheme represents the terminal colors associated to the level parsing for each logging level.
type LevelColorScheme map[Level][]fmtc.Color

// TerminalColorScheme defines terminal colors that are tied to a log level and a field.
type TerminalColorScheme map[string]LevelColorScheme

// TerminalTheme contains the logging theme configuration for terminal logging
type TerminalTheme struct {
	Template string
	Schemes  TerminalColorScheme
}

var (
	TerminalThemeDefault = &TerminalTheme{
		Template: `{{ level . }}{{ timestamp " HH:mm:ss " . }} {{ .Message }}`,
		Schemes: TerminalColorScheme{
			"level": LevelColorScheme{
				DebugLevel: []fmtc.Color{fmtc.Bold, fmtc.DarkGray},
				InfoLevel:  []fmtc.Color{fmtc.Bold, fmtc.Cyan},
				WarnLevel:  []fmtc.Color{fmtc.Bold, fmtc.Yellow},
				ErrorLevel: []fmtc.Color{fmtc.Bold, fmtc.Red},
				FatalLevel: []fmtc.Color{fmtc.Bold, fmtc.Red},
				PanicLevel: []fmtc.Color{fmtc.Bold, fmtc.Red},
			},
			"timestamp": LevelColorScheme{
				DebugLevel: []fmtc.Color{fmtc.BgBlack, fmtc.DarkGray},
				InfoLevel:  []fmtc.Color{fmtc.BgBlack, fmtc.Cyan},
				WarnLevel:  []fmtc.Color{fmtc.BgBlack, fmtc.Yellow},
				ErrorLevel: []fmtc.Color{fmtc.BgBlack, fmtc.Red},
				FatalLevel: []fmtc.Color{fmtc.BgBlack, fmtc.Red},
				PanicLevel: []fmtc.Color{fmtc.BgBlack, fmtc.Red},
			},
		},
	}

	TerminalThemeAlternative = &TerminalTheme{
		Template: `{{ scheme "level" (string " " .Level " ") . }}{{ timestamp " HH:mm:ss " . }} {{ .Message }}`,
		Schemes: TerminalColorScheme{
			"level": LevelColorScheme{
				DebugLevel: []fmtc.Color{fmtc.Bold, fmtc.DarkGray},
				InfoLevel:  []fmtc.Color{fmtc.Bold, fmtc.White, fmtc.BgBlue},
				WarnLevel:  []fmtc.Color{fmtc.Black, fmtc.BgYellow},
				ErrorLevel: []fmtc.Color{fmtc.Bold, fmtc.White, fmtc.BgRed},
				FatalLevel: []fmtc.Color{fmtc.Bold, fmtc.White, fmtc.BgRed},
				PanicLevel: []fmtc.Color{fmtc.Bold, fmtc.White, fmtc.BgRed},
			},
			"timestamp": LevelColorScheme{
				DebugLevel: []fmtc.Color{fmtc.BgBlack, fmtc.DarkGray},
				InfoLevel:  []fmtc.Color{fmtc.BgBlack, fmtc.Cyan},
				WarnLevel:  []fmtc.Color{fmtc.BgBlack, fmtc.Yellow},
				ErrorLevel: []fmtc.Color{fmtc.BgBlack, fmtc.Red},
				FatalLevel: []fmtc.Color{fmtc.BgBlack, fmtc.Red},
				PanicLevel: []fmtc.Color{fmtc.BgBlack, fmtc.Red},
			},
		},
	}

	TerminalThemeCliApp = &TerminalTheme{
		Template: `{{ timestamp " HH:mm:ss " . }} {{ message . "           " }}`,
		Schemes: TerminalColorScheme{
			"timestamp": LevelColorScheme{
				DebugLevel: []fmtc.Color{fmtc.Gray},
				InfoLevel:  []fmtc.Color{fmtc.Cyan},
				WarnLevel:  []fmtc.Color{fmtc.Yellow},
				ErrorLevel: []fmtc.Color{fmtc.Red},
				FatalLevel: []fmtc.Color{fmtc.Red},
				PanicLevel: []fmtc.Color{fmtc.Red},
			},
			"message": LevelColorScheme{
				DebugLevel: []fmtc.Color{fmtc.Gray},
				InfoLevel:  []fmtc.Color{fmtc.White},
				WarnLevel:  []fmtc.Color{fmtc.Yellow},
				ErrorLevel: []fmtc.Color{fmtc.LightRed},
				FatalLevel: []fmtc.Color{fmtc.LightRed},
				PanicLevel: []fmtc.Color{fmtc.LightRed},
			},
		},
	}
)
