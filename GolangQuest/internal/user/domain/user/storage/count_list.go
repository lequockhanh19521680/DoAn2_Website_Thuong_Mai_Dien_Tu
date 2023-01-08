package storage

import (
	"context"
	"fmt"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/user/entities"
)

func (u userStorage) CountList(ctx context.Context, input *paging.GetListInput) (int64, error) {
	var count int64
	db := wrap_gorm.GetDB()

	query := db.Model(entities.User{})

	if input.Type == paging.CURSOR_PAGING {
		query = query.Where("id > ?", input.Current())
	}

	if input.Filter.GetFields() != nil {
		for k, v := range *input.Filter.GetFields() {
			query = query.Where(fmt.Sprintf("`User`.`%s` = ?", k), v)
		}
	}

	if input.Filter.GetSearch() != nil {
		for k, v := range *input.Filter.GetSearch() {
			query = query.Where(fmt.Sprintf("`User`.`%s` LIKE ?", k), "%"+v+"%")
		}
	}

	if input.Filter.GetSort() != nil {
		for k, v := range *input.Filter.GetSort() {
			sort := "ASC"
			if v == "DESC" {
				sort = v
			}
			query = query.Order(fmt.Sprintf(`%s %s`, k, sort))
		}
	}

	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
