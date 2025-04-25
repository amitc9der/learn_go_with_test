package pointererrors

import (
	"errors"
	"fmt"
)

type BitCoin int

type Wallet struct {
	balance BitCoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amt BitCoin) {
	w.balance += amt
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amt BitCoin) error {
	if w.balance < amt {
		return ErrInsufficientFunds
	}

	w.balance -= amt
	return nil
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
