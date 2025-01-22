package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iakigarci/go-ddd-microservice-template/config"
	"github.com/iakigarci/go-ddd-microservice-template/internal/adapters/inbound/http/handlers"
)

func NewRouter(config *config.Config) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	indexRoutes := router.Group("/")
	{
		indexRoutes.GET("/health", handlers.HealthCheck)
	}

	router.Run(fmt.Sprintf(":%d", config.HTTP.Port))
	return router
}
