package user

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Backend/delivery/http/user/convert"
	"github.com/eNViDAT0001/Backend/delivery/http/user/io"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s userHandler) AdminUpdateUser() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.UpdateUserReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		inputSto, err := convert.AdminUpdateReqToUpdateUserInput(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		id, _ := strconv.Atoi(cc.Param("user_id"))

		err = s.userUC.UpdateUser(newCtx, uint32(id), inputSto)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update user success")
	}
}
