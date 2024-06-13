package build

import (
	"fmt"
	"os/exec"

	"github.com/harish876/hypefx/internal/generators/template"
	"github.com/spf13/cobra"
)

func build(cmd *cobra.Command, args []string) {
	//this needs to set dynamically. autodetect this
	templateParams := template.TemplateParams{
		BasePath:            "/home/harish/personal/hypefx/examples/app",
		BaseImportPath:      "github.com/harish876/hypefx/examples/app",
		TemplateName:        "routes",
		RouteDirPackageName: "routes",
		DestinationDir:      "/home/harish/personal/hypefx/examples/routes/routes.go",
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
