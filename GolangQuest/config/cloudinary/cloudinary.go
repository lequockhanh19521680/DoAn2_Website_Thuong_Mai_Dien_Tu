package cloudinary

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

var (
	mediaServer *cloudinary.Cloudinary
)

func GetMediaServer() *cloudinary.Cloudinary {
	if mediaServer != nil {
		return mediaServer
	}

	cld, err := cloudinary.NewFromParams("damzcas3k", "332611833886276", "CApeUOYt-JIgUSj9LvSAs6rO610")
	if err != nil {
		panic(err)
	}

	cld.Config.URL.Secure = true
	mediaServer = cld
	return mediaServer
}
