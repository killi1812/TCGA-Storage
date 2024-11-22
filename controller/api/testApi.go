package api

import (
	"TCGA-storage/storage"
	"encoding/json"
	"net/http"
)

type TestController struct {
}

func NewTestController() *TestController {
	return &TestController{}
}

func (this *TestController) RegisterEndpoints() error {
	http.HandleFunc("/api/ping-minio", func(w http.ResponseWriter, r *http.Request) {
		s := storage.MinioStorage{}
		rez := s.CheckBucket("test")
		json.NewEncoder(w).Encode(rez)
	})

	return nil
}
