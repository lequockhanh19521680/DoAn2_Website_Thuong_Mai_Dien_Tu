package entities

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryParentID *uint  `gorm:"column:category_parent_id"`
	Name             string `gorm:"column:name"`
	ImagePath        string `gorm:"column:image_path"`
}

func (Category) TableName() string {
	return "Category"
}
