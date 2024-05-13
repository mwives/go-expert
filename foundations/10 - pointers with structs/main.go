package main

import "fmt"

type Account struct {
	Balance float64
}

func NewAccount() *Account {
	return &Account{Balance: 0}
}

// Changes made in a copy of the Account
func (a Account) Simulate(value float64) float64 {
	a.Balance += value
	fmt.Printf("New (simulated) balance: %v\n", a.Balance)
	return a.Balance
}

// Using pointer to change the actual field 😎
func (a *Account) Deposit(value float64) float64 {
	a.Balance += value
	fmt.Printf("New (actual) balance: %v\n", a.Balance)
	return a.Balance
}

func main() {
	myAccount := Account{
		Balance: 420,
	}

	myAccount.Simulate(999)
	myAccount.Deposit(246)

}
