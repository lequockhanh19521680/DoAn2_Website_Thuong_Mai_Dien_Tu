package storage

import (
	address "github.com/eNViDAT0001/Backend/internal/address/domain/address"
)

type addressStorage struct {
}

func NewAddressStorage() address.Storage {
	return &addressStorage{}
}
