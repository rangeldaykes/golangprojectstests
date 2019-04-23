package mylogger1

import (
	"container/list"
)

type LogContext struct {
	methodStack  list.List
	logPrefix    string
	contextCount int
}

func (lc *LogContext) GetMessage(message string) string {
	return lc.logPrefix + message
}

func (lc *LogContext) Start(operation string) string {
	lc.methodStack.PushBack(operation)

	operation = "-> " + operation

	lc.contextCount++
	lc.logPrefix = "  " + lc.logPrefix

	return operation
}

func (lc *LogContext) End() string {
	var operation string

	if lc.contextCount == 0 {
		return ""
	}

	f := lc.methodStack.Back()
	lc.methodStack.Remove(f)

	operation = f.Value.(string)

	if lc.contextCount == 1 {
		operation = lc.GetMessage("<- " + operation)

		lc.contextCount--
		lc.logPrefix = lc.logPrefix[2:]

	} else {
		lc.contextCount--
		lc.logPrefix = lc.logPrefix[2:]

		operation = lc.GetMessage("<- " + operation)
	}

	return operation
}
