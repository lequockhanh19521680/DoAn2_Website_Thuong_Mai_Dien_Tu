package usecase

import (
	"github.com/eNViDAT0001/Backend/internal/store/domain/favorite"
)

type favoriteUseCase struct {
	favoriteSto favorite.UseCase
}

func NewFavoriteUseCase(
	favoriteSto favorite.UseCase,

) favorite.UseCase {
	return &favoriteUseCase{
		favoriteSto: favoriteSto,
	}
}
