package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/go-stomp/stomp"
)

const defaultPort = ":61613"

var serverAddr = flag.String("server", "192.168.228.150:61623", "STOMP server endpoint")
var messageCount = flag.Int("count", 999999, "Number of messages to send/receive")
var queueName = flag.String("queue", "/topic/systrac.Bridge.TempoReal.dbDispBusNet", "Destination queue")
var helpFlag = flag.Bool("help", false, "Print help text")
var stop = make(chan bool)

// these are the default options that work with RabbitMQ
var options []func(*stomp.Conn) error = []func(*stomp.Conn) error{
	stomp.ConnOpt.Login("guest", "guest"),
	stomp.ConnOpt.Host("/"),
	stomp.ConnOpt.HeartBeat(time.Duration(99999)*time.Hour, time.Duration(99999)*time.Hour),
	stomp.ConnOpt.Header("content-length", "no"),
}

func main() {
	flag.Parse()
	if *helpFlag {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	subscribed := make(chan bool)
	go recvMessages(subscribed)

	// wait until we know the receiver has subscribed
	<-subscribed

	//go sendMessages()

	<-stop
	<-stop
}

func sendMessages() {
	//defer func() {
	//stop <- true
	//}()

	conn, err := stomp.Dial("tcp", *serverAddr, options...)
	if err != nil {
		println("cannot connect to server", err.Error())
		return
	}

	for i := 1; i <= *messageCount; i++ {
		//text := fmt.Sprintf("Message #%d", i)
		text := fmt.Sprint(`{"Mensagem":[{"Codigo":248,"HistoricoPeriodico":0,"ListaHistoricosBus":[{"DataHora":"2019-01-27T13:24:21","DataHoraGmt":"2019-01-27T16:24:21","Latitude":-28681207,"Longitude":-49364587,"Velocidade":51,"Parado":"F","Chuva":"T","AberturaPorta":"F","FechamentoPorta":"F","Hodometro":434896,"Horimetro":1395376,"HodometroSobra":645,"Consumo":null,"ListaRFID":[]}],"DataUltimoPeriodicoConhecido":"0001-01-01T00:00:00","DataUltimoPeriodicoConhecidoGMT":"0001-01-01T00:00:00"}]}`)

		//text := fmt.Sprint(`{"Mensagem":[{"Codigo":248,"HistoricoPeriodico":0}]}`)

		//fmt.Println(text)
		//err = conn.Send(*queueName, "text/plain", []byte(text), nil)

		err = conn.Send(
			*queueName,
			"text/plain",
			[]byte(text),
			stomp.SendOpt.NoContentLength,
		)

		if err != nil {
			println("failed to send to server", err)
			return
		}

		time.Sleep(time.Duration(2000) * time.Millisecond)
	}
	println("sender finished")
}

/* func(sendTimeout, recvTimeout time.Duration) func(*Conn) error {
	return func(c *Conn) error {
		c.options.WriteTimeout = sendTimeout
		c.options.ReadTimeout = recvTimeout
		return nil
	}
} */
func escreve(mm string) {
	println("Actual:", mm)
}

func recvMessages(subscribed chan bool) {
	defer func() {
		stop <- true
	}()

	conn, err := stomp.Dial("tcp", *serverAddr, options...)

	if err != nil {
		println("cannot connect to server", err.Error())
		return
	}

	sub, err := conn.Subscribe(*queueName, stomp.AckAuto)
	if err != nil {
		println("cannot subscribe to", *queueName, err.Error())
		return
	}
	close(subscribed)

	idx := 0
	for msg := range sub.C {
		if msg.Err != nil {
			fmt.Println("err")
		}

		idx++
		println(idx)

		actualText := string(msg.Body)
		go escreve(actualText)

		if idx >= 999 {
			println("acabou")
		}
	}

	// for i := 1; i <= *messageCount; i++ {
	// 	msg := <-sub.C
	// 	expectedText := fmt.Sprintf("Message #%d", i)
	// 	actualText := string(msg.Body)
	// 	if expectedText != actualText {
	// 		println("Expected:", expectedText)
	// 		println("Actual:", actualText)
	// 	}
	// }
	println("receiver finished")

}
