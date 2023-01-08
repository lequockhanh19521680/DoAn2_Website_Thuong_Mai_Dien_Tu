package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/internal/user/domain/user/storage/io"
	client "github.com/eNViDAT0001/Backend/internal/verification/domain/client"
)

type clientUseCase struct {
	clientSto client.Storage
}

func (c clientUseCase) CreateClient(ctx context.Context, input *io.CreateUserInput) error {
	//TODO implement me
	panic("implement me")
}

func (c clientUseCase) UpdateClient(ctx context.Context, userID uint32, input *io.UpdateUserInput) error {
	//TODO implement me
	panic("implement me")
}

func NewClientUseCase(clientSto client.Storage) client.UseCase {
	return &clientUseCase{clientSto: clientSto}
}
