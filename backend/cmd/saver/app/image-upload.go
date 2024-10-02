package app

import (
	"context"

	"github.com/natalie-richards/wine-app/graph/model"
)

func (a *App) UploadImageFromClient(ctx context.Context, req *model.UploadImageRequest) (*string, error) {
	// Take base64 encoded image
	// Connect to Google Cloud Storage
	// Upload image and return public URL on success
	url, err := a.CloudClient.UploadFile(ctx, req.Base64Encoding)
	if err != nil {
		return nil, err
	}
	return url, nil
}
