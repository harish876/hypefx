package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/harish876/hypefx/internal/generators/template"
	"github.com/stretchr/testify/assert"
)

func TestTemplateGeneration(t *testing.T) {
	pwd, err := os.Getwd()
	assert.NoError(t, err, "Error getting pwd")
	templateParams := template.TemplateParams{
		BasePath:            fmt.Sprintf("%s/generator/app", pwd),
		BaseImportPath:      "github.com/harish876/hypefx/tests/generator",
		TemplateName:        "routes",
		RouteDirPackageName: "routes",
		DestinationDir:      fmt.Sprintf("%s/generator/routes", pwd),
	}
	err = template.Generator(templateParams)
	assert.NoError(t, err, "unable to generate templates")
}
