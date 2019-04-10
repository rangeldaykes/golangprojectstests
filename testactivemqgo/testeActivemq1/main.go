package main

import (
	"fmt"
	"testactivemqgo/testeActivemq1/activemq"
)

func main() {
	fmt.Println("inicio")

	mq := activemq.NewActiveMQ("127.0.0.1:61623")

	destino := "/topic/systrac.Bridge.TempoReal.dbDispBusNet"

	if err := mq.Subscribe(destino, handler); err != nil {
		fmt.Println("AMQ ERROR:", err)
	}

	_ = mq.Send(destino, "test")

	fmt.Println("fim")
}

func handler(err error, msg string) {
	fmt.Println("chegou")
	fmt.Println("AMQ MSG:", err, msg)
}
