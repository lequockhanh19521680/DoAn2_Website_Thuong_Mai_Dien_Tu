package paging

import (
	"fmt"
	"gorm.io/gorm"
)

func SetCountListPagingQuery(input *GetListInput, tableName string, query *gorm.DB) {

	if input.Type == CURSOR_PAGING {
		query = query.Where(fmt.Sprintf("%s.id > ?", tableName), input.Current())
	}

	if input.Filter.GetFields() != nil {
		for k, v := range *input.Filter.GetFields() {
			query = query.Where(fmt.Sprintf("`%s`.`%s` = ?", tableName, k), v)
		}
	}

	if input.Filter.GetSearch() != nil {
		for k, v := range *input.Filter.GetSearch() {
			query = query.Where(fmt.Sprintf("`%s`.`%s` LIKE ?", tableName, k), "%"+v+"%")
		}
	}

	if input.Filter.GetSort() != nil {
		for k, v := range *input.Filter.GetSort() {
			sort := "ASC"
			if v == "DESC" {
				sort = v
			}
			query = query.Order(fmt.Sprintf(`%s.%s %s`, tableName, k, sort))
		}
	}
}
func SetPagingQuery(input *GetListInput, tableName string, query *gorm.DB) {

	query = query.Limit(input.PerPage())
	if input.Type == CURSOR_PAGING {
		query = query.Where(fmt.Sprintf("%s.id > ?", tableName), input.Current())
	} else {
		offset := Offset(input.Current(), input.PerPage())
		query = query.Offset(offset)
	}

	if input.Filter.GetFields() != nil {
		for k, v := range *input.Filter.GetFields() {
			query = query.Where(fmt.Sprintf("`%s`.`%s` = ?", tableName, k), v)
		}
	}

	if input.Filter.GetSearch() != nil {
		for k, v := range *input.Filter.GetSearch() {
			query = query.Where(fmt.Sprintf("`%s`.`%s` LIKE ?", tableName, k), "%"+v+"%")
		}
	}

	if input.Filter.GetSort() != nil {
		for k, v := range *input.Filter.GetSort() {
			sort := "ASC"
			if v == "DESC" {
				sort = v
			}
			query = query.Order(fmt.Sprintf(`%s.%s %s`, tableName, k, sort))
		}
	}
}
