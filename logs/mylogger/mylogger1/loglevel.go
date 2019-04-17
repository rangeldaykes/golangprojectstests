package mylogger1

type LogLevel int

const (
	LEVEL_DEBUG LogLevel = 1 + iota
	LEVEL_LOG
	LEVEL_ERROR
)
