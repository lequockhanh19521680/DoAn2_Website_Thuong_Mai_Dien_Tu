package order

import (
	"context"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *orderHandler) DeleteOrder() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		orderID, err := strconv.Atoi(cc.Param("order_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		err = s.orderUC.DeleteOrder(newCtx, uint(orderID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete order success")
	}
}
