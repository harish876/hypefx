package input

import (
	"net/http"

	"github.com/harish876/hypefx/components/props"
	"github.com/harish876/hypefx/tests/generator/views/form"
	renderer "github.com/harish876/hypefx/internal/generators/template"
	"github.com/labstack/echo/v4"
)

// @get
func RenderFormInput(c echo.Context) error {
	return renderer.Render(c, http.StatusOK, form.Form(props.FormValues{}, props.FormErrors{}, false, ""))
}

// @post
func PostFormInput(c echo.Context) error {
	return renderer.Render(c, http.StatusOK, form.Form(props.FormValues{}, props.FormErrors{}, false, ""))
}
