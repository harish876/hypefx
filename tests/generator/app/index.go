package app

import (
	"net/http"

	"github.com/harish876/hypefx/internal/generators/template"
	"github.com/harish876/hypefx/tests/generator/views/home"
	"github.com/labstack/echo/v4"
)

// @get
func HomeHandler(c echo.Context) error {
	return template.Render(c, http.StatusOK, home.Home())
}
