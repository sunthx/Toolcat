package file

import (
	"net/http"
		"io/ioutil"
	"guid"
		"encoding/json"
	"path"
		"common"
	"time"
	"models"
)


func UploadFileHandler(writer http.ResponseWriter, request *http.Request){
	request.ParseMultipartForm(1024*1024*1024)
	file := request.MultipartForm.File["uploaded"][0]

	fileReader,err := file.Open()
	if err != nil {
		return
	}

	data,err := ioutil.ReadAll(fileReader)
	if err != nil {
		return
	}

	fileId := guid.New().Content
	fileExt := path.Ext(file.Filename)
	fileName := fileId + fileExt
	result,err := common.SaveFile(fileName,data)
	if err != nil || !result{
		return
	}

	currentFile := models.File{
		FileId:fileId,
		FileName:fileName,
		Url: request.Host + "/static/" + fileId,
	}

	currentFile.Base.Date =time.Now().String()

	value,err := json.Marshal(&currentFile)
	if err != nil{
		return
	}

	writer.Header().Set("Content-Type","application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write(value)
}




