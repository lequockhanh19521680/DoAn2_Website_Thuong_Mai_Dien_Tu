package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/internal/user/entities"
)

func (u userStorage) DeleteUserByIDs(ctx context.Context, IDs []uint) error {
	tableName := entities.User{}.TableName()
	db := wrap_gorm.GetDB()

	err := db.Table(tableName).Where("id IN ?", IDs).Delete(&entities.User{}).Error

	if err != nil {
		return err
	}

	return nil
}
