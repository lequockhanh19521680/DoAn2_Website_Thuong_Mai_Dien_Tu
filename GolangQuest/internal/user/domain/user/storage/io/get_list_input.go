package io

import (
	"github.com/eNViDAT0001/Backend/external/paging"
)

type GetListInput struct {
	Marker int
	Limit  int
	Total  int
	Type   paging.PagingType
	Filter paging.FilterList
}

func (s GetListInput) PagingType() paging.PagingType {
	return s.Type
}
func (s GetListInput) Count() int {
	return s.Total
}
func (s GetListInput) PerPage() int {
	if s.Limit < 1 {
		return 20
	}
	return s.Limit
}
func (s GetListInput) Current() int {
	return s.Marker
}
