package paybroker

import "fmt"

type PaymentBroker struct {}


func (p *PaymentBroker) ProcessPayment(amount float32) bool{
	fmt.Println("Processing payment with broker")
	return true
}