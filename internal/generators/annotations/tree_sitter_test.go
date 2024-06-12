package annotations

import (
	"fmt"
	"testing"
)

func TestGetHandlerDetailsFromAnnotations(t *testing.T) {
	sourceCode := []byte(`
package form

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/harish876/hypefx/components/props"
	"github.com/harish876/hypefx/examples/views/form"
	renderer "github.com/harish876/hypefx/internal/generators/template"
	"github.com/labstack/echo/v4"
)

// @post
func PostSomethingSlick(c echo.Context) error {
	formValues := props.FormValues{
		FirstName:            c.FormValue("first_name"),
		LastName:             c.FormValue("last_name"),
		Email:                c.FormValue("email"),
		Password:             c.FormValue("password"),
		PasswordConfirmation: c.FormValue("password_confirmation"),
		MarketingAccept:      c.FormValue("marketing_accept"),
	}

	var formErrors props.FormErrors
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
		return renderer.Render(c, 422, form.Form(formValues, formErrors, false, ""))
	}

	return renderer.Render(c, http.StatusOK, form.Form(props.FormValues{}, props.FormErrors{}, true, "Success"))
}

// @get
func GetSomethingSlick(c echo.Context) error {
	return renderer.Render(c, http.StatusOK, form.Form(props.FormValues{}, props.FormErrors{}, false, ""))
}
`)
	handlerDetails, packageName, err := GetHandlerDetailsFromAnnotations(sourceCode)
	if err != nil {
		t.Fail()
	}
	fmt.Println(handlerDetails, packageName)
}
func TestGetHandlerDetailsFromAnnotationsNoAnnotations(t *testing.T) {
	sourceCode := []byte(`
package form

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/harish876/hypefx/components/props"
	"github.com/harish876/hypefx/examples/views/form"
	renderer "github.com/harish876/hypefx/internal/generators/template"
	"github.com/labstack/echo/v4"
)

func PostSomethingSlick(c echo.Context) error {
	formValues := props.FormValues{
		FirstName:            c.FormValue("first_name"),
		LastName:             c.FormValue("last_name"),
		Email:                c.FormValue("email"),
		Password:             c.FormValue("password"),
		PasswordConfirmation: c.FormValue("password_confirmation"),
		MarketingAccept:      c.FormValue("marketing_accept"),
	}

	var formErrors props.FormErrors
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
		return renderer.Render(c, 422, form.Form(formValues, formErrors, false, ""))
	}

	return renderer.Render(c, http.StatusOK, form.Form(props.FormValues{}, props.FormErrors{}, true, "Success"))
}

//unrelated comment

func GetSomethingSlick(c echo.Context) error {
	return renderer.Render(c, http.StatusOK, form.Form(props.FormValues{}, props.FormErrors{}, false, ""))
}
`)
	handlerDetails, packageName, err := GetHandlerDetailsFromAnnotations(sourceCode)
	if err != nil {
		t.Fail()
	}
	fmt.Println(handlerDetails, packageName)
}
