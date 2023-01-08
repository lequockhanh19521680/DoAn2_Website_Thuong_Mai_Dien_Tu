package user

import (
	"context"

	"github.com/eNViDAT0001/Backend/delivery/http/user/convert"
	"github.com/eNViDAT0001/Backend/delivery/http/user/io"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s userHandler) CreateUser() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.CreateUserReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		inputSto, err := convert.CreateReqToCreateUserInput(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		userID, err := s.userUC.CreateUser(newCtx, inputSto)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(userID)
	}
}
