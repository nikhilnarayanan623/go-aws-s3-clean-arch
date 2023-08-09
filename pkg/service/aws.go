package service

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/config"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/service/interfaces"
)

type awsService struct {
	service    *s3.S3
	bucketName string
}

// To get a new file service
func NewUploadService(cfg config.Config) (interfaces.UploadService, error) {

	// create a new session of aws
	session, err := session.NewSession(&aws.Config{
		Region:      &cfg.AwsRegion,
		Credentials: credentials.NewStaticCredentials(cfg.AwsAccessKeyID, cfg.AwsSecretKey, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session for aws s3 error: %w", err)
	}

	// create new aws service using the aws session
	service := s3.New(session)

	return &awsService{
		service:    service,
		bucketName: cfg.AwsBucketName,
	}, nil
}

// To upload one file
func (c *awsService) UploadOne(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {

	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file \nerror:%w", err)
	}

	awsFileID := uuid.New().String()

	_, err = c.service.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: &c.bucketName,
		Key:    &awsFileID,
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file to aws bucket \nerror: %w", err)
	}

	return awsFileID, nil
}

// To upload multiple files
func (c *awsService) UploadMany(ctx context.Context, filesHeaders *[]multipart.FileHeader) (awsFileIDs []string, err error) {
	return nil, nil
}
