package entities

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	OptionID uint           `gorm:"column:option_id"`
	OrderID  uint           `gorm:"column:order_id"`
	Quantity int            `gorm:"column:quantity"`
	EndTime  datatypes.Time `gorm:"column:end_time"`
}

func (Reservation) TableName() string {
	return "Reservation"
}
