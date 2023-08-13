package gateway

import "github.com/IamP5/ms-wallet/wallet-core/internal/entity"

type ClientGateway interface {
	FindByID(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
