package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Backend/internal/product/entities"
)

func (s productStorage) ListProduct(ctx context.Context, input io.ListProductInput) ([]entities.Product, error) {
	result := make([]entities.Product, 0)
	db := wrap_gorm.GetDB()

	query := db.Model(entities.Product{})

	paging.SetPagingQuery(&input.Paging, entities.Product{}.TableName(), query)

	if len(input.ProductIDs) > 0 {
		query = query.Where("Product.id IN ?", input.ProductIDs)
	}
	if len(input.CategoryIDs) > 0 {
		query = query.Where("Product.category_id IN ?", input.CategoryIDs)
	}

	err := query.Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
