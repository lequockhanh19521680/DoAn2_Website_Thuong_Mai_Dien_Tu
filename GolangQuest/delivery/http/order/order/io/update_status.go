package io

import "github.com/eNViDAT0001/Backend/internal/order/entities"

type UpdateOrderStatusReq struct {
	Status entities.OrderStatus `json:"status"`
}
