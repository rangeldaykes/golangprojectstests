package main

import (
	"fmt"
	"strconv"

	"github.com/rangeldaykes/gosimplefilelog"
)

func main() {

	log := gosimplefilelog.NewLogger(
		gosimplefilelog.NewFileLogger(
			"Logteste",
			"/home/rangelsantos/"),
		gosimplefilelog.LEVEL_LOG)

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
