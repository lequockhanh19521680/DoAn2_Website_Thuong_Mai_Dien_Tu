package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Backend/internal/product/entities"
)

func (s productStorage) CreateProductOptions(ctx context.Context, input []io.ProductOptionCreateForm) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.ProductOption{}.TableName()).
		Create(&input).Error
	if err != nil {
		return err
	}
	return err
}
