package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api/handler/response"
)

type FileUseCase interface {
	UploadSingleFile(ctx context.Context, fileDetails request.SingleFile) error
	GetAllFiles(ctx context.Context, pagination request.Pagination) ([]response.SingleFile, error)
}
