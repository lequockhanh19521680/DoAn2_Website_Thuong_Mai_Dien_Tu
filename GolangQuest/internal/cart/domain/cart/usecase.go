package cart

import (
	"context"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/cart/entities"
)

type UseCase interface {
	GetDetailByID(ctx context.Context, cartID uint) (entities.CartDetail, error)
	ListCartByUserID(ctx context.Context, userID uint, filter paging.GetListInput) (carts []entities.CartDetail, total int64, err error)
	DeleteCart(ctx context.Context, userID uint, cartID uint) error
}
