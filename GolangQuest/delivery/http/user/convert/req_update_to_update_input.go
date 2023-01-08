package convert

import (
	ioHandler "github.com/eNViDAT0001/Backend/delivery/http/user/io"
	ioSto "github.com/eNViDAT0001/Backend/internal/user/domain/user/storage/io"
	"github.com/jinzhu/copier"
)

func UpdateReqToUpdateUserInput(input *ioHandler.UpdateUserReq) (*ioSto.UpdateUserInput, error) {
	var result ioSto.UpdateUserInput
	err := copier.Copy(&result, input)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func AdminUpdateReqToUpdateUserInput(input *ioHandler.UpdateUserReq) (*ioSto.UpdateUserInput, error) {
	var result ioSto.UpdateUserInput
	err := copier.Copy(&result, input)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
