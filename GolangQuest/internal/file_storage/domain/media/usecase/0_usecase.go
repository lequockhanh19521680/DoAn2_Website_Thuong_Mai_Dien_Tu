package usecase

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/eNViDAT0001/Backend/config/cloudinary"
	"github.com/eNViDAT0001/Backend/internal/file_storage/domain/media"
	"mime/multipart"
)

type mediaUseCase struct {
	mediaSto media.Storage
}

func (u *mediaUseCase) UploadMedia(ctx context.Context, files []*multipart.FileHeader, folder cloudinary.MediaFolderType) ([]*uploader.UploadResult, error) {
	return u.mediaSto.UploadMedia(ctx, files, folder)
}
func (u *mediaUseCase) DeleteMedia(ctx context.Context, publicID string) (*uploader.DestroyResult, error) {
	return u.mediaSto.DeleteMedia(ctx, publicID)
}

func NewMediaUseCase(mediaSto media.Storage) media.UseCase {
	return &mediaUseCase{mediaSto: mediaSto}
}
