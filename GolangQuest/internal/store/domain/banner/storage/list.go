package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/store/entities"
)

func (b bannerStorage) ListBanner(ctx context.Context, filter paging.GetListInput) ([]entities.Banner, error) {
	result := make([]entities.Banner, 0)

	db := wrap_gorm.GetDB()

	query := db.Model(entities.Banner{})

	paging.SetPagingQuery(&filter, entities.Banner{}.TableName(), query)

	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
