package product

import (
	"context"
	"github.com/eNViDAT0001/Backend/api"
	"github.com/eNViDAT0001/Backend/delivery/http/product/convert"
	ioHandler "github.com/eNViDAT0001/Backend/delivery/http/product/io"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s *productHandler) CreateSpecificationTree() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.BadRequest(err)
			return
		}

		var input ioHandler.ProductSpecificationsCreateTreeReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		if len(input.Specification) < 1 {
			cc.ResponseError(api.NewBadRequestError("Specification", nil, "Empty specifications"))
		}

		for i := range input.Specification {
			input.Specification[i].Specification.ProductID = uint(productID)
			for j := 0; j < len(input.Specification[i].Options); j++ {
				input.Specification[i].Options[j].ProductID = uint(productID)
			}
		}

		inputUC, err := convert.CreateSpecificationArrayReqToCreateSpecificationForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		specID, err := s.productUC.CreateSpecification(newCtx, inputUC[0])
		if err != nil {
			cc.ResponseError(err)
			return
		}
		for i := 1; i < len(inputUC); i++ {
			inputUC[i].Specification.SpecificationID = &specID
			id, err := s.productUC.CreateSpecification(newCtx, inputUC[i])
			if err != nil {
				if err == gorm.ErrRegistered {
					cc.Conflict(err)
				}
				cc.ResponseError(err)
				return
			}
			specID = id
		}

		cc.Ok("Create Product Specification Success")
	}
}
