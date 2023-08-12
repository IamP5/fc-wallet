package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Jane Doe", "j@j")
	account2 := NewAccount(client2)

	account1.Deposit(1000)
	account2.Deposit(1000)

	transaction, err := NewTransaction(account1, account2, 100)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 900.0, account1.Balance)
	assert.Equal(t, 1100.0, account2.Balance)
}

func TestCreateTransactionWithInsufficientFunds(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Jane Doe", "j@j")
	account2 := NewAccount(client2)

	account1.Deposit(1000)
	account2.Deposit(1000)

	transaction, err := NewTransaction(account1, account2, 2000)

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Error(t, err, "insufficient funds")
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}

func TestCreateTransactionWithNegativeAmount(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Jane Doe", "j@j")
	account2 := NewAccount(client2)

	account1.Deposit(1000)
	account2.Deposit(1000)

	transaction, err := NewTransaction(account1, account2, -100)

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Error(t, err, "amount must be greater than zero")
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}

func TestCreateTransactionWithSameAccount(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	account1 := NewAccount(client1)

	account1.Deposit(1000)

	transaction, err := NewTransaction(account1, account1, 100)

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Error(t, err, "account from and account to must be different")
	assert.Equal(t, 1000.0, account1.Balance)
}

func TestCreateTransactionWithNilAccountFrom(t *testing.T) {
	client2, _ := NewClient("Jane Doe", "j@j")
	account2 := NewAccount(client2)

	account2.Deposit(1000)

	transaction, err := NewTransaction(nil, account2, 100)

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Error(t, err, "account to is required")
	assert.Equal(t, 1000.0, account2.Balance)
}

func TestCreateTransactionWithNilAccountTo(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	account1 := NewAccount(client1)

	account1.Deposit(1000)

	transaction, err := NewTransaction(account1, nil, 100)

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Error(t, err, "account to is required")
	assert.Equal(t, 1000.0, account1.Balance)
}
