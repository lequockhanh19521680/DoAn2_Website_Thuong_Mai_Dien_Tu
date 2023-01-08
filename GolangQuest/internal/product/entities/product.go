package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProviderID uint   `gorm:"column:provider_id"`
	CategoryID uint   `gorm:"column:category_id"`
	UserID     uint   `gorm:"column:user_id"`
	Name       string `gorm:"column:name"`
	Price      string `gorm:"column:price"`
	Discount   int    `gorm:"column:discount"`
}

func (Product) WithFields() []string {
	return []string{"name", "discount", "provider_id", "user_id", "category_id"}
}
func (Product) SearchFields() []string {
	return []string{"name", "price", "id"}
}
func (Product) SortFields() []string {
	return []string{"name", "discount", "price", "provider_id", "user_id", "category_id", "id"}
}
func (Product) TableName() string {
	return "Product"
}
