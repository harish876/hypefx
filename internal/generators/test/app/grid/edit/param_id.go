package edit

import (
	"net/http"

	"github.com/harish876/hypefx/components/grid"
	"github.com/harish876/hypefx/examples/db"
	"github.com/harish876/hypefx/examples/services"
	"github.com/harish876/hypefx/internal/generators/template"
	model "github.com/harish876/hypefx/internal/generators/test/app/grid"
	"github.com/labstack/echo/v4"
)

// @get
func RenderEditGridHandler(c echo.Context) error {
	id := c.Param("id")
	row := services.FilterById(db.GridData, id)
	return template.Render(c, http.StatusOK, grid.RenderEditableRow(&model.GridModel, model.Columns, row))
}
