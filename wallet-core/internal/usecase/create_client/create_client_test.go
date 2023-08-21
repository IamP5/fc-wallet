package create_client

import (
	"github.com/IamP5/ms-wallet/wallet-core/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUseCase(m)

	output, err := uc.Execute(CreateClientInputDTO{
		Name:  "John Doe",
		Email: "j@j",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "j@j", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
