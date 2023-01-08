package order_item

import (
	"context"
	"github.com/eNViDAT0001/Backend/internal/order/entities"
)

type Storage interface {
	ListByOrderID(ctx context.Context, orderID uint) ([]entities.OrderItem, error)
}
