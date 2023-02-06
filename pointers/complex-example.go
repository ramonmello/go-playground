package main

import "fmt"

// A struct to represent a bank account
type Account struct {
	balance float64
	owner   string
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
	fmt.Printf("Deposited %.2f to account of %s\n", amount, a.owner)
}

func (a *Account) Withdraw(amount float64) {
	if a.balance < amount {
		fmt.Println("Insufficient balance")
		return
	}
	a.balance -= amount
	fmt.Printf("Withdrew %.2f from account of %s\n", amount, a.owner)
}

func (a *Account) Balance() float64 {
	return a.balance
}

func main() {
	// Create an account with initial balance of 100 and owner "John"
	account := &Account{balance: 100, owner: "John"}

	// Deposit 50 to the account
	account.Deposit(50)

	// Withdraw 30 from the account
	account.Withdraw(30)

	// Print the current balance of the account
	fmt.Printf("Current balance of %s's account: %.2f\n", account.owner, account.Balance())
}
