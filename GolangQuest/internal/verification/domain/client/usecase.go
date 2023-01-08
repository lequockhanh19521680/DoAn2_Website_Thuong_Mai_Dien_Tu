package user

import (
	"context"
	"github.com/eNViDAT0001/Backend/internal/user/domain/user/storage/io"
)

type UseCase interface {
	CreateClient(ctx context.Context, input *io.CreateUserInput) error
	UpdateClient(ctx context.Context, userID uint32, input *io.UpdateUserInput) error
}
