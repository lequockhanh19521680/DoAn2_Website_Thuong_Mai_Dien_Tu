package provider

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gin"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/eNViDAT0001/Backend/internal/store/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s providerHandler) ListProviderByUserID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		paginator, err := wrap_gin.GetPagingParams(cc.Context, entities.Provider{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		providers, total, err := s.providerUC.ListProviderByUserID(newCtx, uint(userID), paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CURSOR_PAGING && len(providers) > 0 {
			paginator.Marker = int(providers[len(providers)-1].ID)
		}

		cc.OkPaging(paginator, providers)
	}
}
