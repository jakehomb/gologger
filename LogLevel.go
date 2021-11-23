package gologger

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	SUCC
	ERROR
)

func getLogLevel(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case SUCC:
		return "SUCC"
	case ERROR:
		return "ERROR"
	default:
		return "ERROR"
	}
}
