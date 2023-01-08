package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/store/entities"
)

func (b bannerStorage) CountListBanner(ctx context.Context, filter paging.GetListInput, bannerID uint) (total int64, err error) {
	var count int64

	db := wrap_gorm.GetDB()

	query := db.Model(entities.Banner{})

	paging.SetCountListPagingQuery(&filter, entities.Banner{}.TableName(), query)

	if bannerID > 0 {
		query = query.Where("id = ?", bannerID)
	}

	err = query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
