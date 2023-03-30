package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vasolovev/secret_santa/internal/entity"
	"github.com/vasolovev/secret_santa/internal/usecase"
	"github.com/vasolovev/secret_santa/pkg/logger"
)

type groupRoutes struct {
	u usecase.Groups
	l logger.Interface
}

func newGroupRoutes(handler *gin.RouterGroup, u usecase.Groups, l logger.Interface) {
	r := &groupRoutes{u, l}

	h := handler.Group("/group")
	{
		h.POST("/group", r.createGroup)
		h.GET("/groups", r.getAllGroups)
	}
}

type inpGroup struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// @Summary Create Group
// @Description Create a new group
// @Tags         Group
// @ID create-group
// @Accept json
// @Produce json
// @Param Group body inpGroup true "Group object to be created"
// @Success 201 {object} int "Created Group"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Not Found"
// @Router /address [post]
func (r *groupRoutes) createGroup(c *gin.Context) {
	var group inpGroup
	if err := c.ShouldBindJSON(&group); err != nil {
		r.l.Error(err, "http - v1 - Address")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	id, err := r.u.Create(c.Request.Context(), entity.Group{
		Name:        group.Name,
		Description: group.Description,
	})
	if err != nil {
		r.l.Error(err, "http - v1 - Group")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusCreated, id)
}

// @Summary Get all groups
// @Description Get all groups
// @Tags         Group
// @ID get-all-groups
// @Produce json
// @Success 200 {array} entity.Address "List of Address"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Not Found"
// @Router /address [get]
func (r *groupRoutes) getAllGroups(c *gin.Context) {
	Addresss, err := r.u.GetAll(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - Group")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, Addresss)
}
