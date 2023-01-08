package cart

import (
	"context"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/cart/entities"
)

type Storage interface {
	GetDetailByID(ctx context.Context, cartID uint) (entities.CartDetail, error)
	ListCartByUserID(ctx context.Context, userID uint, filter paging.GetListInput) ([]entities.CartDetail, error)
	CountListCartByUserID(ctx context.Context, userID uint, filter paging.GetListInput) (total int64, err error)
	DeleteCart(ctx context.Context, userID uint, cartID uint) error
}
