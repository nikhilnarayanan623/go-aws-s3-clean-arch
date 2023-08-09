package db

import (
	"fmt"

	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/config"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database : %w", err)
	}

	err = db.AutoMigrate(
		domain.File{},
		domain.SingleFile{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed migrate tables : %w", err)
	}

	return db, nil
}
