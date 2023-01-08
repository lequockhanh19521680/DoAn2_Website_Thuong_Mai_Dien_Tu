package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/internal/store/entities"
)

func (c categoryStorage) ListCategories(ctx context.Context) ([]entities.Category, error) {
	var result []entities.Category
	db := wrap_gorm.GetDB()
	err := db.Model(entities.Category{}).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, err
}
