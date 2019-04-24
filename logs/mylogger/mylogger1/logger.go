package mylogger1

import (
	"fmt"
	"runtime/debug"
)

type Logger struct {
	instance ITypeLogger
	loglevel LogLevel
	context  LogContext
}

func NewLogger(typelogger ITypeLogger, Level LogLevel) *Logger {
	logger := &Logger{instance: typelogger, loglevel: Level, context: LogContext{}}
	return logger
}

func (l *Logger) Log(message string) {
	l.logBase(LEVEL_LOG, message)
}

func (l *Logger) Debug(message string) {
	l.logBase(LEVEL_DEBUG, message)
}

func (l *Logger) Assert(condition bool, message string) {
	if !condition {
		msg := fmt.Sprintf("%s %s", "Assert: ", message)
		l.logBase(LEVEL_ASSERT, msg)
	}
}

func (l *Logger) Error(err error, message string) {

	msg := fmt.Sprintf("%s - %s \n %s \n %s", "Erro: ", message, err.Error(), debug.Stack())

	l.logBase(LEVEL_ERROR, msg)
}

func (l *Logger) ContextStart(message string) {
	message = l.context.Start(message)
	l.logBase(LEVEL_LOG, message)
}

func (l *Logger) ContextEnd() {
	message := l.context.End()
	l.logBase(LEVEL_LOG, message)
}

func (l *Logger) logBase(level LogLevel, message string) {
	if l.loglevel > level {
		return
	}

	if l.instance != nil {
		if message != "" {
			message = ": " + l.context.GetMessage(message)
		}

		l.instance.Log(level, message)
	}
}
