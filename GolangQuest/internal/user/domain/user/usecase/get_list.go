package usecase

import (
	"context"
	"github.com/eNViDAT0001/Backend/external/paging"
	"github.com/eNViDAT0001/Backend/internal/user/entities"
)

func (u userUseCase) GetUserList(ctx context.Context, input *paging.GetListInput) ([]*entities.User, error) {
	total, err := u.userSto.CountList(ctx, input)
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return nil, nil
	}
	result, err := u.userSto.GetUserList(ctx, input)
	if err != nil {
		return nil, err
	}

	input.Total = int(total)

	if input.Type == paging.CURSOR_PAGING {
		input.Marker = int(result[len(result)-1].ID)
	}

	return result, nil
}
