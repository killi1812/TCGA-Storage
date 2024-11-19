package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("App started")
	http.Handle("/", http.FileServer(http.Dir("./wwwroot/")))
	fmt.Println("listeing in http://localhost:3050")
	err := http.ListenAndServe(":3050", nil)
	if err != nil {
		print(err)
	}
}
