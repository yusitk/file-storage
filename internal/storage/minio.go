package storage

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

const bucketName = "file-storage"

func InitMinIO() {
	endpoint := "localhost:9000"
	accessKey := "admin"
	secretKey := "password"

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatal("Failed to connect to MinIO:", err)
	}

	MinioClient = client

	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatal("Failed to check bucket:", err)
	}

	if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatal("Failed to create bucket:", err)
		}
		log.Println("Bucket created:", bucketName)
	} else {
		log.Println("Bucket already exists:", bucketName)
	}
}
