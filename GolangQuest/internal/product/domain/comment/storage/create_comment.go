package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/wrap_gorm"
	"github.com/eNViDAT0001/Backend/internal/product/domain/comment/storage/io"
	"github.com/eNViDAT0001/Backend/internal/product/entities"
	"gorm.io/gorm/clause"
)

func (c commentStorage) CreateComment(ctx context.Context, comment io.CreateComment) (commentID uint, err error) {
	db := wrap_gorm.GetDB()
	err = db.Table(entities.Comment{}.TableName()).Clauses(clause.OnConflict{
		DoNothing: true,
	}).
	Create(&comment).Error
	if err != nil {
		return 0, err
	}
	return comment.ID, nil
}
