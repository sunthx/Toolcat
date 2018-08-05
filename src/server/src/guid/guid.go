package guid

import (
		"time"
		"github.com/google/uuid"
	"net/http"
	"encoding/json"
		"models"
)

func New() models.Guid{
	newGuid := uuid.Must(uuid.NewRandom())
	result := models.Guid{
		Content:newGuid.String(),
	}

	result.Base.Date = time.Now().String()
	return result
}

func GuidHandler(writer http.ResponseWriter,request *http.Request) {

	newGuid := New()
	value,err := json.Marshal(&newGuid)
	if err != nil{
		return
	}

	writer.Header().Set("Content-Type","application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write(value)
	return
}

