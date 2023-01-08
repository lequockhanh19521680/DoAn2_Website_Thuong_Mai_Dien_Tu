package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/store/entities"
)

func (s providerStorage) CountListProviderByUserID(ctx context.Context, userID uint, filter paging.GetListInput) (total int64, err error) {
	db := wrap_gorm.GetDB()
	query := db.Model(entities.Provider{}).Where("user_id = ?", userID)
	paging.SetCountListPagingQuery(&filter, entities.Provider{}.TableName(), query)
	err = query.Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
