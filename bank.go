package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	sync.Mutex
}

func (b *BankAccount) Balance() (balance int) {
	return b.balance
}

func (b *BankAccount) Withdraw(amount int) (newBalance int) {
	b.Lock()
	b.balance = b.balance - amount
	b.Unlock()
	return b.balance
}

func (b *BankAccount) Deposit(amount int) (newBalance int) {
	b.Lock()
	b.balance = b.balance + amount
	b.Unlock()
	return b.balance
}

func New(startingAmount int) *BankAccount {
	return &BankAccount{
		balance: startingAmount,
	}
}

func main() {
	myBankAccount := New(120)
	fmt.Println("starting amount: ", myBankAccount.Balance())

	go secondTransaction(myBankAccount)

	fmt.Println("depositing 100 on main go routine")
	myBankAccount.Deposit(100)
	fmt.Println("checking balance on main go routine after depositing 100: ", myBankAccount.Balance())

	go thirdTransaction(myBankAccount)

	time.Sleep(time.Second * 3) // Just to prevent the main routine from ending immediately
}

func secondTransaction(b *BankAccount) {
	fmt.Println("withdrawing 100 on second go routine")
	b.Withdraw(100)
	fmt.Println("checking balance on second go routine after withdrawing 100: ", b.Balance())
}

func thirdTransaction(b *BankAccount) {
	fmt.Println("depositing 500 on third go routine")
	b.Deposit(500)
	fmt.Println("checking balance on third go routine after depositing 500: ", b.Balance())
}
