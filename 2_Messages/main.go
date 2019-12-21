package main

import (
	"fmt"
)

/*
DebitCard payment  500
DebitCard payment  300

 */

type CreditAccount struct {}

func (c *CreditAccount) process(amount float32) {
	fmt.Println("DebitCard payment ", amount)
}

func CreateCreditCardAccount(chargeCh chan float32){
	creditAccount := new(CreditAccount)
	go func (charge chan float32) {
		for amount := range chargeCh {
			creditAccount.process(amount)
		}
	}(chargeCh)
}



func main() {
	cha := make(chan float32)
	CreateCreditCardAccount(cha)
	cha <- 500
	cha <- 300
	fmt.Scanln()
}


