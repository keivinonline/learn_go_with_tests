package wallet

import (
	"errors"
	"fmt"
)

// Defined in fmt package
// Defines how type is printer when using %s prints
type Stringer interface {
	String() string
}

// from: "got 0 want 10"
// to:  "got 0 BTC want 10 BTC"
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// custom type
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("[Deposit]address of wallet balance is %p \n", &w.balance)
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	
	fmt.Printf("[Withdraw]address of wallet balance is %p \n", &w.balance)
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	// signify no issues
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
