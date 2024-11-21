package main

import (
	"TCGA-storage/config"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("App started")
	err := config.LoadConfig()
	if err != nil {
		fmt.Print(err)
		panic(1)
	}

	http.Handle("/", http.FileServer(http.Dir("./wwwroot/")))
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Test")
	})
	fmt.Printf("listeing in http://localhost:%s\n", config.GetPort())
	err = http.ListenAndServe(fmt.Sprintf(":%s", config.GetPort()), nil)
	if err != nil {
		fmt.Print(err)
		panic(2)
	}
}
