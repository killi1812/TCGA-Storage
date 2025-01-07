package scrapper

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Run() error {
	ok := fileLock.TryLock()
	if !ok {
		return fmt.Errorf("Another operation is in progress try later")
	}
	go func() {

		err := download()
		if err != nil {
			fmt.Errorf("%s", err.Error())
		}

		err = upload()
		if err != nil {
			fmt.Errorf("%s", err.Error())
		}

		err = clean()
		if err != nil {
			fmt.Errorf("%s", err.Error())
		}
		defer fileLock.Unlock()
	}()

	return nil
}

func download() error {
	fmt.Printf("Scrapper started\n")
	programPath := filepath.Join(path, prog)
	cmd := exec.Command(programPath)
	_, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Scrapper failed\n")
		return err
	}

	fmt.Printf("Scrapper succeded\n")
	return nil
}

func upload() error {
	workingDir := filepath.Join(data)

	files, err := os.ReadDir(workingDir)
	if err != nil {
		return err
	}

	var errors error = nil
	fmt.Printf("Uploading files\n")
	for _, file := range files {
		f, err := os.Open(filepath.Join(workingDir, file.Name()))

		if err != nil {
			fmt.Printf("Failed to read file %s, error: %s\n", file.Name(), err.Error())

			errors = fmt.Errorf("%s\n%s", errors, err)
			continue
		}

		stats, _ := file.Info()
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

func clean() error {
	workingDir := filepath.Join(data)
	fmt.Printf("Cleaning %s\n", workingDir)

	err := os.RemoveAll(workingDir)
	if err != nil {
		return err
	}

	fmt.Printf("Cleaning %s finished\n", workingDir)
	return nil
}
