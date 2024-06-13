package main

import (
	"github.com/harish876/hypefx/examples/handlers"
	"github.com/harish876/hypefx/examples/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/", "assets")

	app.HTTPErrorHandler = handlers.NotFoundErrorHandler

	routes.RegisterRoutes(app)
	// routes.RegisterHomeRoutes(app)
	// routes.RegisterGridRoutes(app)
	// routes.RegisterFormRoutes(app)

	app.Logger.Fatal(app.Start(":42069"))
}
