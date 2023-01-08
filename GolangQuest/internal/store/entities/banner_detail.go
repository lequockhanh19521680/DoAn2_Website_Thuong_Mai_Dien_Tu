package entities

import (
	"time"
)

type BannerDetail struct {
	BannerID  uint      `gorm:"column:banner_id"`
	ProductID uint      `gorm:"column:product_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

//	type BannerDetail struct {
//		BannerID  uint
//		ProductID uint
//		CreatedAt time.Time
//		UpdatedAt time.Time
//	}
func (BannerDetail) TableName() string {
	return "BannerDetail"
}
