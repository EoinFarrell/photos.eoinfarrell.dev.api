package s3

import (
	"github.com/eoinfarrell/photos.eoinfarrell.dev.api/pkg/models"
	"net/http"
	"net/url"
)

func GetImage(info models.ImageInfo) (resp *http.Response, err error) {
	return getImage(info.Name)
}

func getImage(imageName string) (resp *http.Response, err error) {
	url, _ := url.Parse("https://photos-efarrell.s3.eu-west-1.amazonaws.com/" + imageName)
	escapedUrl := ("https://photos-efarrell.s3.eu-west-1.amazonaws.com" + url.EscapedPath())
	return http.Get(escapedUrl)
}
