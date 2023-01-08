package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/internal/cart/entities"
)

func (c cartItemStorage) UpdateCartItem(ctx context.Context, itemID uint, cartID uint, quantity int) error {
	db := wrap_gorm.GetDB()
	err := db.Model(entities.CartItem{}).
		Where("id = ?", itemID).
		Where("cart_id = ?", cartID).
		Update("quantity", quantity).
		Error
	if err != nil {
		return nil
	}
	return nil
}
