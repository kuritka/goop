package main

/*
List of funds
**********
cash paid
true
**********
List of funds
Some dollars
0
**********
List of funds
0
List of funds
0

 */

import (
	"fmt"
	"strings"
)

type Account struct {
}

func (a *Account) AvailableFunds() float32 {
	fmt.Println("List of funds")
	return 0
}

func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Process payment")
	return true
}


type CreditAccount struct {
	//syntactic sugar allows us to do this (using embeded types):
	//ca := new(CreditAccount)
	//ca.AvailableFunds()
	Account
}

func (a *CashAccount) ProcessPayment(amount float32) bool {
	fmt.Println("cash paid")
	return true
}


func (a *CashAccount) AvailableFunds() float32 {
	fmt.Println("Some dollars")
	return 0
}



type CashAccount struct {
	//syntactic sugar allows us to do this (using embeded types):
	//ca := new(CreditAccount)
	//ca.AvailableFunds()
	Account
}



type HybridAccount struct{
	CashAccount
	CreditAccount
	creditAccountHidden CreditAccount
}

func (a *HybridAccount ) AvailableFunds() float32{
	return a.CreditAccount.AvailableFunds() + a.CashAccount.AvailableFunds()
}

func main(){
	ca := new(CreditAccount)
	ca.AvailableFunds()
	fmt.Println(strings.Repeat("*",10))
	h := new (HybridAccount)
	fmt.Println(h.ProcessPayment(50))
	fmt.Println(strings.Repeat("*",10))
	fmt.Println(h.AvailableFunds())
	fmt.Println(strings.Repeat("*",10))
	//we shouldnt expose externally, we can implement methodin Hybrid acccount or hide behind some interface etc..
	fmt.Println(h.CreditAccount.AvailableFunds())
	fmt.Println(h.creditAccountHidden.AvailableFunds())
}


