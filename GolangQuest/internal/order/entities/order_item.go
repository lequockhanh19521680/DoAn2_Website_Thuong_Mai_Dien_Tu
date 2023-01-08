package entities

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID         uint   `gorm:"column:order_id"`
	ProviderID      uint   `gorm:"column:provider_id"`
	ProductID       uint   `gorm:"column:product_id"`
	ProductOptionID uint   `gorm:"column:product_option_id"`
	Name            string `gorm:"column:name"`
	Price           int    `gorm:"column:price"`
	Option          string `gorm:"column:option"`
	Quantity        int    `gorm:"column:quantity"`
	Discount        int    `gorm:"column:discount"`
	Image           string `gorm:"column:image"`
}

func (OrderItem) WithFields() []string {
	return []string{"order_id", "category_id", "product_id", "product_option_id"}
}
func (OrderItem) SearchFields() []string {
	return []string{"name", "price", "option", "quantity", "discount"}
}
func (OrderItem) SortFields() []string {
	return []string{"order_id", "category_id", "product_id", "product_option_id", "id"}
}
func (OrderItem) TableName() string {
	return "OrderItem"
}
