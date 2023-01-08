package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/internal/product/entities"
)

func (s productStorage) DeleteProductByIDs(ctx context.Context, IDs []uint) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.Product{}.TableName()).
		Where("id IN ?", IDs).
		Delete(&entities.Product{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
