package models

type ImageInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TagInfo struct {
	ID   int    `json:"id"`
	Pid  string `json:"pid"`
	Name string `json:"name"`
}

type ImageAndTag struct {
	ImageInfo ImageInfo `json:"image-info"`
	TagInfo   TagInfo   `json:"tag-info"`
	TagCount  int       `json:"count"`
}
