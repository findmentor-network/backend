package controller

import (
	"fmt"
	"github.com/findmentor-network/backend/internal/person"
	mongohelper "github.com/findmentor-network/backend/pkg/mongoextentions"
	"github.com/findmentor-network/backend/pkg/pagination"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	repository person.Repository
}

func NewController(repository person.Repository) *Controller {
	return &Controller{repository: repository}
}

func NewHandlers(instance *echo.Echo, r *Controller) {

	grp := instance.Group("/api/v1/")
	grp.GET("persons", r.Get)
	grp.GET("persons/peer/:slug", r.GetBySlug)
	grp.GET("persons/filter", r.Get)
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
func (r Controller) Get(c echo.Context) error {

	pg := pagination.NewFromRequest(c.Request(), -1)
	query := mongohelper.QueryOf()
	query.Add("mentor", c.QueryParam("mentor"))
	if isHireable, err := strconv.ParseBool(c.QueryParam("isHireable")); err == nil {
		query.Add("isHireable", isHireable)
	}

	items, err := r.repository.Get(c.Request().Context(), query, pg)
	if err != nil {
		panic(err)
	}
	if len(items) == 0 {
		panic(person.NotFoundError)
	}

	return c.JSON(http.StatusOK, items)
}

func (r Controller) GetBySlug(c echo.Context) error {

	slug := c.Param("slug")
	if len(slug)==0{
		return c.JSON(http.StatusInternalServerError,fmt.Sprintf("slug can not be null or empty"))
	}
	query := mongohelper.QueryOf()
	query.Add("slug",slug)
	person,err := r.repository.Get(c.Request().Context(),query,nil)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK,person)
}
