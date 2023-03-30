package v1

import (
	"net/http"

	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

type groupRoutes struct {
	t usecase.group
	l logger.Interface
}

func newgroupRoutes(handler *gin.RouterGroup, t usecase.group, l logger.Interface) {
	r := &groupRoutes{t, l}

	h := handler.Group("/group")
	{
		h.GET("/history", r.history)
		h.POST("/do-translate", r.doTranslate)
	}
}

type historyResponse struct {
	History []entity.group `json:"history"`
}

// @Summary     Show history
// @Description Show all group history
// @ID          history
// @Tags  	    group
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /group/history [get]
func (r *groupRoutes) history(c *gin.Context) {
	groups, err := r.t.History(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, historyResponse{groups})
}

type doTranslateRequest struct {
	Source      string `json:"source"       binding:"required"  example:"auto"`
	Destination string `json:"destination"  binding:"required"  example:"en"`
	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    group
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up group"
// @Success     200 {object} entity.group
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /group/do-translate [post]
func (r *groupRoutes) doTranslate(c *gin.Context) {
	var request doTranslateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	group, err := r.t.Translate(
		c.Request.Context(),
		entity.group{
			Source:      request.Source,
			Destination: request.Destination,
			Original:    request.Original,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusInternalServerError, "group service problems")

		return
	}

	c.JSON(http.StatusOK, group)
}
