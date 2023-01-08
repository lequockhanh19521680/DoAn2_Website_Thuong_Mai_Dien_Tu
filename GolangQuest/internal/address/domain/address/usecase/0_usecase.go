package usecase

import (
	address "github.com/eNViDAT0001/Backend/internal/address/domain/address"
)

type addressUseCase struct {
	addressSto address.Storage
}

func NewAddressUseCase(addressSto address.Storage) address.UseCase {
	return &addressUseCase{
		addressSto: addressSto,
	}
}
