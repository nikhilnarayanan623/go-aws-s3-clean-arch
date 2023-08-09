//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api/handler"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/config"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/db"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/repository"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/service"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {

	wire.Build(
		db.ConnectDatabase,
		repository.NewFileRepository,
		service.NewUploadService,
		usecase.NewFileUseCase,
		handler.NewFileHandler,
		http.NewServerHTTP,
	)

	return &http.Server{}, nil
}
