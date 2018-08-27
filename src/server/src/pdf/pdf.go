package pdf

import (
	"net/http"
	"common"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func PdfFileHandler(writer http.ResponseWriter,request *http.Request){
	convertToPng("")
}

func convertToPng(pdfName string) (string,error)  {
	destPath := common.GetRandomFilePath() + ".jpg"
	imagick.Initialize()
	//imagick.Initialize()
	//defer imagick.Terminate()
	//
	//mw := imagick.NewMagickWand()
	//defer mw.Destroy()
	//
	//if err := mw.SetResolution(300,300); err != nil{
	//	return "",err
	//}
	//
	//if err := mw.ReadImage(pdfName); err != nil {
	//	return "",err
	//}
	//
	//if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_REMOVE); err != nil{
	//	return "",err
	//}
	//
	//if err := mw.SetCompressionQuality(95);err != nil{
	//	return "",err
	//}
	//
	//mw.SetIteratorIndex(0)
	//
	//if err:= mw.SetFormat("jpg");err != nil{
	//	return "",err
	//}
	//
	//if err := mw.WriteImage(destPath); err != nil{
	//	return "",err
	//}

	return destPath,nil
}