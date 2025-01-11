package api

import (
	"TCGA-storage/parser"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type DataController struct {
}

func NewDataController() *DataController {
	return &DataController{}
}

func (this *DataController) RegisterEndpoints() error {
	http.HandleFunc("/api/ping-mongo", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Test")
	})
	http.HandleFunc("/api/data/upload", this.upload)
	return nil
}

func (this *DataController) upload(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("filename")
	if err != nil {
		fmt.Printf("Error reading file \n\t\n")
		return
	}
	defer file.Close()

	p := parser.GetTsvParser()

	data, err := p.Parse(file)
	if err != nil {
		fmt.Printf("failed parsing \n")
	}

	writer := strings.Builder{}
	json.NewEncoder(&writer).Encode(data)
	fmt.Printf("writer: %v\n", writer.String())
}
