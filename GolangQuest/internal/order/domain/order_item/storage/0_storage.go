package storage

import (
	"github.com/eNViDAT0001/Backend/internal/order/domain/order_item"
)

type orderItemStorage struct {
}

func NewOrderItemStorage() order_item.Storage {
	return &orderItemStorage{}
}
