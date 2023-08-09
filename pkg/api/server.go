package http

import (
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/api/routes"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/go-aws-s3-clean-arch/cmd/api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine *gin.Engine
	port   string
}

func NewServerHTTP(cfg config.Config, fileHandler interfaces.FileHandler) *Server {

	engine := gin.New()
	engine.Use(gin.Logger())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	docs.SwaggerInfo.BasePath = "/file"

	api := engine.Group("/file")

	routes.SetupFileRoutes(api, fileHandler)

	return &Server{
		engine: engine,
		port:   cfg.Port,
	}
}

func (c *Server) Start() {
	c.engine.Run(c.port)
}
