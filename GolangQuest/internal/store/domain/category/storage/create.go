package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	ioSto "github.com/eNViDAT0001/Backend/internal/store/domain/category/storage/io"
	"github.com/eNViDAT0001/Backend/internal/store/entities"
)

func (c categoryStorage) CreateCategory(ctx context.Context, input *ioSto.CategoryForm) error {
	db := wrap_gorm.GetDB()

	err := db.Table(entities.Category{}.TableName()).Create(input).Error

	if err != nil {
		return err
	}

	return nil
}
