package storage

import (
	"TCGA-storage/config"
	"context"
	"fmt"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// This will be a singleton
// if there will be a need for a connection pool it will be added
// TODO change this to conn pool after more resarch
var minioClientInstance *minio.Client
var lock sync.Mutex

const bucketName = "patient-data"

func Setup() error {
	fmt.Println("Setting up Minio client")

	// Initialize minio client object.
	minioClient, err := minio.New(config.Conf.MinioConn.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Conf.MinioConn.AccessKeyID, config.Conf.MinioConn.SecretAccessKey, ""),
		Secure: config.Conf.MinioConn.UseSSL,
	})
	if err != nil {
		return fmt.Errorf("Failed to create MinIO client\n%s", err.Error())
	}

	minioClientInstance = minioClient

	lock.Lock()
	defer lock.Unlock()
	ok, err := minioClientInstance.BucketExists(context.Background(), bucketName)
	if err != nil {
		return fmt.Errorf("Error pinging minio service with config(%s, %s)\n%s\n", config.Conf.MinioConn.Endpoint, config.Conf.MinioConn.AccessKeyID, err.Error())
	}

	if !ok {
		fmt.Printf("Making new bucket: %s\n", bucketName)
		err = minioClientInstance.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("Error making new bucket: %s, error: \n%s\n", bucketName, err.Error())
		}
	}

	fmt.Println("MinIO client Setup successfully")
	return nil
}
