package wrap_gin

import (
	"github.com/eNViDAT0001/Backend/api"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/gin-gonic/gin"
)

func GetPagingParams(cc *gin.Context, filter paging.EntityFilter) (paginator paging.GetListInput, err error) {
	if err = cc.BindQuery(&paginator); err != nil {
		return paginator, err
	}

	search := cc.QueryArray("search[]")
	fields := cc.QueryArray("fields[]")
	sort := cc.QueryArray("sorts[]")

	paginator.Filter = paging.NewFilterBuilder().
		WithSearch(search).
		WithFields(fields).
		WithSorts(sort).
		Build()

	inValidField, val := paging.ValidateFilter(paginator.Filter, filter)
	if len(inValidField) > 0 {
		return paginator, api.NewBadRequestError(inValidField, val, "invalid key and value")
	}
	return paginator, err
}
