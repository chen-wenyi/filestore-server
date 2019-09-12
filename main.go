package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)

	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		fmt.Printf("Fail to start server, err: %s", error.Error())
	}
}
