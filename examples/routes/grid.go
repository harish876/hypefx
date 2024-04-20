package routes

import (
	"github.com/harish876/hypefx/examples/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterGridRoutes(e *echo.Echo) {
	apiGroup := e.Group("/grid")

	apiGroup.GET("", handlers.GridHandler)
	apiGroup.POST("", handlers.GridFilterHandler)
	apiGroup.GET("/:id", handlers.GridRowHandler)
	apiGroup.PUT("/:id", handlers.UpdateGridRowHandler)
	apiGroup.DELETE("/:id", handlers.DeleteGridRowHandler)
	apiGroup.GET("/edit/:id", handlers.RenderEditGridHandler)
}
