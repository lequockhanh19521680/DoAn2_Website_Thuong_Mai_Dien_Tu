package entities

import (
	"gorm.io/gorm"
)

type ProductOption struct {
	gorm.Model
	SpecificationID *uint  `gorm:"column:specification_id"`
	ProductID       uint   `gorm:"column:product_id"`
	Name            string `gorm:"column:name"`
	Price           int    `gorm:"column:price"`
	Quantity        int    `gorm:"column:quantity"`
}

func (ProductOption) TableName() string {
	return "ProductOption"
}
