package mylogger1

type Logger struct {
	instance ITypeLogger
	loglevel LogLevel
	//config   string
	//getlevel LogLevel
}

func NewLogger(typelogger ITypeLogger, Level LogLevel) *Logger {
	logger := &Logger{instance: typelogger, loglevel: Level}
	return logger
}

func (l Logger) Log(message string) {
	l.logBase(LEVEL_LOG, message)
}

func (l Logger) Debug(message string) {
	l.logBase(LEVEL_DEBUG, message)
}

func (l Logger) Error(err error) {
	l.logBase(LEVEL_ERROR, err.Error())
}

func (l Logger) logBase(level LogLevel, message string) {
	if l.loglevel > level {
		return
	}

	if l.instance != nil {
		if message != "" {
			message = ": " + message
		}

		l.instance.Log(level, message)
	}
}
