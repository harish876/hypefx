package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/", "assets")
	app.Logger.Fatal(app.Start(":42070"))
}
