package storage

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/eNViDAT0001/Backend/config/cloudinary"
	"mime/multipart"
)

func (m mediaStorage) UploadMedia(ctx context.Context, files []*multipart.FileHeader, folder cloudinary.MediaFolderType) ([]*uploader.UploadResult, error) {
	ms := cloudinary.GetMediaServer()
	result := make([]*uploader.UploadResult, 0)
	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			return nil, err
		}
		resp, err := ms.Upload.Upload(ctx, f, uploader.UploadParams{
			Folder:      cloudinary.GetMediaFolder(folder),
			DisplayName: file.Filename,
			EagerAsync:  api.Bool(true),
		})
		if err != nil {
			return nil, err
		}
		result = append(result, resp)
		err = f.Close()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
