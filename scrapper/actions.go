package scrapper

import (
	"fmt"
	"os"
	"os/exec"
)

func Run() error {
	ok := fileLock.TryLock()
	if !ok {
		return fmt.Errorf("Another operation is in progress try later")
	}
	go func() {

		//err := download()
		//if err != nil {
		// fmt.Errorf("%s", err.Error())
		//}

		err := upload()
		if err != nil {
			fmt.Errorf("%s", err.Error())
		}
		defer fileLock.Unlock()
	}()

	return nil
}

func download() error {
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

func upload() error {

	//TODO: Rework path combinging and path saving

	files, err := os.ReadDir(path + data)
	if err != nil {
		return err
	}

	var errors error = nil
	fmt.Printf("Uploading files\n")
	for _, file := range files {
		f, err := os.Open(path + data + "/" + file.Name())

		if err != nil {
			fmt.Printf("Failed to read file %s, error: %s\n", file.Name(), err.Error())

			errors = fmt.Errorf("%s\n%s", errors, err)
			continue
		}

		stats, _ := file.Info()
		fmt.Printf("%v\n", file.Name())
		err = minioStorage.UploadFile(f, file.Name(), stats.Size())

		if err != nil {
			fmt.Printf("Failed to Upload file %s\n", file.Name())

			errors = fmt.Errorf("%s\n%s", errors, err)
			continue
		}

	}

	fmt.Printf("Finished uploading files\n")
	return errors
}
