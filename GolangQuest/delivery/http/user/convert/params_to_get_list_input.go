package convert

import (
	ioHandler "github.com/eNViDAT0001/Backend/delivery/http/user/io"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/jinzhu/copier"
)

func ParamsToGetListInput(input *ioHandler.GetUserListParams) (*paging.GetListInput, error) {
	var result paging.GetListInput
	err := copier.Copy(&result, input)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
