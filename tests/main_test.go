package tests

import (
	"testing"

	"github.com/harish876/hypefx/internal/generators/template"
)

func TestTemplateGeneration(t *testing.T) {
	templateParams := template.TemplateParams{
		BasePath:            "/home/harish/personal/hypefx/tests/generator/app",
		BaseImportPath:      "github.com/harish876/hypefx/tests/generator",
		TemplateName:        "routes",
		RouteDirPackageName: "routes",
		DestinationDir:      "/home/harish/personal/hypefx/tests/generator/routes",
	}
	if err := template.Generator(templateParams); err != nil {
		t.Fatalf("unable to generate templates: %v", err)
	}
}
