package convert

import (
	ioHttpHandler "github.com/eNViDAT0001/Backend/delivery/http/store/banner/io"
	ioSto "github.com/eNViDAT0001/Backend/internal/store/domain/banner/storage/io"
	"github.com/jinzhu/copier"
)

func UpdateReqToUpdateBannerInput(input *ioHttpHandler.BannerUpdateReq) (ioSto.BannerUpdateForm, error) {
	var result ioSto.BannerUpdateForm
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}
	return result, nil
}
