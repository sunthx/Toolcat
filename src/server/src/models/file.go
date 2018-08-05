package models

type File struct {
	Base
	FileId string `json:fileId`
	FileName string `json:fileName`
	Url string `json:url`
}
