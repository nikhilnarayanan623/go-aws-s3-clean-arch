package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api/handler/request"
)

type FileUseCase interface {
	UploadSingleFile(ctx context.Context, fileDetails request.SingleFile) error
}
