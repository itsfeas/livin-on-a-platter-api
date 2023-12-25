package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type FireStorage struct {
	*storage.Client
}

var fireStorage FireStorage

func (f *FireStorage) Connect() error {
	home, err := os.Getwd()
	credFile := os.Getenv("DB_CREDENTIALS")
	if err != nil {
		return err
	}

	ctx := context.Background()

	// Fetch the service account key JSON file contents
	fmt.Println("storage - finding firebase json @: " + home + credFile)
	opt := option.WithCredentialsFile(home + credFile)

	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	f.Client = client
	return nil
}

// streamFileUpload uploads an object via a stream.
func (f *FireStorage) StreamBufferUpload(buf *bytes.Buffer, bucket, object string) error {
	// bucket := "bucket-name"
	// object := "object-name"
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := f.Client.Bucket(bucket).Object(object).NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.

	if _, err := io.Copy(wc, buf); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}
	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %w", err)
	}
	// fmt.Fprintf(w, "%v uploaded to %v.\n", object, bucket)
	return nil
}

// streamFileUpload uploads an object via a stream.
func (f *FireStorage) StreamFileUpload(file *multipart.File, bucket, object string) error {
	// bucket := "bucket-name"
	// object := "object-name"
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := f.Client.Bucket(bucket).Object(object).NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.

	if _, err := io.Copy(wc, *file); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}
	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %w", err)
	}
	// fmt.Fprintf(w, "%v uploaded to %v.\n", object, bucket)
	return nil
}

func GetStorage() *FireStorage {
	return &fireStorage
}
