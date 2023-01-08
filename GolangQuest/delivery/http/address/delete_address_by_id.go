package address

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s addressHandler) DeleteAddressByID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		userID, _ := strconv.Atoi(cc.Param("user_id"))
		addressID, _ := strconv.Atoi(cc.Param("address_id"))
		ids := []uint{uint(addressID)}
		err := s.addressUC.DeleteAddressByIDs(newCtx, uint(userID), ids)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete Address Success")
	}
}
