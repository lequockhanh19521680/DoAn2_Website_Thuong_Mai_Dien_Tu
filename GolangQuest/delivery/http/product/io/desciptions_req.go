package io

import "mime/multipart"

type ProductDescriptionsCreateReq struct {
	ID        uint
	ProductID uint                  `form:"product_id"`
	Name      string                `form:"name"`
	File      *multipart.FileHeader `form:"descriptions_md"`
}

type ProductDescriptionsUpdateReq struct {
	Name *string               `form:"name"`
	File *multipart.FileHeader `form:"descriptions_path"`
}
