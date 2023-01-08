package user

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Backend/delivery/http/user/convert"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s userHandler) GetUserDetailByID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		id, _ := strconv.Atoi(cc.Param("user_id"))

		user, err := s.userUC.GetUserDetailByID(newCtx, uint(id))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result, err := convert.UserEntityToUserRes(user)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(result)
	}
}
