package product

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Backend/api"
	"github.com/eNViDAT0001/Backend/delivery/http/product/convert"
	ioHandler "github.com/eNViDAT0001/Backend/delivery/http/product/io"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *productHandler) CreateDescriptions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input ioHandler.ProductDescriptionsCreateReq
		if err := cc.ShouldBind(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		productID, _ := strconv.Atoi(cc.Param("product_id"))
		if input.ProductID != uint(productID) {
			cc.Conflict(api.NewConflictError("Descriptions.ProductID", input.ProductID, "ProductID and Descriptions.ProductID does not match"))
			return
		}

		inputUC, err := convert.CreateDescriptionsReqToCreateDescriptionsForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		descriptionsID, err := s.productUC.CreateDescriptions(newCtx, inputUC)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		result := map[string]interface{}{
			"DescriptionsID": descriptionsID,
		}
		cc.Ok(result)
	}
}
