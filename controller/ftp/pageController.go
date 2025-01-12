package ftp

import (
	"fmt"
	"net/http"
)

type PageController struct {
}

func NewPageController() *PageController {
	return &PageController{}
}

func (this *PageController) RegisterEndpoints() error {
	http.HandleFunc("/app/patient-details/", this.patientDetails)
	http.Handle("/", http.FileServer(http.Dir("./wwwroot/")))
	return nil
}

func (this PageController) patientDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving file\n ")
	http.ServeFile(w, r, "./wwwroot/app/patient-details.html")
}
