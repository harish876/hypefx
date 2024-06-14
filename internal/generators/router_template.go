// router_template.go

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Route struct {
	PackageName string
	RouteGroup  string
	RouteName   string
	Method      string
}

func Runner(basePath string) {
	var routes []Route
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), "index.go") {
			relativePath := strings.TrimPrefix(path, basePath)
			relativePath = strings.TrimPrefix(relativePath, string(os.PathSeparator))
			dir, _ := filepath.Split(relativePath)
			dir = strings.TrimSuffix(dir, string(os.PathSeparator))

			paths := strings.Split(dir, "/")
			packageName := paths[len(paths)-1]

			routeName := strings.ToLower(strings.TrimSuffix(info.Name(), ".go"))
			if routeName == "index" {
				routeName = "/"
			}
			routes = append(routes, Route{
				PackageName: packageName,
				RouteGroup:  dir,
				RouteName:   routeName,
			})
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path: %v\n", err)
		return
	}

	err = generateHandlerFile(routes)
	if err != nil {
		fmt.Printf("Error generating handler file: %v\n", err)
		return
	}
}

func generateHandlerFile(routes []Route) error {
	const tpl = `
package route

import (
{{- range $i, $a := .Imports }}
	"{{$a}}"
{{- end }}
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
{{- range $i, $a := .Routes }}

	apiGroup{{$i}} := e.Group("/{{.RouteGroup}}")

	apiGroup{{$i}}.GET("", {{.PackageName}}.GET)
	apiGroup{{$i}}.POST("", {{.PackageName}}.POST)

{{- end }}
}
`

	importMap := make(map[string]struct{})
	for _, route := range routes {
		importMap[route.RouteGroup] = struct{}{}
	}

	imports := make([]string, 0, len(importMap))
	for pkg := range importMap {
		imports = append(imports, fmt.Sprintf("github.com/harish876/hypefx/internal/generators/test/app/%s", pkg))
	}

	data := struct {
		Imports []string
		Routes  []Route
	}{
		Imports: imports,
		Routes:  routes,
	}

	file, err := os.Create("test/routes/form.go")
	if err != nil {
		return err
	}
	defer file.Close()

	tmpl, err := template.New("routes").Parse(tpl)
	if err != nil {
		return err
	}

	return tmpl.Execute(file, data)
}
