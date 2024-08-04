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
	Name      string
	Directive string
	Page      string
	Url       string
	FileName  string
	IsStatic  bool
	IsDebug   bool
}

func Generator(templateParams TemplateParams) error {
	var routes []Route
	err := filepath.Walk(
		templateParams.BasePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				relativePath := strings.TrimPrefix(path, templateParams.BasePath)
				pathToFile := strings.TrimSuffix(relativePath, "/"+info.Name())

				if strings.HasSuffix(info.Name(), "index.go") ||
					strings.Contains(info.Name(), "param") {

					fileContents, err := os.ReadFile(path)
					if err != nil {
						return err
					}

					handlerDetails, packageName, err := annotations.GetHandlerDetailsFromAnnotations(
						fileContents,
					)
					if err != nil {
						return err
					}
					if len(handlerDetails) == 0 {
						//revisit this block. bug, if one method is missed, there is not comparator
						return fmt.Errorf(
							"no annotations found. generator is unable to map the handler names to the routes. please add annotations to your routes for autogeneration or unset the routing flag: %s",
							path,
						)
					}

					routeName := pathToFile

					if strings.Contains(info.Name(), "param") {
						fileName := strings.TrimSuffix(info.Name(), ".go")
						paramFileFormat := strings.Split(fileName, "_")
						if len(paramFileFormat) < 2 {
							return fmt.Errorf(
								"invalid param file format. it should be param_[param_name].go",
							)
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

						var handler Handler
						handler.Name = val.HandlerName
						handler.Directive = annotations.FromEnum(val.Direcive)
						handler.Page = strings.ReplaceAll(val.DirectiveParams["page"], `"`, "")
						handler.Url = val.DirectiveParams["url"]
						handler.IsStatic = val.Direcive == annotations.STATIC
						if _, ok := val.DirectiveParams["debug"]; ok {
							handler.IsDebug = true

						}
						if handler.IsStatic && !handler.IsDebug {
							slog.Info("Generator", "Params", val.DirectiveParams)
							handlerStaticFileGeneration()
						}
						handlers = append(handlers, handler)
					}

					routes = append(routes, Route{
						PackageName: packageName,
						RouteGroup:  pathToFile,
						RouteName:   routeName,
						Handlers:    handlers,
					})
				} else {
					slog.Error("Generator", "file types other than index.go not implemented yet.file name", path)
				}
			} else {
				//do something specific to different file patterns
				return nil
			}
			return nil
		},
	)

	if err != nil {
		slog.Error("Generator", "Error walking the path:", err.Error())
		return err
	}

	err = generateHandlerFile(routes, templateParams)
	if err != nil {
		slog.Error("Generator", "Error generating handler file:", err.Error())
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
		{{- if .IsStatic}}
			apiGroup{{$idx}}.File("","public/{{.Page}}")
		{{- else }}
			apiGroup{{$idx}}.{{.Directive}}("",{{$route.PackageName}}.{{.Name}})
		{{- end }}
	{{- end }}

{{- end }}
}
`, templateParams.RouteDirPackageName)

	importMap := make(map[string]struct{})
	for _, route := range routes {
		slog.Debug("generateHanlderFile", "imports", route.RouteGroup)
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

	file, err := os.Create(filepath.Join(templateParams.DestinationDir, "routes.go"))
	if err != nil {
		return err
	}
	defer file.Close()

	tmpl, err := template.New(templateParams.TemplateName).Parse(tpl)
	if err != nil {
		return err
	}

	return tmpl.Execute(file, data)
}

func handlerStaticFileGeneration() error {
	return nil
}
