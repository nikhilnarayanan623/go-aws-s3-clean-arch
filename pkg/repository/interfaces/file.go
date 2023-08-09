package interfaces

import (
	"context"
	"time"

	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/domain"
)

type FileRepository interface {
	Transaction(trxFunc func(trxRepo FileRepository) error) error
	SaveFile(ctx context.Context, file domain.File) (fileID uint, err error)
	SaveSingleFile(ctx context.Context, file domain.SingleFile) error
	FindAllFiles(ctx context.Context, pagination request.Pagination) ([]SingleFileDetails, error)
}

type SingleFileDetails struct {
	ID          uint      `json:"id" gorm:"primaryKey;not null"`
	FileID      uint      `json:"file_id"`
	UploadID    string    `json:"upload_id" gorm:"not null"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	UploadedAt  time.Time `json:"uploaded_at" gorm:"not null"`
}
