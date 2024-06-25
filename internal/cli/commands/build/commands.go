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
	configs, err := commands.GetAllConfig()
	if err != nil {
		DisplayError(fmt.Errorf("unable to read hypeconfig.json: %s.\nuse hypefx generate [module_name] to get things started", err))
		return
	}
	appDir, err := commands.FromConfig(configs, "appDir")
	if err != nil || appDir == nil {
		DisplayError(err)
		return
	}
	module, err := commands.FromConfig(configs, "module")
	if err != nil || module == nil {
		DisplayError(err)
		return
	}
	routesDir, err := commands.FromConfig(configs, "routesDir")
	if err != nil || routesDir == nil {
		DisplayError(err)
		return
	}

	_, err = commands.FromConfig(configs, "routing")
	if err != nil {
		DisplayError(errors.Join(err, fmt.Errorf("enable routing to automagically build routes")))
		return
	}

	os.MkdirAll(routesDir.(string), os.ModePerm)
	if file, err := os.Create(filepath.Join(routesDir.(string), ROUTES_FILE)); err != nil {
		slog.Error("build", "creating routes path", err.Error())
		defer file.Close()
	}
	slog.Debug("build", "creating routes path dir", "created routes path dir successfully")

	templateParams := template.TemplateParams{
		BasePath:            appDir.(string),        //config.get("appDir")
		BaseImportPath:      module.(string),        //config.get("module")
		TemplateName:        commands.TEMPLATE_NAME, // constant.routes
		RouteDirPackageName: commands.TEMPLATE_NAME, // constant.routes
		DestinationDir:      routesDir.(string),     //config.get("routeDir")
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
