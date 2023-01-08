package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Backend/internal/store/domain/category/storage/io"
)

func (u *categoryUseCase) UpdateCategory(ctx context.Context, categoryID uint, input *ioSto.CategoryForm) error {
	return u.categorySto.UpdateCategory(ctx, categoryID, input)
}
