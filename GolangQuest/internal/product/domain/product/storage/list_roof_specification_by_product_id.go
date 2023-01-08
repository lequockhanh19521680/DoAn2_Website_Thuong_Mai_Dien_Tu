package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/internal/product/entities"
)

func (s productStorage) GetRoofSpecificationByProductID(ctx context.Context, productID uint, specID *uint) (entities.ProductSpecification, error) {

	var result entities.ProductSpecification
	db := wrap_gorm.GetDB()
	query := db.Model(entities.ProductSpecification{}).
		Where("product_id = ?", productID)
	if specID != nil {
		query = query.Where("specification_id = ?", specID)
	} else {
		query = query.Where("specification_id IS NULL")
	}
	err := query.First(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}
