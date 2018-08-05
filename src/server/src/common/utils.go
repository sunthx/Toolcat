package common

import (
	"net/http"
	"strings"
	"os"
	"io/ioutil"
)

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


func SaveFile(dir string,fileName string,data []byte) (bool,error){
	exist,err := PathExists(dir)
	if err == nil && !exist{
		os.Mkdir(dir,os.ModePerm)
	}

	filePath := dir + "/" + fileName
	os.Create(filePath)
	ioutil.WriteFile(dir + "/"+fileName,data, os.ModePerm)

	return true,nil
}