package mylogger1

type ILogger interface {
	Log(LogLevel level, string message)
	Assert(string message)
}
