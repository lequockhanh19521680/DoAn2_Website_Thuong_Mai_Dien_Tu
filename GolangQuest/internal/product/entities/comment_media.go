package entities

import (
	"github.com/eNViDAT0001/Backend/external/enum"
	"gorm.io/gorm"
)

type CommentMedia struct {
	gorm.Model
	CommentID uint           `gorm:"column:comment_id"`
	PublicID  string         `gorm:"column:public_id"`
	MediaPath string         `gorm:"column:media_path"`
	MediaType enum.MediaType `gorm:"column:media_type"`
}

func (CommentMedia) TableName() string {
	return "CommentMedia"
}
