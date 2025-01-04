package scrapper

import (
	"TCGA-storage/storage"
	"fmt"
	"os"
	"os/exec"
)

const prog = "PPPK_Scrapper"
const path = ""
const data = "Data"

func Run() error {

	fmt.Printf("Scrapper started\n")
	cmd := exec.Command("./" + path + prog)
	_, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Scrapper failed\n")
		return err
	}

	fmt.Printf("Scrapper succeded\n")
	return nil
}

func Upload() error {

	minio := storage.New()
	files, err := os.ReadDir(path + data)
	if err != nil {
		return err
	}

	var errors error = nil

	for _, file := range files {
		f, err := os.Open(path + file.Name())

		if err != nil {
			fmt.Printf("Failed to read file %s\n", file.Name())

			errors = fmt.Errorf("%s\n%s", errors, err)
			continue
		}

		stats, _ := file.Info()

		err = minio.UploadFile(f, stats.Size())

		if err != nil {
			fmt.Printf("Failed to Upload file %s\n", file.Name())

			errors = fmt.Errorf("%s\n%s", errors, err)
			continue
		}

	}

	return errors
}
