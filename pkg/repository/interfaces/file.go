package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/domain"
)

type FileRepository interface {
	Transaction(trxFunc func(trxRepo FileRepository) error) error
	SaveFile(ctx context.Context, file domain.File) (fileID uint, err error)
	SaveSingleFile(ctx context.Context, file domain.SingleFile) error
}
