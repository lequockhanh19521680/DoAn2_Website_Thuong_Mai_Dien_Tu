package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/internal/store/entities"
)

func (s providerStorage) GetProviderByID(ctx context.Context, providerID uint) (entities.Provider, error) {
	var result entities.Provider
	db := wrap_gorm.GetDB()

	err := db.Model(entities.Provider{}).Where("id", providerID).First(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}
