package api

import (
	"TCGA-storage/db"
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
	http.HandleFunc("/api/patients", this.getAllPatients)
	http.HandleFunc("/api/patients/delete", this.deleteAllPatients)
	return nil
}

func (this *DataController) getPatientData(w http.ResponseWriter, r *http.Request) {
	tmp := strings.Split(r.URL.Path, "/")
	patientCode := tmp[len(tmp)-1]

	//TODO:Get Data from mongo

	patient, err := db.Read(patientCode)
	if err != nil {
		fmt.Printf("Failed retriving patient data, error: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Get data from files
	p := parser.GetGeneParser()
	files, err := this.store.GetAllReaders()
	if err != nil {
		fmt.Printf("Failed retriving files \n")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(files) == 0 {
		fmt.Printf("No files \n")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	counter := make(chan rune, len(files))
	dataChan := make(chan parser.PatientGenesExpressions, 1)
	for _, file := range files {
		go func() {
			data, err := p.Parse(file, patientCode)
			defer file.Close()
			if err == nil {
				dataChan <- data
			}
			counter <- 'a'
		}()
	}
	for i := 0; i < len(files); i++ {
		<-counter
	}

	if len(dataChan) > 0 {
		dto := dto.NewPatientGensDto(patient, <-dataChan)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(dto)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
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
		fmt.Printf("Dailed parsing \n")
		return
	}
	//TODO: insertsy empty

	err = db.InsertMany(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Failed Saving to mongo \n")
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("Succesfly inserted %d patients\n", len(data)))
}

func (this *DataController) getAllPatients(w http.ResponseWriter, r *http.Request) {
	data, err := db.ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Eeror reading patients\n")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (this *DataController) deleteAllPatients(w http.ResponseWriter, r *http.Request) {
	err := db.DeleteAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Eeror reading patients\n")
		return
	}

	w.WriteHeader(http.StatusOK)
}
