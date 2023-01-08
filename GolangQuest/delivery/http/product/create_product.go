package product

import (
	"context"

	"github.com/eNViDAT0001/Backend/delivery/http/product/convert"
	ioHandler "github.com/eNViDAT0001/Backend/delivery/http/product/io"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *productHandler) CreateProduct() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input ioHandler.ProductCreateReq
		if err := cc.ShouldBind(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		userID, _ := strconv.Atoi(cc.Param("user_id"))
		input.UserID = uint(userID)
		providerID, _ := strconv.Atoi(cc.Param("provider_id"))
		input.ProviderID = uint(providerID)

		inputUC, err := convert.CreateProductReqToCreateProductForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		productID, err := s.productUC.CreateProduct(newCtx, inputUC)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		result := map[string]interface{}{
			"ProductID": productID,
		}
		cc.Ok(result)
	}
}
