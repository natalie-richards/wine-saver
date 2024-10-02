package cloudstorage

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

type storageClient struct {
	*storage.Client
}

type Client interface {
	UploadFile(ctx context.Context, object string) (*string, error)
}

func NewClient(ctx context.Context) Client {
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(fmt.Sprintf("storage.NewClient: %v", err))
	}
	defer client.Close()
	return &storageClient{client}
}

// uploadFile uploads an object.
func (c *storageClient) UploadFile(ctx context.Context, object string) (*string, error) {
	bucket := "bucket-9292024"

	// Open local file.
	f, err := os.Open("notes.txt")
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	o := c.Bucket(bucket).Object(object)

	// Optional: set a generation-match precondition to avoid potential race
	// conditions and data corruptions. The request to upload is aborted if the
	// object's generation number does not match your precondition.
	// For an object that does not yet exist, set the DoesNotExist precondition.
	o = o.If(storage.Conditions{DoesNotExist: true})
	// If the live object already exists in your bucket, set instead a
	// generation-match precondition using the live object's generation number.
	// attrs, err := o.Attrs(ctx)
	// if err != nil {
	//      return fmt.Errorf("object.Attrs: %w", err)
	// }
	// o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	// Upload an object with storage.Writer.
	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return nil, fmt.Errorf("io.Copy: %w", err)
	}
	if err := wc.Close(); err != nil {
		return nil, fmt.Errorf("Writer.Close: %w", err)
	}

	return &object, err
}
