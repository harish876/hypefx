package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/", "assets")

	/* Starting Point for registering all routes. Uncomment this to see the generated routes work*/
	//routes.RegisterRoutes(app)

	app.Logger.Fatal(app.Start(":42070"))
}
