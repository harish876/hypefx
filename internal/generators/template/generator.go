package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"log/slog"

	"github.com/harish876/hypefx/internal/generators/annotations"
)

type Route struct {
	PackageName string
	RouteGroup  string
	RouteName   string
	Handlers    []Handler
}

type Handler struct {
	Name       string
	HttpMethod string
}

func Generator(templateParams TemplateParams) error {
	var routes []Route
	err := filepath.Walk(templateParams.BasePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			relativePath := strings.TrimPrefix(path, templateParams.BasePath)
			pathToFile := strings.TrimSuffix(relativePath, "/"+info.Name())

			if strings.HasSuffix(info.Name(), "index.go") || strings.Contains(info.Name(), "param") {

				fileContents, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				handlerDetails, packageName, err := annotations.GetHandlerDetailsFromAnnotations(fileContents)
				if err != nil {
					return err
				}
				if len(handlerDetails) == 0 {
					//revisit this block. bug, if one method is missed, there is not comparator
					return fmt.Errorf("no annotations found. generator is unable to map the handler names to the routes. please add annotations to your routes for autogeneration or unset the routing flag: %s", path)
				}

				routeName := pathToFile

				if strings.Contains(info.Name(), "param") {
					fileName := strings.TrimSuffix(info.Name(), ".go")
					paramFileFormat := strings.Split(fileName, "_")
					if len(paramFileFormat) < 2 {
						return fmt.Errorf("invalid param file format. it should be param_[param_name].go")
					}
					routeName = pathToFile + "/:" + paramFileFormat[1]
				}

				slog.Debug(
					"Generator",
					"Handler Map: ", handlerDetails,
					"Package Name: ", packageName,
					"Full Path: ", path,
					"Route Name: ", routeName,
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
					RouteName:   routeName,
					Handlers:    handlers,
				})
			}
		} else {
			//do something specific to different file patterns
			//fmt.Println("file types other than index.go not implemented yet.continuing execution for other...")
			return nil
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

	apiGroup{{$idx}} := e.Group("{{$route.RouteName}}")

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
		imports = append(imports, fmt.Sprintf("%s/app%s", templateParams.BaseImportPath, pkg))
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
