package comment

import (
	"context"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/product/domain/comment/storage/io"
)

type Storage interface {
	ListCommentByProductID(ctx context.Context, productID uint, filter paging.GetListInput) ([]io.CommentDetail, error)
	CountListCommentByProductID(ctx context.Context, productID uint, filter paging.GetListInput) (int64, error)
	GetCommentDetailByID(ctx context.Context, commentID uint) (io.CommentDetail, error)
	CreateComment(ctx context.Context, comment io.CreateComment) (commentID uint, err error)
	CreateCommentMedia(ctx context.Context, media []io.CreateCommentMedia) error
}
