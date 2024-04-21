package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/harish876/hypefx/components/notfound"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func NotFoundErrorHandler(err error, c echo.Context) {
	// Check if the error is a 404 Not Found error
	if h, ok := err.(*echo.HTTPError); ok {
		if h.Code == http.StatusNotFound {
			Render(c, http.StatusNotFound, notfound.NotFound())
		}
	}
}
