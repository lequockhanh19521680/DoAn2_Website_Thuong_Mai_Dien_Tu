package comment

import (
	"context"
	"github.com/eNViDAT0001/Backend/api"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/eNViDAT0001/Backend/internal/product/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s commentHandler) ListCommentByProductID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator := paging.GetListInput{}
		if err := cc.BindQuery(&paginator); err != nil {
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

		inValidField, val := paging.ValidateFilter(paginator.Filter, entities.Comment{})
		if len(inValidField) > 0 {
			cc.ResponseError(api.NewBadRequestError(inValidField, val, "invalid key and value"))
			return
		}

		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		comments, total, err := s.commentUC.ListCommentByProductID(newCtx, uint(productID), paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CURSOR_PAGING && len(comments) > 0 {
			paginator.Marker = int(comments[len(comments)-1].ID)
		}

		cc.OkPaging(paginator, comments)
	}
}
