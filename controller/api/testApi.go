package api

import (
	"encoding/json"
	"net/http"
)

type TestController struct {
}

func NewTestController() *TestController {
	return &TestController{}
}

func (this *TestController) RegisterEndpoints() error {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Test")
	})

	return nil
}
