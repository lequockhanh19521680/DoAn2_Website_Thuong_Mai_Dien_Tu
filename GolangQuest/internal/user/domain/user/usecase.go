package user

import (
	"context"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/user/domain/user/storage/io"
	"github.com/eNViDAT0001/Backend/internal/user/entities"
)

type UseCase interface {
	GetUserDetailByID(ctx context.Context, ID uint) (*entities.User, error)
	GetUserList(ctx context.Context, input *paging.GetListInput) ([]*entities.User, error)
	CreateUser(ctx context.Context, input *io.CreateUserInput) (userID uint, err error)
	UpdateUser(ctx context.Context, userID uint32, input *io.UpdateUserInput) error
	SetPassword(ctx context.Context, userID uint32, password string, newPassword string) error
	DeleteUserByIDs(ctx context.Context, IDs []uint) error
	GetUserByUsername(ctx context.Context, username string) (*entities.User, error)
}
