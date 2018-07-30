package main

import (
	"./guid"
	"net/http"
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Print("toolcat running ...")
	http.HandleFunc("/guid/new",guidHandleRequest)
	http.ListenAndServe(":8090",nil)
}

func guidHandleRequest(writer http.ResponseWriter,request *http.Request) {

	newGuid := guid.New()
	value,err := json.Marshal(&newGuid)
	if err != nil{
		return
	}

	writer.Header().Set("Content-Type","application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write(value)
	return
}
