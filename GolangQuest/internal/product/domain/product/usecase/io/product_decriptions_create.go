package io

import "mime/multipart"

type ProductDescriptionsCreate struct {
	ProductID uint
	Name      string
	File      *multipart.FileHeader
}
