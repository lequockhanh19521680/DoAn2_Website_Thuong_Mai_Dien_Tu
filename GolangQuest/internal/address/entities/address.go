package entities

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID       *string `gorm:"column:user_id"`
	Name         *string `gorm:"column:name"`
	Gender       *bool   `gorm:"column:gender"`
	Phone        *string `gorm:"column:phone"`
	ProvinceCode *string `gorm:"column:province_code"`
	DistrictCode *string `gorm:"column:district_code"`
	WardCode     *string `gorm:"column:ward_code"`
	Street       *string `gorm:"column:street"`
}

func (Address) TableName() string {
	return "Address"
}
