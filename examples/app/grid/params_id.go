package grid

import (
	"net/http"

	"github.com/harish876/hypefx/components/grid"
	"github.com/harish876/hypefx/examples/db"
	"github.com/harish876/hypefx/examples/services"
	"github.com/harish876/hypefx/internal/generators/template"
	"github.com/labstack/echo/v4"
)

// @get
func GridRowHandler(c echo.Context) error {
	id := c.Param("id")
	row := services.FilterById(db.GridData, id)
	return template.Render(c, http.StatusOK, grid.RenderRow(&GridModel, Columns, row))
}

// @put
func UpdateGridRowHandler(c echo.Context) error {
	id := c.Param("id")
	var newData = db.GridDataRow{
		Id:       id,
		Status:   c.FormValue("status"),
		Position: c.FormValue("position"),
	}
	updatedRow := services.UpdateById(db.GridData, id, newData)
	return template.Render(c, http.StatusOK, grid.RenderRow(&GridModel, Columns, updatedRow))
}

// @delete
func DeleteGridRowHandler(c echo.Context) error {
	id := c.Param("id")
	services.DeleteById(id)
	return c.NoContent(http.StatusOK)
}
