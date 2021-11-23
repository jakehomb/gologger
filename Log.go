package gologger

import (
	"fmt"
)

type Log struct {
	LogLevel LogLevel
	LogMsg   string
}

func (l *Log) ConsoleString() string {
	switch l.LogLevel {
	case DEBUG:
		return fmt.Sprintf("[%s] %s", getLogLevel(l.LogLevel), fmt.Sprint(cyan, l.LogMsg, reset))
	case INFO:
		return fmt.Sprintf("[%s] %s", getLogLevel(l.LogLevel), fmt.Sprint(gray, l.LogMsg, reset))
	case WARN:
		return fmt.Sprintf("[%s] %s", getLogLevel(l.LogLevel), fmt.Sprint(yellow, l.LogMsg, reset))
	case SUCC:
		return fmt.Sprintf("[%s] %s", getLogLevel(l.LogLevel), fmt.Sprint(green, l.LogMsg, reset))
	case ERROR:
		return fmt.Sprintf("[%s] %s", getLogLevel(l.LogLevel), fmt.Sprint(red, l.LogMsg, reset))
	default:
		return fmt.Sprintf("[%s] %s", getLogLevel(l.LogLevel), fmt.Sprint(reset, l.LogMsg))
	}
}

func (l *Log) String() string {
	return fmt.Sprintf("[%s] %s", getLogLevel(l.LogLevel), l.LogMsg)
}

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var cyan = "\033[36m"
var gray = "\033[37m"
var yellow = "\033[33m"
