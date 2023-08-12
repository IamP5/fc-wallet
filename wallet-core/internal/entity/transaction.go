package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
	err := transaction.Validate()
	if err != nil {
		return nil, err
	}

	transaction.Commit()

	return transaction, nil
}

func (t *Transaction) Commit() {
	t.AccountFrom.Withdraw(t.Amount)
	t.AccountTo.Deposit(t.Amount)
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if t.AccountTo == nil {
		return errors.New("account to is required")
	}

	if t.AccountFrom == nil {
		return errors.New("account from is required")
	}

	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient funds")
	}

	if t.AccountTo == t.AccountFrom {
		return errors.New("account from and account to must be different")
	}

	return nil
}
