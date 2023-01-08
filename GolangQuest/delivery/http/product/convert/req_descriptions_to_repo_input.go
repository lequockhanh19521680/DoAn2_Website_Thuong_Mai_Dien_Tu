package convert

import (
	ioHttpHandler "github.com/eNViDAT0001/Backend/delivery/http/product/io"
	ioUC "github.com/eNViDAT0001/Backend/internal/product/domain/product/usecase/io"
	"github.com/jinzhu/copier"
)

func UpdateDescriptionsReqToUpdateDescriptionsForm(input *ioHttpHandler.ProductDescriptionsUpdateReq) (ioUC.ProductDescriptionsUpdate, error) {
	var result ioUC.ProductDescriptionsUpdate
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}
	return result, nil
}
func CreateDescriptionsReqToCreateDescriptionsForm(input *ioHttpHandler.ProductDescriptionsCreateReq) (ioUC.ProductDescriptionsCreate, error) {
	var result ioUC.ProductDescriptionsCreate
	err := copier.Copy(&result, &input)
	if err != nil {
		return result, err
	}

	return result, nil
}
