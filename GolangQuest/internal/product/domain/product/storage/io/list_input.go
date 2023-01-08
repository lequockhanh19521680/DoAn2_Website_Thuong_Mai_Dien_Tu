package io

import "github.com/eNViDAT0001/Backend/external/paging"

type ListProductInput struct {
	ProductIDs  []uint
	CategoryIDs []uint
	Paging      paging.GetListInput
}
