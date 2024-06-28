package build

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/harish876/hypefx/internal/cli/commands"
	"github.com/harish876/hypefx/internal/generators/template"
	"github.com/spf13/cobra"
)

var (
	ROUTES_FILE = "routes.go"
)

func build(cmd *cobra.Command, args []string) {
	//single disk read
	config, err := commands.GetConfig()
	if err != nil {
		DisplayError(fmt.Errorf("unable to read hypeconfig.json: %s.\nuse hypefx generate [module_name] to get things started", err))
		return
	}
	if config.AppDir == "" {
		DisplayError(err)
		return
	}
	if config.Module == "" {
		DisplayError(err)
		return
	}
	if config.RoutesDir == "" {
		DisplayError(err)
		return
	}

	if !config.Routing {
		DisplayError(errors.Join(err, fmt.Errorf("enable routing to automagically build routes")))
		return
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
		DisplayError(fmt.Errorf("error running template generator : %v", err))
		return
	}
	fmtCmd := exec.Command("gofmt", "-w", templateParams.DestinationDir)
	_, err = fmtCmd.CombinedOutput()
	if err != nil {
		DisplayError(fmt.Errorf("error formatting the routes file : %v", err))
		return
	}
	DisplaySuccessMessage(fmt.Sprintf("Routes generated at %s", templateParams.DestinationDir))
}

var BuildCmd = &cobra.Command{
	Use:     "build",
	Short:   "Builds the routes using the app directory",
	Long:    `Builds the routes dynamically using the app directory. This provides a neat compile time FBR`,
	Example: "hype build",
	Run:     build,
}
