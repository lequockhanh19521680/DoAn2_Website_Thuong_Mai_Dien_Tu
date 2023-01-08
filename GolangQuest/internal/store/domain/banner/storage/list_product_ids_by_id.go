package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/store/entities"
)

func (b bannerStorage) ListProductIDsByBannerID(ctx context.Context, bannerID uint, filter paging.GetListInput) ([]uint, error) {
	result := make([]uint, 0)

	db := wrap_gorm.GetDB()

	query := db.Table(entities.Banner{}.TableName()).
		Select("BannerDetail.product_id").
		Joins("JOIN BannerDetail ON BannerDetail.banner_id = Banner.id").
		Where("Banner.id = ?", bannerID).
		Where("Banner.deleted_at IS NULL")

	paging.SetPagingQuery(&filter, entities.Banner{}.TableName(), query)

	err := query.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}
