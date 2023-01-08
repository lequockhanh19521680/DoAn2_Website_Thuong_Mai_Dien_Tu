package cart

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gin"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/eNViDAT0001/Backend/internal/cart/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s *cartHandler) ListCartByUserID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		paginator, err := wrap_gin.GetPagingParams(cc.Context, entities.CartDetail{})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		carts, total, err := s.cartUC.ListCartByUserID(newCtx, uint(userID), paginator)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		paginator.Total = int(total)
		if paginator.Type == paging.CURSOR_PAGING && len(carts) > 0 {
			paginator.Marker = int(carts[len(carts)-1].ID)
		}

		cc.OkPaging(paginator, carts)
	}
}
