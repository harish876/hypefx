package handlers

import (
	"net/http"

	"github.com/harish876/hypefx/internal/cli/commands/generate/scaffolding/views/welcome"
	"github.com/labstack/echo/v4"
)

// @get
func WelcomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, welcome.Welcome())
}
