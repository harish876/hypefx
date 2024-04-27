package main

import (
	"github.com/harish876/hypefx/scaffolding/handlers"
	"github.com/harish876/hypefx/scaffolding/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/", "assets")

	app.HTTPErrorHandler = handlers.NotFoundErrorHandler
	routes.RegisterWelcomeRoutes(app)
	app.Logger.Fatal(app.Start(":42070"))
}
