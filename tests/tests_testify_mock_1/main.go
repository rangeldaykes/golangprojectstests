package main

import (
	"fmt"
	"tests_testify_mock_1/messageservice"
)

// A "Production" Example
func main() {
	fmt.Println("Hello World")

	smsService2 := messageservice.SMSService{}
	myService := messageservice.MyService{MessageService: smsService2}
	myService.ChargeCustomer(100)
}
