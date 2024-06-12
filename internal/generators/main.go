package main

import (
	"log"

	"github.com/harish876/hypefx/internal/generators/template"
)

func main() {
	templateParams := template.TemplateParams{
		BasePath:            "/home/harish/personal/hypefx/examples/app",
		BaseImportPath:      "github.com/harish876/hypefx/examples/app",
		TemplateName:        "routes",
		RouteDirPackageName: "routes",
		DestinationDir:      "/home/harish/personal/hypefx/examples/routes/routes.go",
	}
	if err := template.Generator(templateParams); err != nil {
		log.Fatalf("error running template generator: %s", err)
	}
}
