package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"

	"github.com/eNViDAT0001/Backend/internal/address/entities"
)

func (a *addressStorage) GetWardWithDistrictCode(ctx context.Context, districtCode string) ([]entities.Ward, error) {
	result := make([]entities.Ward, 0)

	tableName := entities.Ward{}.TableName()
	db := wrap_gorm.GetDB()
	//TODO: Need Join To Get AdministrativeRegion And AdministrativeUnit

	err := db.Table(tableName).
		Where("district_code = ?", districtCode).
		Where("deleted_at IS NULL").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
