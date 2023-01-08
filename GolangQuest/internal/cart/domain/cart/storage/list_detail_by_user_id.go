package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/cart/entities"
	"gorm.io/gorm"
)

func (c cartStorage) ListCartByUserID(ctx context.Context, userID uint, filter paging.GetListInput) ([]entities.CartDetail, error) {
	result := make([]entities.CartDetail, 0)
	db := wrap_gorm.GetDB()
	query := db.Table(entities.CartDetail{}.TableName()).Where("user_id = ?", userID)

	paging.SetPagingQuery(&filter, entities.CartDetail{}.TableName(), query)

	query = query.Scan(&result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	if query.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	return result, err
}
