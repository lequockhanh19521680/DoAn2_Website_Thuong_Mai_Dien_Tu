package entities

import "gorm.io/gorm"

type Provider struct {
	gorm.Model
	UserID    uint   `gorm:"column:user_id"`
	Name      string `gorm:"column:name"`
	ImagePath string `gorm:"column:image_path"`
}

func (Provider) WithFields() []string {
	return []string{"name"}
}
func (Provider) SearchFields() []string {
	return []string{"name"}
}
func (Provider) SortFields() []string {
	return []string{"name", "id"}
}
func (Provider) TableName() string {
	return "Provider"
}
