package routes

import (
	"github.com/harish876/hypefx/examples/app"
	"github.com/harish876/hypefx/examples/app/form"
	"github.com/harish876/hypefx/examples/app/grid"
	"github.com/harish876/hypefx/examples/app/grid/edit"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	apiGroup0 := e.Group("/form")
	apiGroup0.POST("", form.FormSubmissionHandler)
	apiGroup0.GET("", form.FormHandler)

	apiGroup1 := e.Group("/grid/edit/:id")
	apiGroup1.GET("", edit.RenderEditGridHandler)

	apiGroup2 := e.Group("/grid")
	apiGroup2.POST("", grid.GridFilterHandler)

	apiGroup3 := e.Group("/grid/:id")
	apiGroup3.GET("", grid.GridRowHandler)
	apiGroup3.PUT("", grid.UpdateGridRowHandler)
	apiGroup3.DELETE("", grid.DeleteGridRowHandler)

	apiGroup4 := e.Group("")
	apiGroup4.GET("", app.HomeHandler)
}
