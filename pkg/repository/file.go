package repository

import (
	"context"
	"time"

	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/domain"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/utils"

	"gorm.io/gorm"
)

type fileDatabase struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) interfaces.FileRepository {
	return &fileDatabase{
		db: db,
	}
}

func (c *fileDatabase) Transaction(trxFunc func(trxRepo interfaces.FileRepository) error) error {

	trxDB := c.db.Begin()

	trxRepo := NewFileRepository(trxDB)

	err := trxFunc(trxRepo)
	if err != nil {
		trxDB.Rollback()
		return err
	}

	err = trxDB.Commit().Error
	if err != nil {
		return utils.PrependMessageToError(err, "failed to commit transaction")
	}

	return nil
}

func (c *fileDatabase) SaveFile(ctx context.Context, file domain.File) (fileID uint, err error) {

	query := `INSERT INTO files (upload_id) VALUES($1) RETURNING id`
	err = c.db.Raw(query, file.UploadID).Scan(&fileID).Error

	return
}

func (c *fileDatabase) SaveSingleFile(ctx context.Context, file domain.SingleFile) error {

	query := `INSERT INTO single_files (file_id, name, description, uploaded_at) 
	VALUES ($1, $2, $3, $4)`

	err := c.db.Exec(query, file.FileID, file.Name, file.Description, time.Now()).Error

	return err
}
