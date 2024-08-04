package app

import (
	"net/http"

	"github.com/harish876/hypefx/internal/generators/template"
	"github.com/harish876/hypefx/tests/generator/views/home"
	"github.com/labstack/echo/v4"
)

// @get
// @static(url="/about",page="home.html",debug="true")
func HomeHandler(c echo.Context) error {
	return template.Render(c, http.StatusOK, home.Home())
}
