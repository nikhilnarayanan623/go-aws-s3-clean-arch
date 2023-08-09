package routes

import (
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api/handler/interfaces"

	"github.com/gin-gonic/gin"
)

func SetupFileRoutes(api *gin.RouterGroup, fileHandler interfaces.FileHandler) {

	upload := api.Group("/uploads")
	{
		upload.POST("/one", fileHandler.UploadSingleFile)
	}

}
