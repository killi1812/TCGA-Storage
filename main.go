package main

import (
	"TCGA-storage/config"
	"TCGA-storage/controller/api"
	"TCGA-storage/controller/ftp"
	"TCGA-storage/storage"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("App started")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			signal.Ignore(sig)
			cleanup()
		}
	}()

	err := config.LoadConfig()
	if err != nil {
		fmt.Print(err)
		panic(1)
	}

	err = storage.Setup()

	err = config.RegiserControllers([]config.Controller{
		ftp.NewPageController(),
		api.NewTestController(),
	})

	fmt.Printf("Listeing on http://localhost:%s\n", config.GetPort())
	err = http.ListenAndServe(fmt.Sprintf(":%s", config.GetPort()), nil)
	if err != nil {
		fmt.Print(err)
		panic(2)
	}
}

func cleanup() {
	fmt.Println("\nBye 👋")
	os.Exit(0)
}
