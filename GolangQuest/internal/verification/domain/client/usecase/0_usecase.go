package usecase

import (
	"context"
	client "github.com/eNViDAT0001/Backend/internal/verification/domain/client"
	"github.com/eNViDAT0001/Backend/internal/verification/domain/jwt/storage/io"
)

type clientStorage struct {
}

func (c clientStorage) FindClientByKey(clientKey string) (io.Client, error) {
	//TODO implement me
	panic("implement me")
}

func (c clientStorage) FindClientById(clientId string) (io.Client, error) {
	//TODO implement me
	panic("implement me")
}

func (c clientStorage) CheckUserExist(ctx context.Context, email string, phone string) error {
	//TODO implement me
	panic("implement me")
}

func NewClientStorage() client.Storage {
	return &clientStorage{}
}
