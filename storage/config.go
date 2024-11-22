package storage

import (
	"fmt"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

//This will be a singleton
//if there will be a need for a connection pool it will be added

var lock sync.Mutex
var MinioClient *minio.Client

func Setup() error {
	fmt.Println("Setting up Minio client")
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return fmt.Errorf("Failed to create MinIO client\n%s", err.Error())
	}

	MinioClient = minioClient
	fmt.Println("MinIO client Setup successfully ")
	return nil
}
