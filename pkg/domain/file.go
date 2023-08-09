package domain

import "time"

type SingleFile struct {
	ID          uint      `json:"id" gorm:"primaryKey;not null"`
	FileID      uint      `json:"file_id"`
	File        File      `json:"file" gorm:"foreignKey:FileID"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	UploadedAt  time.Time `json:"uploaded_at" gorm:"not null"`
}

type File struct {
	ID       uint   `json:"id" gorm:"primaryKey;not null"`
	UploadID string `json:"upload_id" gorm:"not null"`
}
