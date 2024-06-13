package main

import (
	"flag"
	"log"

	"github.com/harish876/hypefx/internal/generators/template"
)

func main() {
	var MODE = flag.String("mode", "test", "mode for running the cli temporarily")
	flag.Parse()
	var templateParams template.TemplateParams
	if *MODE == "main" {
		templateParams = template.TemplateParams{
			BasePath:            "/home/harish/personal/hypefx/examples/app",
			BaseImportPath:      "github.com/harish876/hypefx/examples/app",
			TemplateName:        "routes",
			RouteDirPackageName: "routes",
			DestinationDir:      "/home/harish/personal/hypefx/examples/routes/routes.go",
		}
	} else {
		templateParams = template.TemplateParams{
			BasePath:            "test/app",
			BaseImportPath:      "github.com/harish876/hypefx/internal/generators/test/app",
			TemplateName:        "routes",
			RouteDirPackageName: "routes",
			DestinationDir:      "test/routes/routes.go",
		}
	}

	if err := template.Generator(templateParams); err != nil {
		log.Fatalf("error running template generator: %s", err)
	}
}
