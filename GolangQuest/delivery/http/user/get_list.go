package user

import (
	"context"
	"github.com/eNViDAT0001/Backend/api"
	"github.com/eNViDAT0001/Backend/delivery/http/user/convert"
	"github.com/eNViDAT0001/Backend/delivery/http/user/io"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/eNViDAT0001/Backend/internal/user/entities"
	"github.com/gin-gonic/gin"
)

func (s userHandler) GetUserList() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var paginator io.GetUserListParams
		err := cc.BindQuery(&paginator)
		if err != nil {
			cc.BadRequest(err)
			return
		}

		search := cc.QueryArray("search[]")
		fields := cc.QueryArray("fields[]")
		sort := cc.QueryArray("sorts[]")

		paginator.Filter = paging.NewFilterBuilder().
			WithSearch(search).
			WithFields(fields).
			WithSorts(sort).
			Build()

		inValidField, val := paging.ValidateFilter(paginator.Filter, entities.User{})
		if len(inValidField) > 0 {
			cc.ResponseError(api.NewBadRequestError(inValidField, val, "invalid key and value"))
			return
		}

		inputUC, err := convert.ParamsToGetListInput(&paginator)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		users, err := s.userUC.GetUserList(newCtx, inputUC)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result, err := convert.ArrUserEntityToArrUserRes(users)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		if len(result) == 0 {
			cc.NoContent()
			return
		}

		cc.OkPaging(inputUC, result)
	}
}
