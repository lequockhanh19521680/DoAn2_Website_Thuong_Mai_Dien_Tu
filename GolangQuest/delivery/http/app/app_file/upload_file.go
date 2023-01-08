package app_file

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/cloudinary"
	"github.com/eNViDAT0001/Backend/delivery/http/app/app_file/io"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *mediaHandler) UploadMedia() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var form io.FileForm
		err := c.ShouldBind(&form)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		result, err := s.mediaUC.UploadMedia(newCtx, form.File, cloudinary.Product)

		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(result)
	}
}
