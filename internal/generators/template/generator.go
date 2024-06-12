package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/harish876/hypefx/internal/generators/annotations"
)

type Route struct {
	PackageName string
	RouteGroup  string
	Handlers    []Handler
}

type Handler struct {
	Name       string
	HttpMethod string
}

// path to the app directory
func Generator(templateParams TemplateParams) error {
	var routes []Route
	err := filepath.Walk(templateParams.BasePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			relativePath := strings.TrimPrefix(path, templateParams.BasePath)
			pathToFile := strings.TrimSuffix(relativePath, "/"+info.Name())

			if !strings.HasSuffix(info.Name(), "index.go") {
				//do something specific to different file patterns
				fmt.Println("file types other than index.go not implemented yet.continuing execution for other...")
				return nil

			}

			fileContents, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			handlerDetails, packageName, err := annotations.GetHandlerDetailsFromAnnotations(fileContents)
			if err != nil {
				return err
			}
			if len(handlerDetails) == 0 {
				return fmt.Errorf("no annotations found. generator is unable to map the handler names to the routes. please add annotations to your routes for autogeneration or unset the routing flag: %s", path)
			}

			fmt.Println(
				"Handler Map: ", handlerDetails,
				"\nPackage Name: ", packageName,
				"\nFull Path: ", path,
				"\n Route Group: ", pathToFile,
				"\n",
			)

			var handlers []Handler
			for _, val := range handlerDetails {

				handlers = append(handlers, Handler{
					Name:       val.HandlerName,
					HttpMethod: annotations.FromEnum(val.Method),
				})
			}

			routes = append(routes, Route{
				PackageName: packageName,
				RouteGroup:  pathToFile,
				Handlers:    handlers,
			})
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path: %v\n", err)
		return err
	}

	err = generateHandlerFile(routes, templateParams)
	if err != nil {
		fmt.Printf("Error generating handler file: %v\n", err)
		return err
	}
	return nil
}

func generateHandlerFile(routes []Route, templateParams TemplateParams) error {
	var tpl = fmt.Sprintf(`
package %s

import (
{{- range $i, $a := .Imports }}
	"{{$a}}"
{{- end }}
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
{{- range $idx, $route := .Routes }}

	apiGroup{{$idx}} := e.Group("{{$route.RouteGroup}}")

	{{- range .Handlers}}
		apiGroup{{$idx}}.{{.HttpMethod}}("",{{$route.PackageName}}.{{.Name}})
	{{- end }}

{{- end }}
}
`, templateParams.RouteDirPackageName)

	importMap := make(map[string]struct{})
	for _, route := range routes {
		importMap[route.RouteGroup] = struct{}{}
	}

	imports := make([]string, 0, len(importMap))
	for pkg := range importMap {
		imports = append(imports, fmt.Sprintf("%s%s", templateParams.BaseImportPath, pkg))
	}

	data := struct {
		Imports []string
		Routes  []Route
	}{
		Imports: imports,
		Routes:  routes,
	}

	file, err := os.Create(templateParams.DestinationDir) //revisit this madness
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
