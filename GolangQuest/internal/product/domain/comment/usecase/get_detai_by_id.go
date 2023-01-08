package usecase

import (
	"context"
	ioSto "github.com/eNViDAT0001/Backend/internal/product/domain/comment/storage/io"
)

func (u *commentUseCase) GetCommentDetailByID(ctx context.Context, commentID uint) (ioSto.CommentDetail, error) {

	return u.commentSto.GetCommentDetailByID(ctx, commentID)
}
