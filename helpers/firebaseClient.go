package helpers

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func NewFirebaseStorageClient(ctx context.Context) *storage.BucketHandle {
	opt := option.WithCredentialsFile("firebase-admin-sdk.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v\n", err)
	}

	storageClient, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firebase Storage client: %v\n", err)
	}

	bucketName := os.Getenv("BUCKET_NAME_FIREBASE")
	bucket, err := storageClient.Bucket(bucketName)
	if err != nil {
		log.Fatalf("Error accessing default bucket: %v\n", err)
	}

	return bucket
}
