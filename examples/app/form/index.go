package form

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/harish876/hypefx/components/props"
	app "github.com/harish876/hypefx/examples/app"
	"github.com/harish876/hypefx/examples/views/form"
	"github.com/labstack/echo/v4"
)

// @post
func FormSubmissionHandler(c echo.Context) error {
	formValues := props.FormValues{
		FirstName:            c.FormValue("first_name"),
		LastName:             c.FormValue("last_name"),
		Email:                c.FormValue("email"),
		Password:             c.FormValue("password"),
		PasswordConfirmation: c.FormValue("password_confirmation"),
		MarketingAccept:      c.FormValue("marketing_accept"),
	}

	var formErrors props.FormErrors
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(emailRegex)

	fmt.Println(formValues)
	if formValues.Password != formValues.PasswordConfirmation {
		formErrors.Password = "Password and Password Confirmation must match"
	}
	if ok := regex.MatchString(formValues.Email); !ok {
		formErrors.Email = "Invalid email format"
	}

	if formErrors.Email != "" || formErrors.Password != "" {
		fmt.Println(formErrors)
		return app.Render(c, 422, form.Form(formValues, formErrors, false, ""))
	}

	return app.Render(c, http.StatusOK, form.Form(props.FormValues{}, props.FormErrors{}, true, "Success"))
}

// @get
func FormHandler(c echo.Context) error {
	return app.Render(c, http.StatusOK, form.Form(props.FormValues{}, props.FormErrors{}, false, ""))
}
