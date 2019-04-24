package mylogger1_test

import (
	"mylogger1"
	"testing"
)

func TestLogger(t *testing.T) {

	//log := mylogger1.NewLogger(mylogger1.NewConsoleLoger(), mylogger1.LEVEL_LOG)

	log := mylogger1.NewLogger(
		mylogger1.NewFileLogger(
			"Logteste",
			"/home/rangelsantos/discod/golangprojectstests/"),
		mylogger1.LEVEL_LOG)

	for i := 0; i < 10000; i++ {
		log.Log("a")
		log.Log("TestLogger1")

		//str := "ab"
		//i, err := strconv.Atoi(str)
		//if err != nil {
		//log.Error(err, "conversão ruin")
		//}
		//fmt.Println(i)

		log.Assert(i == 1, "i não tem o valor 1")

		log.ContextStart("start1")
		log.Log("after start context 1")

		log.ContextStart("start2")
		log.Log("after start context 2")
		log.ContextEnd()

		log.ContextEnd()

		log.Log("after end context 1")
	}
}
