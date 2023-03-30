package v1

import (
	"net/http"

	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

type participantRoutes struct {
	t usecase.participant
	l logger.Interface
}

func newparticipantRoutes(handler *gin.RouterGroup, t usecase.participant, l logger.Interface) {
	r := &participantRoutes{t, l}

	h := handler.Group("/participant")
	{
		h.GET("/history", r.history)
		h.POST("/do-translate", r.doTranslate)
	}
}

type historyResponse struct {
	History []entity.participant `json:"history"`
}

// @Summary     Show history
// @Description Show all participant history
// @ID          history
// @Tags  	    participant
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /participant/history [get]
func (r *participantRoutes) history(c *gin.Context) {
	participants, err := r.t.History(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, historyResponse{participants})
}

type doTranslateRequest struct {
	Source      string `json:"source"       binding:"required"  example:"auto"`
	Destination string `json:"destination"  binding:"required"  example:"en"`
	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    participant
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up participant"
// @Success     200 {object} entity.participant
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /participant/do-translate [post]
func (r *participantRoutes) doTranslate(c *gin.Context) {
	var request doTranslateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	participant, err := r.t.Translate(
		c.Request.Context(),
		entity.participant{
			Source:      request.Source,
			Destination: request.Destination,
			Original:    request.Original,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusInternalServerError, "participant service problems")

		return
	}

	c.JSON(http.StatusOK, participant)
}
