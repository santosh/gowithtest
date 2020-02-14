// Package wallet simulates Bitcoin financial system.
package wallet

import (
	"fmt"
)

// Bitcoin represents a bitcoin currency, underlaying datatype is int
type Bitcoin int

// Wallet has balance to represent a bitcoin wallet.
type Wallet struct {
	balance Bitcoin
}

// Deposit adds amount to existing Wallet.balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns value of Wallet.balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
