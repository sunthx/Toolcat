package models

type Guid struct {
	Base
	Content   string `json:content`
}