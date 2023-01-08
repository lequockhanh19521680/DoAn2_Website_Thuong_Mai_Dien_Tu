package io

import "mime/multipart"

type ProductMedia struct {
	Files []*multipart.FileHeader `form:"files"`
}
