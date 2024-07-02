package static

import (
	"fmt"
	"html/template"
	"log/slog"
	"os"

	"github.com/harish876/hypefx/internal/cli/commands"
)

func Generator(staticFiles []commands.StaticFile, basePath string, packagePath string) error {
	tmpl := `package {{.PackagePath}}

import (
{{- range .Imports }}
    "{{ . }}"
{{- end }}
)

//RUN THIS FILE TO REGENERATE THE STATIC PAGES

func main() {
{{- range $idx,$file := .StaticFiles }}
    f_{{$idx}}, err := os.Create("{{ $file.FileName }}")
    if err != nil {
        log.Fatalf("failed to create output file '{{ $file.FileName }}': %v", err)
    }
    defer f_{{$idx}}.Close()

    err = {{ $file.Package }}.{{ $file.Component }}().Render(context.Background(), f_{{$idx}})
    if err != nil {
        log.Fatalf("failed to write output file '{{ $file.FileName }}': %v", err)
    }
    log.Printf("{{ $file.FileName }} created successfully.")

{{- end }}
}
`

	importsMap := map[string]string{
		"context": "",
		"log":     "",
		"os":      "",
	}

	for _, files := range staticFiles {
		package_name := fmt.Sprintf("%s/%s", basePath, files.Package)
		importsMap[package_name] = ""
	}

	imports := make([]string, 0, len(importsMap))
	for key, _ := range importsMap {
		imports = append(imports, key)
	}

	data := struct {
		StaticFiles []commands.StaticFile
		Imports     []string
		PackagePath string
	}{
		StaticFiles: staticFiles,
		Imports:     imports,
		PackagePath: packagePath,
	}

	t := template.Must(template.New("goTemplate").Parse(tmpl))

	outputFile, err := os.Create("public/static_runner.go") //revisit
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer outputFile.Close()

	err = t.Execute(outputFile, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	slog.Info("static_runner.go created successfully.")
	return nil
}
