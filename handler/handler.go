package handler

import (
	"encoding/json"
	"filestore-server/meta"
	"filestore-server/util"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// return html page
		bytes, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "Internal Server Error")
			return
		}
		io.WriteString(w, string(bytes))
	} else if r.Method == http.MethodPost {
		// store file stream to local directory
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data, err: %s", err.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "/tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Failed to create file, err: %s", err.Error())
			return
		}
		defer newFile.Close()
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into file, err: %s", err.Error())
			return
		}

		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)

		fmt.Printf("%v", fileMeta)

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload Finished!")
}

func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form["filehash"][0] // same as r.Form.Get("filehash")
	if fMeta, ok := meta.GetFileMeta(filehash); ok {
		bytes, err := json.Marshal(fMeta)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		io.WriteString(w, string(bytes))
	} else {
		io.WriteString(w, "Cannot find file meta!")
	}
}

func FileQueryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	limitCount, _ := strconv.Atoi(r.Form.Get("limit"))
	fileMetas := meta.GetLatestFileMetas(limitCount)
	bytes, err := json.Marshal(fileMetas)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form.Get("filehash")
	fileMeta, ok := meta.GetFileMeta(filehash)
	if ok {
		file, err := os.Open(fileMeta.Location)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("cannot find file at path: %s", fileMeta.Location)
			return
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("cannot read file at path: %s, err:%s", fileMeta.Location, err.Error())
			return
		}
		w.Header().Set("Content-Disposition", "attachment;filename=\""+fileMeta.FileName+"\"")
		w.Write(data)
	} else {
		io.WriteString(w, "File not found!")
	}
}
