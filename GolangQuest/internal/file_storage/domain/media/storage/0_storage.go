package storage

import (
	"context"
	"github.com/eNViDAT0001/Backend/config/cloudinary"
	"github.com/eNViDAT0001/Backend/internal/file_storage/domain/media"
	"mime/multipart"
)

type mediaStorage struct {
}

func (m mediaStorage) UpdateMedia(ctx context.Context, file *multipart.FileHeader, folder cloudinary.MediaFolderType) error {
	//TODO implement me
	panic("implement me")
}

func NewMediaStorage() media.Storage {
	return &mediaStorage{}
}
