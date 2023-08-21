package create_transaction

import (
	"context"
	"github.com/IamP5/ms-wallet/wallet-core/internal/entity"
	"github.com/IamP5/ms-wallet/wallet-core/internal/event"
	"github.com/IamP5/ms-wallet/wallet-core/internal/usecase/mocks"
	"github.com/IamP5/ms-wallet/wallet-core/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "john@j")
	account1 := entity.NewAccount(client1)
	account1.Deposit(1000)

	client2, _ := entity.NewClient("Jane Doe", "jane@j")
	account2 := entity.NewAccount(client2)
	account2.Deposit(1000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	ctx := context.Background()

	uc := NewCreateTransactionUseCase(mockUow, dispatcher, transactionCreatedEvent)
	output, err := uc.Execute(ctx, inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
