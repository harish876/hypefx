package handlers

import (
	"net/http"

	"github.com/harish876/hypefx/scaffolding/views/welcome"
	"github.com/labstack/echo/v4"
)

func WelcomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, welcome.Welcome())
}
