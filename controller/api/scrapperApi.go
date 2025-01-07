package api

import (
	"TCGA-storage/scrapper"
	"TCGA-storage/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

type ScrapperController struct {
	store *storage.MinioStorage
}

func NewScrapperController() *ScrapperController {
	return &ScrapperController{store: storage.New()}
}

func (this *ScrapperController) RegisterEndpoints() error {
	http.HandleFunc("/api/scrape/", this.scrape)
	return nil
}

func (this *ScrapperController) scrape(w http.ResponseWriter, r *http.Request) {
	err := scrapper.Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Scrapping requestfailed with error: %s\n", err.Error())
		json.NewEncoder(w).Encode("Try again later")

	} else {
		w.WriteHeader(http.StatusOK)
	}
}
