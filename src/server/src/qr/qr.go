package qr

import (
	"github.com/boombuler/barcode/qr"
	"os"
	"common"
	"image/png"
	"net/http"
)

func QrHandler(writer http.ResponseWriter,request *http.Request) {
	generate("")
}

func generate(content string) string{
	qrCode, _ := qr.Encode(content, qr.M, qr.Auto)
	fileId := common.GetRandomId()
	file,_ := os.Create(common.GetFullFilePath(fileId))
	defer file.Close()

	png.Encode(file,qrCode)
	return fileId
}
