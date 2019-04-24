package mylogger1_test

import (
	"fmt"
	"mylogger1"
	"strconv"
	"testing"
)

func TestLogger(t *testing.T) {
	//mylogger1.NewLogger(mylogger1.NewFileLogger("Logteste", ""), mylogger1.LEVEL_LOG)

	log := mylogger1.NewLogger(
		mylogger1.NewFileLogger(
			"Logteste",
			"/home/rangelsantos/Documentos/"),
		mylogger1.LEVEL_LOG)

	log.Log("TestLogger1")

	str := "ab"
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Error(err, "conversão ruin")
	}
	fmt.Println(i)

	log.Assert(i == 1, "i não tem o valor 1")

	log.ContextStart("start1")
	log.Log("after start context 1")

	log.ContextStart("start2")
	log.Log("after start context 2")
	log.ContextEnd()

	log.ContextEnd()

	log.Log("after end context 1")
}
