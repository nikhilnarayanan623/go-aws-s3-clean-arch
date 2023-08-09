// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api/handler"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/config"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/db"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/repository"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/service"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	fileRepository := repository.NewFileRepository(gormDB)
	uploadService, err := service.NewUploadService(cfg)
	if err != nil {
		return nil, err
	}
	fileUseCase := usecase.NewFileUseCase(fileRepository, uploadService)
	fileHandler := handler.NewFileHandler(fileUseCase)
	server := http.NewServerHTTP(cfg, fileHandler)
	return server, nil
}