package storage

import (
	"github.com/eNViDAT0001/Backend/internal/product/domain/product"
)

type productStorage struct {
}

func NewProductStorage() product.Storage {
	return &productStorage{}
}
