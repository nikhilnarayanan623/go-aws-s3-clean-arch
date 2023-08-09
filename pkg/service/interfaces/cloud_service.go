package interfaces

import (
	"context"
	"mime/multipart"
)

type CloudService interface {
	UploadOne(ctx context.Context, fileHeader *multipart.FileHeader) (uploadID string, err error)
	UploadMany(ctx context.Context, filesHeaders *[]multipart.FileHeader) (uploadIDs []string, err error)
	GetOneUrl(ctx context.Context, uploadID string) (url string, err error)
}
