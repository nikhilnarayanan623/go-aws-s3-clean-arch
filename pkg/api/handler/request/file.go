package request

import "mime/multipart"

type SingleFile struct {
	Name        string                `json:"name" binding:"required"`
	Description string                `json:"description" binding:"required"`
	FileHeader  *multipart.FileHeader `json:"-"`
}
