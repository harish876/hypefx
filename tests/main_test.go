package tests

import (
	"os"
	"testing"

	"github.com/harish876/hypefx/internal/generators/template"
)

func TestTemplateGeneration(t *testing.T) {
	templateParams := template.TemplateParams{
		BasePath:            "/home/harish/personal/hypefx/tests/generator/app",
		BaseImportPath:      "github.com/harish876/hypefx/tests/generator",
		TemplateName:        "routes",
		RouteDirPackageName: "routes",
		DestinationDir:      "/home/harish/personal/hypefx/tests/generator/routes/routes.go",
	}
	if err := template.Generator(templateParams); err != nil {
		t.Fatalf("unable to generate templates: %v", err)
	}

	fileContents, err := os.ReadFile(templateParams.DestinationDir)
	if err != nil {
		t.Fatalf("unable to open the destination dir")
	}
	_ = fileContents
}
