package app

import (
	"net/http"

	"github.com/harish876/hypefx/examples/views/home"
	"github.com/labstack/echo/v4"
)

// @get
func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, home.Home())
}
