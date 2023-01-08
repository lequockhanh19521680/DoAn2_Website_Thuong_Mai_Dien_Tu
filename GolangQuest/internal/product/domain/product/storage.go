package product

import (
	"context"
	"github.com/eNViDAT0001/Backend/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/Backend/internal/product/entities"
)

type Storage interface {
	ListProductsPreview(ctx context.Context, input io.ListProductInput) ([]io.ProductPreviewItem, error)
	ListCountProductsPreview(ctx context.Context, input io.ListProductInput) (total int64, err error)

	ListCountProduct(ctx context.Context, input io.ListProductInput) (total int64, err error)
	ListProduct(ctx context.Context, input io.ListProductInput) ([]entities.Product, error)

	ListProductMediaByProductID(ctx context.Context, productID uint) ([]entities.ProductMedia, error)

	GetProductSpecificationRoofByProductID(ctx context.Context, productID uint) ([]entities.ProductSpecification, error)
	GetRoofSpecificationByProductID(ctx context.Context, productID uint, specID *uint) (entities.ProductSpecification, error)
	GetSpecificationByID(ctx context.Context, id uint) (entities.ProductSpecification, error)
	GetSpecificationTreeByProductID(ctx context.Context, productID uint) ([]io.ProductSpecificationFullDetail, error)
	GetProductDescriptionsByProductID(ctx context.Context, productID uint) ([]entities.ProductDescriptions, error)
	GetProductDetailByID(ctx context.Context, productID uint) (entities.Product, error)

	UpdateProduct(ctx context.Context, productID uint, product io.ProductUpdateForm) error
	UpdateProductDescriptions(ctx context.Context, descriptionsID uint, descriptions io.ProductDescriptionsUpdateForm) error
	UpdateProductOptions(ctx context.Context, optionsID uint, options io.ProductOptionUpdateInput) error
	UpdateProductSpecification(ctx context.Context, specID uint, specifications io.ProductSpecificationUpdateForm) error

	CreateProduct(ctx context.Context, product io.ProductCreateForm) (productID uint, err error)
	CreateProductDescriptions(ctx context.Context, descriptions io.ProductDescriptionsCreateForm) (productID uint, err error)
	CreateProductMedia(ctx context.Context, media []io.CreateProductMedia) error
	CreateProductOptions(ctx context.Context, input []io.ProductOptionCreateForm) error
	CreateSpecification(ctx context.Context, input io.ProductSpecificationCreateForm) (specID uint, err error)

	DeleteProductByIDs(ctx context.Context, IDs []uint) error
	DeleteProductByID(ctx context.Context, ID uint) error
}
