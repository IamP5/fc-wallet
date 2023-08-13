package gateway

import "github.com/IamP5/ms-wallet/wallet-core/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
