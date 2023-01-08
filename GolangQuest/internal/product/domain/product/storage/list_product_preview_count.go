package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Backend/internal/product/entities"
)

func (s productStorage) ListCountProductsPreview(ctx context.Context, input io.ListProductInput) (total int64, err error) {
	db := wrap_gorm.GetDB()

	query := db.Table(entities.Product{}.TableName()).
		Joins("LEFT JOIN ProductMedia ON ProductMedia.product_id = Product.id").
		Joins("LEFT JOIN Comment ON Comment.product_id = Product.id").
		Where("Product.deleted_at IS NULL").
		Group("Product.id")

	paging.SetCountListPagingQuery(&input.Paging, entities.Product{}.TableName(), query)

	if len(input.ProductIDs) > 0 {
		query = query.Where("Product.id IN ?", input.ProductIDs)
	}
	if len(input.CategoryIDs) > 0 {
		query = query.Where("Product.category_id IN ?", input.CategoryIDs)
	}

	err = query.Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
