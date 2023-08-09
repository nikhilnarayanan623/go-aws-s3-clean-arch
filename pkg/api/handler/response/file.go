package response

import "time"

type SingleFile struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UploadedAt  time.Time `json:"uploaded_at"`
	Url         string    `json:"url"`
}
