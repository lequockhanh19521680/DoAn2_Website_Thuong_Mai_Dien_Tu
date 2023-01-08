package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID             uint        `gorm:"column:user_id"`
	ProviderID         uint        `gorm:"column:provider_id"`
	Name               string      `gorm:"column:name"`
	Gender             *bool       `gorm:"column:gender"`
	Phone              string      `gorm:"column:phone"`
	Province           string      `gorm:"column:province"`
	District           string      `gorm:"column:district"`
	Ward               string      `gorm:"column:ward"`
	Street             string      `gorm:"column:street"`
	Quantity           int         `gorm:"column:quantity"`
	Total              int         `gorm:"column:total"`
	Discount           int         `gorm:"column:discount"`
	Status             OrderStatus `gorm:"column:status"`
	StatusDescriptions string      `gorm:"column:status_descriptions"`
}

func (Order) WithFields() []string {
	return []string{"provider_id", "user_id", "name", "gender", "status", "id"}
}
func (Order) SearchFields() []string {
	return []string{"name", "gender", "status", "phone", "province", "district", "ward", "street", "quantity"}
}
func (Order) SortFields() []string {
	return []string{"provider_id", "user_id", "name", "gender", "status", "total", "discount", "id", "province", "district", "ward"}
}
func (Order) TableName() string {
	return "Order"
}

type OrderStatus string

const (
	WaitingOrder    OrderStatus = "WAITING"
	ConfirmedOrder  OrderStatus = "CONFIRMED"
	DeliveringOrder OrderStatus = "DELIVERING"
	DeliveredOrder  OrderStatus = "DELIVERED"
	CancelOrder     OrderStatus = "CANCEL"
)
