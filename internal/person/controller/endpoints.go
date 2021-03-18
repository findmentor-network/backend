package controller

import (
	"github.com/findmentor-network/backend/internal/person"
	"github.com/findmentor-network/backend/pkg/pagination"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	repository person.Repository
}

func NewController(repository person.Repository) *Controller {
	return &Controller{repository: repository}
}

func NewHandlers(instance *echo.Echo, r *Controller) {

	grp := instance.Group("/api/v1/")
	grp.GET("/persons", r.GetAll)

}

// @Summary Get Persons
// @Description Get Persons
// @Tags Person
// @Accept json
// @Produce json
// @Param page path string true "Page"
// @Success 200 {object} person.Person
// @Failure 400 id person.Person
// @Router /api/v1/ [get]
func (r Controller) GetAll(c echo.Context) error {

	pg := pagination.NewFromRequest(c.Request(), 20)

	items, err := r.repository.Get(c.Request().Context(), pg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, items)
}
