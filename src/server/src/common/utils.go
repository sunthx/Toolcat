package common

import (
	"net/http"
	"strings"
	"os"
	"io/ioutil"
	"guid"
)

var fileStoragePath = "./public/files"

func GetRequestUrl(request *http.Request) string{
	scheme := "http://"
	if request.TLS != nil {
		scheme = "https://"
	}

	return strings.Join([]string{scheme, request.Host, request.RequestURI},"")
}

func PathExists(path string)(bool,error){
	_, err := os.Stat(path)
	if err == nil{
		return true, nil
	}

	if os.IsNotExist(err){
		return false,nil
	}

	return false,err
}

func SaveFile(fileName string,data []byte) (bool,error){
	exist,err := PathExists(fileStoragePath)
	if err == nil && !exist{
		os.Mkdir(fileStoragePath,os.ModePerm)
	}

	filePath := fileStoragePath + "/" + fileName
	newFile,_ := os.Create(filePath)
	defer newFile.Close()

	ioutil.WriteFile(fileStoragePath + "/"+fileName,data, os.ModePerm)
	return true,nil
}

func GetRandomId() string{
	return guid.New().Content
}

func GetRandomFilePath() string{
	return GetFullFilePath(guid.New().Content)
}

func GetFullFilePath(fileId string) string{
	return fileStoragePath + "/" + fileId
}