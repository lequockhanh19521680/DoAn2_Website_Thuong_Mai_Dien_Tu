package paging

type GetListInput struct {
	Marker int        `form:"marker"`
	Limit  int        `form:"limit"`
	Total  int        `form:"total"`
	Type   PagingType `form:"type"`
	Filter FilterList
}

func (s GetListInput) PagingType() PagingType {
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
