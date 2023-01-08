package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/cart/entities"
)

func (c cartStorage) CountListCartByUserID(ctx context.Context, userID uint, filter paging.GetListInput) (total int64, err error) {
	var count int64
	db := wrap_gorm.GetDB()
	query := db.Table(entities.CartDetail{}.TableName()).Where("user_id = ?", userID)

	paging.SetCountListPagingQuery(&filter, entities.CartDetail{}.TableName(), query)

	err = query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, err
}
