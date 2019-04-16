package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	file := getFile()

	//output := zerolog.ConsoleWriter{Out: file, TimeFormat: time.RFC3339}
	out := zerolog.NewConsoleWriter(
		func(w *zerolog.ConsoleWriter) {
			// Customize time format
			w.TimeFormat = time.RFC822
			// Customize level formatting
			w.FormatLevel = func(i interface{}) string { return strings.ToUpper(fmt.Sprintf("[%-5s]", i)) }
		},
	)
	out.Out = file
	//logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger := zerolog.New(file).Output(out).With().Timestamp().Logger()

	logger.Info().Msg("hello world")
}

func getFile() *os.File {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)

	logFile, err := os.OpenFile(dir+"/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return logFile
}
