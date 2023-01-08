package usecase

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/cloudinary"
	"github.com/eNViDAT0001/Backend/internal/product/domain/product/storage/io"
	ioUC "github.com/eNViDAT0001/Backend/internal/product/domain/product/usecase/io"
	"mime/multipart"
)

func (u *productUseCase) CreateDescriptions(ctx context.Context, input ioUC.ProductDescriptionsCreate) (newID uint, err error) {
	files := []*multipart.FileHeader{input.File}
	uploadedFiles, err := u.mediaSto.UploadMedia(ctx, files, cloudinary.ProductDescriptions)
	if err != nil {
		return newID, err
	}
	descriptions := io.ProductDescriptionsCreateForm{
		ProductID:        input.ProductID,
		Name:             input.Name,
		PublicID:         uploadedFiles[0].PublicID,
		DescriptionsPath: uploadedFiles[0].URL,
	}
	newID, err = u.productSto.CreateProductDescriptions(ctx, descriptions)
	if err != nil {
		return 0, err
	}
	return newID, nil
}
