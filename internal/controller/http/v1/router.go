// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/vasolovev/secret_santa/docs"
	"github.com/vasolovev/secret_santa/internal/usecase"
	"github.com/vasolovev/secret_santa/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Secret Santa API
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, u *usecase.UseCases) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Routers
	h := handler.Group("/v1")
	{
		newGroupRoutes(h, u.Group, l)
		newParticipantRoutes(h, u.Participant, l)
	}
}
