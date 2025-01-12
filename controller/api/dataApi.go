package api

import (
	"TCGA-storage/dto"
	"TCGA-storage/parser"
	"TCGA-storage/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type DataController struct {
	store *storage.MinioStorage
}

func NewDataController() *DataController {
	return &DataController{store: storage.New()}
}

func (this *DataController) RegisterEndpoints() error {
	http.HandleFunc("/api/ping-mongo", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Test")
	})
	http.HandleFunc("/api/data/upload", this.upload)
	http.HandleFunc("/api/patient/data/", this.getPatientData)
	return nil
}

func (this *DataController) getPatientData(w http.ResponseWriter, r *http.Request) {
	tmp := strings.Split(r.URL.Path, "/")
	patientCode := tmp[len(tmp)-1]

	//Get Data from mongo

	patient := parser.PatientData{}

	//Get data from files
	p := parser.GetGeneParser()
	files, err := this.store.GetAllReaders(patientCode)
	if err != nil {
		fmt.Printf("failed retriving files \n")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(files) == 0 {
		fmt.Printf("no files \n")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	counter := make(chan rune, len(files))
	dataChan := make(chan parser.PatientGenesExpressions, 1)
	for _, file := range files {
		go func() {
			data, err := p.Parse(file, patientCode)
			defer file.Close()
			counter <- 'a'
			if err == parser.PatientNotFound {
				return
			}
			dataChan <- data
		}()
	}
	for i := 0; i < len(files); i++ {
		<-counter
	}

	dto := dto.NewPatientGensDto(patient, <-dataChan)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto)
}

func (this *DataController) upload(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("filename")
	if err != nil {
		fmt.Printf("Error reading file \n\t\n")
		return
	}
	defer file.Close()

	p := parser.GetPatientParser()

	data, err := p.Parse(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("failed parsing \n")
		return
	}

	writer := strings.Builder{}
	json.NewEncoder(&writer).Encode(data)
	fmt.Printf("writer: %v\n", writer.String())
}
