package interfaces

import (
	"context"
	"mime/multipart"
)

type UploadService interface {
	UploadOne(ctx context.Context, fileHeader *multipart.FileHeader) (uploadID string, err error)
	UploadMany(ctx context.Context, filesHeaders *[]multipart.FileHeader) (uploadIDs []string, err error)
	// DownloadOne(ctx context.Context)error
	// DownloadMany(ctx context.Context)error
}
