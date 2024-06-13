package routes

import (
	"github.com/harish876/hypefx/internal/generators/test/app"
	"github.com/harish876/hypefx/internal/generators/test/app/form"
	"github.com/harish876/hypefx/internal/generators/test/app/form/input"
	"github.com/harish876/hypefx/internal/generators/test/app/grid"
	"github.com/harish876/hypefx/internal/generators/test/app/grid/edit"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	apiGroup0 := e.Group("/form")
	apiGroup0.POST("", form.POST)
	apiGroup0.GET("", form.GET)

	apiGroup1 := e.Group("/form/input")
	apiGroup1.GET("", input.RenderFormInput)
	apiGroup1.POST("", input.PostFormInput)

	apiGroup2 := e.Group("/grid/edit/:id")
	apiGroup2.GET("", edit.RenderEditGridHandler)

	apiGroup3 := e.Group("/grid")
	apiGroup3.GET("", grid.GridHandler)
	apiGroup3.POST("", grid.GridFilterHandler)

	apiGroup4 := e.Group("/grid/:id")
	apiGroup4.GET("", grid.GridRowHandler)
	apiGroup4.PUT("", grid.UpdateGridRowHandler)
	apiGroup4.DELETE("", grid.DeleteGridRowHandler)

	apiGroup5 := e.Group("")
	apiGroup5.GET("", app.HomeHandler)
}
