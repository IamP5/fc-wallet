package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j")
	account := NewAccount(client)

	assert.NotNil(t, account)
	assert.Equal(t, 0.0, account.Balance)
	assert.Equal(t, client, account.Client)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestDepositAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j")
	account := NewAccount(client)
	account.Deposit(100)

	assert.Equal(t, 100.0, account.Balance)
}

func TestWithdrawAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j")
	account := NewAccount(client)
	account.Deposit(100)
	account.Withdraw(50)

	assert.Equal(t, 50.0, account.Balance)
}
