package entities

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Banner struct {
	gorm.Model
	UserID     uint           `gorm:"column:user_id"`
	Title      string         `gorm:"column:title"`
	Collection string         `gorm:"column:collection"`
	Discount   string         `gorm:"column:discount"`
	Image      string         `gorm:"column:image"`
	EndTime    datatypes.Date `gorm:"column:end_time"`
}
func (Banner) WithFields() []string {
	return []string{"discount", "collection", "end_time"}
}
func (Banner) SearchFields() []string {
	return []string{"discount", "collection", "end_time"}
}
func (Banner) SortFields() []string {
	return []string{"discount", "collection", "end_time", "id"}
}
func (Banner) TableName() string {
	return "Banner"
}
