package storage

import (
	"context"
	"fmt"
	"os"
)

type MinioStorage struct {
}

func New() *MinioStorage {
	return &MinioStorage{}
}

func (this *MinioStorage) Upload(file *os.File) error {
	return nil
}

func (this *MinioStorage) Download(name string) (*os.File, error) {
	return nil, nil
}

func (this *MinioStorage) CheckBucket(name string) bool {
	lock.Lock()
	defer lock.Unlock()
	ret, err := MinioClient.BucketExists(context.Background(), name)
	if err != nil {
		fmt.Printf("Error accessing bucket: %s\n%s\n", name, err.Error())
		return false
	}
	return ret
}
