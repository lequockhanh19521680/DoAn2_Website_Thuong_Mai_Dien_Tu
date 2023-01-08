package banner

import (
	"github.com/eNViDAT0001/Backend/internal/store/domain/favorite"
)

type favoriteHandler struct {
	favoriteUC favorite.HttpHandler
}

func NewFavoriteHandler(favoriteUC favorite.HttpHandler) favorite.HttpHandler {
	return &favoriteHandler{favoriteUC: favoriteUC}
}
