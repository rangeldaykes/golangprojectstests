package mylogger1

import (
	"log"
	"os"
)

type ConsoleLoger struct {
	logger *log.Logger
}

func NewConsoleLoger() *ConsoleLoger {
	l := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	return &ConsoleLoger{logger: l}
}

func (c *ConsoleLoger) Log(level LogLevel, message string) {
	c.logger.Println(message)
}
