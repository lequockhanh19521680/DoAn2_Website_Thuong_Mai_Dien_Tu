package usecase

import (
	"github.com/eNViDAT0001/Backend/internal/user/domain/user"
)

type userUseCase struct {
	userSto user.Storage
}

func NewUserUseCase(userSto user.Storage) user.UseCase {
	return &userUseCase{
		userSto: userSto,
	}
}
