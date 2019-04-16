// https://github.com/bestmethod/go-logger
package main

import (
	"fmt"
	"os"

	Logger "github.com/bestmethod/go-logger"
)

func main() {
	logger := new(Logger.Logger)
	err := logger.Init("SUBNAME",
		"SERVICENAME",
		Logger.LEVEL_DEBUG|Logger.LEVEL_INFO|Logger.LEVEL_WARN, Logger.LEVEL_ERROR|Logger.LEVEL_CRITICAL,
		Logger.LEVEL_NONE)
	if err != nil {
		fmt.Fprintf(os.Stderr, "CRITICAL Could not initialize logger. Quitting. Details: %s\n", err)
		os.Exit(1)
	}

	// standard logger messages
	logger.Info("This is info message")
	logger.Debug("This is debug message")
	logger.Error("This is error message")
	logger.Warn("This is warning message")
	logger.Critical("This is critical message")

	// logger messages, like Printf (auto-discovery of printf happens, so same functions are used)
	logger.Info("%s %v", "This is info message", 10)
	logger.Debug("%s %v", "This is debug message", 10)
	logger.Error("%s %v", "This is error message", 10)
	logger.Warn("%s %v", "This is warning message", 10)
	logger.Critical("%s %v", "This is critical message", 10)
}
