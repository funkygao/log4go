package log4go

import "strings"

func ToLogLevel(levelStr string) Level {
	level := TRACE
	switch strings.ToLower(levelStr) {
	case "info":
		level = INFO

	case "warn":
		level = WARNING

	case "error":
		level = ERROR

	case "debug":
		level = DEBUG

	case "trace":
		level = TRACE

	case "alarm":
		level = ALARM
	}

	return level
}
