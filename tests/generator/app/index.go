package app

import (
	"context"
	"net/http"
	"os"

	"github.com/harish876/hypefx/internal/generators/template"
	"github.com/harish876/hypefx/tests/generator/views/home"
	"github.com/labstack/echo/v4"
)

// @static(url="/",page="home.html")
func HomeHandler(c echo.Context) error {
	f, _ := os.Open("./public/home.html")
	home.Home().Render(context.Background(), f)
	return template.Render(c, http.StatusOK, home.Home())
}
