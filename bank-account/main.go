package main

import "fmt"

type BankAccount struct {
	balance int
}

func (b *BankAccount) deposit(val int) {
	b.balance += val
}

func (b *BankAccount) withdraw(val int) {
	if b.balance-val < 0 {
		fmt.Println("insufficient funds")
		return
	}
	b.balance -= val
}

func (b *BankAccount) getBalance() int {
	return b.balance
}

func main() {
	account := BankAccount{balance: 0}

	account.deposit(10)
	account.withdraw(50)

	account.deposit(10)
	account.deposit(10)
	account.deposit(10)
	account.deposit(10)

	fmt.Printf("%v\n", account.getBalance())

	account.withdraw(50)

	fmt.Printf("%v\n", account.getBalance())

	account.deposit(10)
	account.deposit(10)

	fmt.Printf("%v\n", account.getBalance())
}
