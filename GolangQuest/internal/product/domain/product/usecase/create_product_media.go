package usecase

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/cloudinary"
	ioSto "github.com/eNViDAT0001/Backend/internal/product/domain/product/storage/io"
	"mime/multipart"
	"strings"
)

func (u *productUseCase) CreateProductMedia(ctx context.Context, productID uint, files []*multipart.FileHeader) error {
	uploadedFiles, err := u.mediaSto.UploadMedia(ctx, files, cloudinary.Product)
	if err != nil {
		return err
	}

	mediaStore := make([]ioSto.CreateProductMedia, 0)
	for _, file := range uploadedFiles {
		ioMediaCreate := ioSto.CreateProductMedia{
			ProductID: productID,
			PublicID:  file.PublicID,
			MediaPath: file.URL,
			MediaType: strings.ToUpper(file.ResourceType),
		}
		mediaStore = append(mediaStore, ioMediaCreate)
	}
	return u.productSto.CreateProductMedia(ctx, mediaStore)
}
