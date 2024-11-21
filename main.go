package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("App started")
	http.Handle("/", http.FileServer(http.Dir("./wwwroot/")))
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Test")
	})
	fmt.Println("listeing in http://localhost:3050")
	err := http.ListenAndServe(":3050", nil)
	if err != nil {
		print(err)
	}
}
