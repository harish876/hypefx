package routes

import (
	"github.com/harish876/hypefx/scaffolding/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterWelcomeRoutes(e *echo.Echo) {
	apiGroup := e.Group("/")

	apiGroup.GET("", handlers.WelcomeHandler)
}
