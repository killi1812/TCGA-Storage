package api

import (
	"TCGA-storage/parser"
	"encoding/json"
	"fmt"
	"net/http"
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

	//fmt.Printf("Started parsig %s", filename)
	p := parser.GetParser()

	data, err := p.Parse(file)
	if err != nil {
		fmt.Printf("failed parsing \n")
	}

	fmt.Printf("parsed data len: %d\n", len(data))
}
