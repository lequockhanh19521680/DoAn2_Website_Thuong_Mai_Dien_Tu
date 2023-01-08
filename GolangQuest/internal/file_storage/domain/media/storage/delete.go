package storage

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/eNViDAT0001/Backend/config/cloudinary"
)

func (m mediaStorage) DeleteMedia(ctx context.Context, publicID string) (*uploader.DestroyResult, error) {
	ms := cloudinary.GetMediaServer()
	result, err := ms.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return nil, err
	}

	return result, err
}
