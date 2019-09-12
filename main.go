package main

import (
	"filestore-server/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)

	log.Printf("%s", "listennig at http://localhost:8080")

	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		fmt.Printf("Fail to start server, err: %s", error.Error())
	}
}
