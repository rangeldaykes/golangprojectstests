package mylogger1

type ITypeLogger interface {
	Log(level LogLevel, message string)
}
