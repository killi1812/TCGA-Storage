package api

import (
	"TCGA-storage/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type TestController struct {
	store *storage.MinioStorage
}

func NewTestController() *TestController {
	return &TestController{store: storage.New()}
}

func (this *TestController) RegisterEndpoints() error {
	http.HandleFunc("/api/ping-minio", func(w http.ResponseWriter, r *http.Request) {
		rez := this.store.CheckBucket("test")
		json.NewEncoder(w).Encode(rez)
	})
	http.HandleFunc("/api/upload", this.upload)
	http.HandleFunc("/api/img/", this.download)
	return nil
}

func (this *TestController) upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		fmt.Printf("Error reading file \n%s\n", err.Error())
		return
	}
	defer file.Close()

	this.store.Upload(file, fileHeader)
	//TODO redirect to img
}

func (this *TestController) download(w http.ResponseWriter, r *http.Request) {
	//TODO name from url
	path := strings.TrimPrefix(r.URL.Path, "/api/img/")
	id := strings.Split(path, "/")[0]

	bytes, err := this.store.Download(id)
	if err != nil {
		fmt.Printf("Failed to download file\n%s\n", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(bytes)
}
