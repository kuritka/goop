package main

import (
	"goop/3_Polymorphism/paybroker"
)


/*
Processing payment with broker

Process finished with exit code 0
*/

type IPaymentOption interface {
	ProcessPayment(float32) bool
}

func main() {
	var option IPaymentOption

	option = &paybroker.PaymentBroker{}

	option.ProcessPayment(3000)
}


