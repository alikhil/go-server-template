package utils

import (
	"fmt"
	lg "log"
)

type LogLevel int32

const (
	Trace LogLevel = iota
	Debug
	Info
	Warn
	Error
	Fatal
)

type Logger struct {
	MinLogLevel LogLevel
}

func (l *Logger) Info(format string, args ...interface{}) {
	if l.MinLogLevel <= Info {
		lg.Printf("[ INFO ]: %s", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Warn(format string, args ...interface{}) {
	if l.MinLogLevel <= Warn {
		lg.Printf("[ WARN ]: %s", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	if l.MinLogLevel <= Fatal {
		lg.Fatalf("[ FATAL ]: %s", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Error(format string, args ...interface{}) {
	if l.MinLogLevel <= Error {
		lg.Printf("[ ERROR ]: %s", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Debug(format string, args ...interface{}) {
	if l.MinLogLevel <= Debug {
		lg.Printf("[ DEBUG ]: %s", fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Trace(format string, args ...interface{}) {
	if l.MinLogLevel <= Trace {
		lg.Printf("[ TRACE ]: %s", fmt.Sprintf(format, args...))
	}
}
