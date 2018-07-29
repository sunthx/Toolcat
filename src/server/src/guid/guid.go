package guid

import (
		"time"
		"../github.com/google/uuid"
)

func New() Guid{
	newGuid := uuid.Must(uuid.NewRandom())
	return Guid{
		Content:newGuid.String(),
		CreatedAt:time.Now().String(),
	}
}

type Guid struct {
	Content string `json:content`
	CreatedAt string `json:createdAt`
}