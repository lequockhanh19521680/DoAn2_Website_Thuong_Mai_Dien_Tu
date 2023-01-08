package category

import (
	"context"
	"github.com/eNViDAT0001/Backend/external/request"
	"github.com/eNViDAT0001/Backend/internal/store/domain/category"
	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryUC category.UseCase
}

func (s *categoryHandler) ListCategories() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		categories, err := s.categoryUC.ListCategories(newCtx)
		if err != nil {
			cc.ResponseError(err)
		}
		cc.Ok(categories)
	}
}

func NewCategoryHandler(categoryUC category.UseCase) category.HttpHandler {
	return &categoryHandler{categoryUC: categoryUC}
}
