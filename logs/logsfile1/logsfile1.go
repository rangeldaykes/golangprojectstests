package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)

	file := filepath.Join(dir, "log.txt")

	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger := log.New(logFile, "Log: ", log.LstdFlags)

	logger.Println("main started")

	logger.Fatalln("fatal message")

	logger.Panicln("panic message")
}
