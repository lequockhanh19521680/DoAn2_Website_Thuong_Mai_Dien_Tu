package entities

import "gorm.io/gorm"

type ProductDescriptions struct {
	gorm.Model
	ProductID        uint   `gorm:"column:product_id"`
	Name             string `gorm:"column:name"`
	PublicID         string `gorm:"column:public_id"`
	DescriptionsPath string `gorm:"column:descriptions_path"`
}

func (ProductDescriptions) WithFields() []string {
	return []string{"name", "discount"}
}
func (ProductDescriptions) SearchFields() []string {
	return []string{"name"}
}
func (ProductDescriptions) SortFields() []string {
	return []string{"name", "discount"}
}
func (ProductDescriptions) TableName() string {
	return "ProductDescriptions"
}
