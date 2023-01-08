package banner

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

func (s bannerHandler) ListProductByBannerID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		bannerID, err := strconv.Atoi(cc.Param("banner_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		paginator := paging.GetListInput{}
		if err := cc.BindQuery(&paginator); err != nil {
			cc.ResponseError(err)
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

		inValidField, val := paging.ValidateFilter(paginator.Filter, entities.Product{})
		if len(inValidField) > 0 {
			cc.ResponseError(api.NewBadRequestError(inValidField, val, "invalid key and value"))
			return
		}

		products, total, err := s.bannerUC.ListProductByBannerID(newCtx, uint(bannerID), paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CURSOR_PAGING && len(products) > 0 {
			paginator.Marker = int(products[len(products)-1].ID)
		}

		cc.OkPaging(paginator, products)
	}
}
