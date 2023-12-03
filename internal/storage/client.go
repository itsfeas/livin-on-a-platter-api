package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
)

type FireStorage struct {
	*storage.Client
}

var fireStorage FireStorage

func (f *FireStorage) Connect() error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	f.Client = client
	defer client.Close()
	return nil
}

// streamFileUpload uploads an object via a stream.
func (f *FireStorage) streamFileUpload(w io.Writer, buf *bytes.Buffer, bucket, object string) error {
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
	fmt.Fprintf(w, "%v uploaded to %v.\n", object, bucket)
	return nil
}

func (f *FireStorage) GetStorage() *FireStorage {
	return &fireStorage
}
