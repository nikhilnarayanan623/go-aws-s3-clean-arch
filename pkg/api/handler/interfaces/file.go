package interfaces

import "github.com/gin-gonic/gin"

type FileHandler interface {
	UploadSingleFile(ctx *gin.Context)
	GetAllFiles(ctx *gin.Context)
}
