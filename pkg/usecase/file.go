package usecase

import (
	"context"

	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/domain"
	repoInterface "github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/repository/interfaces"
	serviceInterface "github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/service/interfaces"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/utils"
)

type fileUseCase struct {
	fileRepo      repoInterface.FileRepository
	uploadService serviceInterface.UploadService
}

func NewFileUseCase(fileRepo repoInterface.FileRepository, uploadService serviceInterface.UploadService) interfaces.FileUseCase {
	return &fileUseCase{
		fileRepo:      fileRepo,
		uploadService: uploadService,
	}
}

func (c *fileUseCase) UploadSingleFile(ctx context.Context, fileDetails request.SingleFile) error {

	uploadID, err := c.uploadService.UploadOne(ctx, fileDetails.FileHeader)

	if err != nil {
		return utils.PrependMessageToError(err, "failed to upload file")
	}

	err = c.fileRepo.Transaction(func(trxRepo repoInterface.FileRepository) error {
		file := domain.File{
			UploadID: uploadID,
		}

		fileID, err := trxRepo.SaveFile(ctx, file)
		if err != nil {
			return utils.PrependMessageToError(err, "failed to save file on database")
		}

		singleFile := domain.SingleFile{
			FileID:      fileID,
			Name:        fileDetails.Name,
			Description: fileDetails.Description,
		}

		err = trxRepo.SaveSingleFile(ctx, singleFile)
		if err != nil {
			return utils.PrependMessageToError(err, "failed to save file details on database")
		}
		return nil
	})

	return err
}
