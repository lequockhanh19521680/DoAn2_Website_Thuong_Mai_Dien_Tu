package entities

import (
	"github.com/eNViDAT0001/Backend/external/enum"
	"gorm.io/gorm"
)

type ProductMedia struct {
	gorm.Model
	ProductID uint           `gorm:"column:product_id"`
	PublicID  string         `gorm:"column:public_id"`
	MediaPath string         `gorm:"column:media_path"`
	MediaType enum.MediaType `gorm:"column:media_type"`
}

func (ProductMedia) TableName() string {
	return "ProductMedia"
}
