package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Backend/internal/product/entities"
)

func (s productStorage) UpdateProductDescriptions(ctx context.Context, descriptionsID uint, descriptions io.ProductDescriptionsUpdateForm) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.ProductDescriptions{}.TableName()).
		Where("id = ?", descriptionsID).
		Updates(&descriptions).Error
	if err != nil {
		return err
	}
	return nil
}
