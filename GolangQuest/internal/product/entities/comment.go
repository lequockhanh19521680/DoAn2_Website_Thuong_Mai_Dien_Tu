package entities

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ProductID    uint   `gorm:"column:product_id"`
	UserID       uint   `gorm:"column:user_id"`
	Name         string `gorm:"column:name"`
	Descriptions string `gorm:"column:descriptions"`
	Rating       int    `gorm:"column:rating"`
}

func (Comment) WithFields() []string {
	return []string{"rating", "description"}
}
func (Comment) SearchFields() []string {
	return []string{"description"}
}
func (Comment) SortFields() []string {
	return []string{"rating", "description", "id"}
}

func (Comment) TableName() string {
	return "Comment"
}
