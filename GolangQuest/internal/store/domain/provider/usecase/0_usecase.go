package usecase

import (
	"github.com/eNViDAT0001/Backend/internal/store/domain/provider"
)

type providerUseCase struct {
	providerSto provider.Storage
}

func NewProviderUseCase(
	providerSto provider.Storage,

) provider.UseCase {
	return &providerUseCase{
		providerSto: providerSto,
	}
}
