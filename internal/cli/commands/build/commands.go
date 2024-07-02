package build

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/harish876/hypefx/internal/cli/commands"
	"github.com/harish876/hypefx/internal/generators/static"
	"github.com/harish876/hypefx/internal/generators/template"
	"github.com/spf13/cobra"
)

var (
	ROUTES_FILE = "routes.go"
)

func build(cmd *cobra.Command, args []string) {
	slog.Debug("build", "args", args)
	config, err := commands.GetConfig()
	if err != nil {
		DisplayError(fmt.Errorf("unable to read hypeconfig.json: %s.\nuse hypefx generate [module_name] to get things started", err))
		return
	}
	switch args[0] {
	case "routes":
		dir, err := buildRoutes(config)
		if err != nil {
			DisplayError(err)
			return
		}
		DisplaySuccessMessage(fmt.Sprintf("Routes generated at %s", dir))
	case "staticPages":
		if err := buildStaticPages(config); err != nil {
			DisplayError(err)
		}
	default:
		dir, err := buildRoutes(config)
		if err != nil {
			DisplayError(err)
			return
		}
		DisplaySuccessMessage(fmt.Sprintf("Routes generated at %s", dir))
		if err := buildStaticPages(config); err != nil {
			DisplayError(err)
		}
	}

}

func buildRoutes(config commands.HypeConfig) (string, error) {
	if config.AppDir == "" {
		return "", fmt.Errorf("appDir is not set. use hypefx set --appDir [value]")
	}
	if config.Module == "" {
		return "", fmt.Errorf("module is not set. use hypefx set --module [value]")
	}
	if config.RoutesDir == "" {
		return "", fmt.Errorf("routesDir is not set. use hypefx set --routesDir [value]")
	}
	if !config.Routing {
		return "", fmt.Errorf("enable routing to automagically build routes")
	}

	routesDir := config.RoutesDir
	os.MkdirAll(routesDir, os.ModePerm)
	if file, err := os.Create(filepath.Join(routesDir, ROUTES_FILE)); err != nil {
		slog.Error("build", "creating routes path", err.Error())
		defer file.Close()
	}
	slog.Debug("build", "creating routes path dir", "created routes path dir successfully")

	templateParams := template.TemplateParams{
		BasePath:            config.AppDir,          //config.get("appDir")
		BaseImportPath:      config.Module,          //config.get("module")
		TemplateName:        commands.TEMPLATE_NAME, // constant.routes
		RouteDirPackageName: commands.TEMPLATE_NAME, // constant.routes
		DestinationDir:      routesDir,              //config.get("routeDir")
	}

	if err := template.Generator(templateParams); err != nil {
		return "", fmt.Errorf("error running template generator : %v", err)
	}
	fmtCmd := exec.Command("gofmt", "-w", templateParams.DestinationDir)
	_, err := fmtCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error formatting the routes file : %v", err)
	}
	return templateParams.DestinationDir, nil
}

func buildStaticPages(config commands.HypeConfig) error {
	os.MkdirAll("public", os.ModePerm)
	staticFiles := config.StaticFiles
	if len(staticFiles) == 0 {
		return nil
	}
	if err := static.Generator(staticFiles, config.Module+"/views", "main"); err != nil {
		return err
	}
	fmtCmd := exec.Command("go", "run", "public/static_runner.go")
	_, err := fmtCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error creating static pages : %v", err)
	}
	return nil
}

var BuildCmd = &cobra.Command{
	Use:     "build",
	Short:   "Builds the routes using the app directory",
	Long:    `Builds the routes dynamically using the app directory. This provides a neat compile time FBR`,
	Example: "hype build",
	Args:    cobra.ExactArgs(1),
	Run:     build,
}
