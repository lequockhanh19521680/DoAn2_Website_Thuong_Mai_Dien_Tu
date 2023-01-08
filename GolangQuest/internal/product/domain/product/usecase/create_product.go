package usecase

import (
	"context"
	ioUC "github.com/eNViDAT0001/Backend/internal/product/domain/product/usecase/io"
)

func (u *productUseCase) CreateProduct(ctx context.Context, productDetail ioUC.ProductDetailCreateForm) (productID uint, err error) {
	productID, err = u.productSto.CreateProduct(ctx, productDetail.Product)
	if err != nil {
		return 0, err
	}
	err = u.CreateProductMedia(ctx, productID, productDetail.Media)
	if err != nil {
		return 0, err
	}
	productDetail.Descriptions.ProductID = productID
	if productDetail.Descriptions.File != nil {
		_, err = u.CreateDescriptions(ctx, productDetail.Descriptions)
		if err != nil {
			return 0, err
		}
	}
	return productID, nil
}
