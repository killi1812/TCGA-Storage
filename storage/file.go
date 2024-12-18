package storage

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

type MinioStorage struct {
}

func New() *MinioStorage {
	return &MinioStorage{}
}

func (this *MinioStorage) Upload(file multipart.File, header *multipart.FileHeader) error {
	lock.Lock()
	defer lock.Unlock()
	_, err := MinioClient.PutObject(
		context.Background(),
		"test",
		header.Filename,
		file,
		header.Size,
		minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return fmt.Errorf("Failed to upload \n%s", err.Error())
	}

	return nil
}

func (this *MinioStorage) Download(name string) ([]byte, error) {
	lock.Lock()
	defer lock.Unlock()

	reader, err := MinioClient.GetObject(context.Background(), "test", name, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	info, err := reader.Stat()
	img := make([]byte, info.Size)
	reader.Read(img)
	return img, nil
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
