package user

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	CreateClient() func(*gin.Context)
	UpdateClient() func(*gin.Context)
}
