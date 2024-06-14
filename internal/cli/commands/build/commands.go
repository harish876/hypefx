package build

import (
	"fmt"
	"os/exec"

	"github.com/harish876/hypefx/internal/cli/commands"
	"github.com/harish876/hypefx/internal/generators/template"
	"github.com/spf13/cobra"
)

func build(cmd *cobra.Command, args []string) {
	//this needs to set dynamically. autodetect this
	basePath, _ := commands.GetConfig("basePath")
	module, _ := commands.GetConfig("module")
	routesDir, _ := commands.GetConfig("routesDir")

	templateParams := template.TemplateParams{
		BasePath:            basePath.(string),      //config.get("basePath")
		BaseImportPath:      module.(string),        //config.get("module")
		TemplateName:        commands.TEMPLATE_NAME, // constant.routes
		RouteDirPackageName: commands.TEMPLATE_NAME, // constant.routes
		DestinationDir:      routesDir.(string),     //config.get("routeDir")
	}

	if err := template.Generator(templateParams); err != nil {
		DisplayError(fmt.Errorf("error running template generator : %v", err))
	}
	fmtCmd := exec.Command("gofmt", "-w", templateParams.DestinationDir)
	_, err := fmtCmd.CombinedOutput()
	if err != nil {
		DisplayError(fmt.Errorf("error formatting the routes file : %v", err))
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
