package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println(time.Now().String())
	fmt.Print("Toolcat Running ...")

	mux := http.NewServeMux()

	//静态文件
	fileServer := http.FileServer(http.Dir("./public/files/"))
	http.StripPrefix("/static/", fileServer)

	//mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	//mux.HandleFunc("/guid/new", guid.GuidHandler)
	//mux.HandleFunc("/file/upload", file.UploadFileHandler)
	//mux.HandleFunc("/pdf/2img",pdf.PdfFileHandler)

	server := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	server.ListenAndServe()
}
