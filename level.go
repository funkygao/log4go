package log4go

import "strings"

func ToLogLevel(levelStr string, defaultLevel Level) (lvl Level) {
	lvl = defaultLevel
	switch strings.ToUpper(levelStr) {
	case "FINEST":
		lvl = FINEST
	case "FINE":
		lvl = FINE
	case "DEBUG":
		lvl = DEBUG
	case "TRACE":
		lvl = TRACE
	case "INFO":
		lvl = INFO
	case "WARNING":
		lvl = WARNING
	case "ERROR":
		lvl = ERROR
	case "CRITICAL":
		lvl = CRITICAL
	default:
	}

	return
}
