package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vasolovev/secret_santa/internal/usecase"

	"github.com/vasolovev/secret_santa/pkg/logger"
)

type ParticipantRoutes struct {
	t usecase.Participants
	l logger.Interface
}

func newParticipantRoutes(handler *gin.RouterGroup, t usecase.Participants, l logger.Interface) {
	// r := &ParticipantRoutes{t, l}

	// h := handler.Group("/participant")
	// {
	// }
}
