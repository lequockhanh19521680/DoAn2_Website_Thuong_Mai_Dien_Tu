package io

import "mime/multipart"

type DeleteProductReq struct {
	IDs []uint `json:"ids"`
}
type ProductUpdateReq struct {
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
	Price      *int   `json:"price"`
	Discount   *int   `json:"discount"`
}

type ProductCreateReq struct {
	UserID     uint
	ProviderID uint

	CategoryID       uint                    `form:"category_id"`
	Name             string                  `form:"name"`
	Discount         int                     `form:"discount"`
	Price            int                     `form:"price"`
	Media            []*multipart.FileHeader `form:"media"`
	DescriptionsName string                  `form:"descriptions_name"`
	File             *multipart.FileHeader   `form:"descriptions_md"`
}
