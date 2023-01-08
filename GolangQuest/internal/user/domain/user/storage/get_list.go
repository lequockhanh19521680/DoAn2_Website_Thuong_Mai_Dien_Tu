package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/user/entities"
)

func (u userStorage) GetUserList(ctx context.Context, input *paging.GetListInput) ([]*entities.User, error) {
	result := make([]*entities.User, 0)

	db := wrap_gorm.GetDB()

	query := db.Model(entities.User{})

	paging.SetPagingQuery(input, entities.User{}.TableName(), query)

	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
