package mylogger1_test

import (
	"mylogger1"
	"testing"
)

func TestLogger(t *testing.T) {
	//mylogger1.NewLogger(mylogger1.NewFileLogger("Logteste", ""), mylogger1.LEVEL_LOG)

	log := mylogger1.NewLogger(
		mylogger1.NewFileLogger(
			"Logteste",
			"/home/rangelsantos/Documentos/"),
		mylogger1.LEVEL_ERROR)

	log.Log("TestLogger")
}
