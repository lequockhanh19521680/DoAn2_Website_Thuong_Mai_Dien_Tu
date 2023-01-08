package order

import (
	"context"
	"github.com/eNViDAT0001/Backend/delivery/http/order/order/convert"
	io2 "github.com/eNViDAT0001/Backend/delivery/http/order/order/io"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *orderHandler) CreateOrder() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io2.CreateOrderReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		orderSto, itemsSto, err := convert.CreateReqToCreateOrderFormInput(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.orderUC.CreateOrder(newCtx, orderSto, itemsSto, input.CartItemsIDS)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Make Order Success")
	}
}
