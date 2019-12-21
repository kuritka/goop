package payment

import "fmt"

type CreditCard struct {
}

func (c *CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Paying with credit card")
	return true
}

type CashAccount struct {}

func (c *CashAccount) ProcessPayment(amount float32) bool {
	fmt.Println("Cash, cash....!")
	return true
}

