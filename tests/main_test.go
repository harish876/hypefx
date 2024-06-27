package tests

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/harish876/hypefx/internal/generators/template"
)

func TestTemplateGeneration(t *testing.T) {
	pwd, _ := os.Getwd()
	slog.SetLogLoggerLevel(slog.LevelDebug)
	fmt.Println("Present Working directory", pwd)
	templateParams := template.TemplateParams{
		BasePath:            filepath.Join(pwd, "generator", "app"),
		BaseImportPath:      "github.com/harish876/hypefx/tests/generator",
		TemplateName:        "routes",
		RouteDirPackageName: "routes",
		DestinationDir:      filepath.Join(pwd, "generator", "routes"),
	}
	if err := template.Generator(templateParams); err != nil {
		t.Fatalf("unable to generate templates: %v", err)
	}
}
