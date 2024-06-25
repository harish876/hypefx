
package routes

import (
	"github.com/harish876/hypefx/tests/generator/app/form"
	"github.com/harish876/hypefx/tests/generator/app/form/input"
	"github.com/harish876/hypefx/tests/generator/app"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	apiGroup0 := e.Group("/form")
		apiGroup0.POST("",form.POST)
		apiGroup0.GET("",form.GET)

	apiGroup1 := e.Group("/form/input")
		apiGroup1.GET("",input.RenderFormInput)
		apiGroup1.POST("",input.PostFormInput)

	apiGroup2 := e.Group("")
		apiGroup2.GET("",app.HomeHandler)
}
