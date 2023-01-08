package provider

import (
	"context"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/store/domain/provider/storage/io"
	"github.com/eNViDAT0001/Backend/internal/store/entities"
)

type Storage interface {
	CreateProvider(ctx context.Context, input io.ProviderForm) (io.ProviderForm, error)
	GetProviderByID(ctx context.Context, provider uint) (entities.Provider, error)
	UpdateProvider(ctx context.Context, providerID uint, input io.ProviderUpdateForm) error
	DeleteProviderByIDs(ctx context.Context, providerID []uint) error
	ListProviderByUserID(ctx context.Context, userID uint, filter paging.GetListInput) ([]entities.Provider, error)
	ListProvider(ctx context.Context) ([]entities.Provider, error)
	CountListProviderByUserID(ctx context.Context, userID uint, filter paging.GetListInput) (total int64, err error)
}
