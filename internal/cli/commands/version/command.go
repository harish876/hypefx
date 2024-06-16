package version

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
)

var (
	MESSAGE string = "Hype"
	VERSION string = "0.0.3"
)

func Welcome(cmd *cobra.Command, args []string, buildFlags map[string]string) {
	var metdata string
	version, _ := cmd.Flags().GetBool("version")
	slog.Debug("welcome", "version", version)
	description := "Build Simple Web Apps using Go,Tailwind,Templ and HTMX"
	if tag, ok := buildFlags["buildEnv"]; ok {
		metdata = fmt.Sprintf("\n1.Running Version - %s\n2.Build Tag - %s\n3.Run hypefx completion --help for autocomplete setup", VERSION, tag)
	}
	model := FromMessage(MESSAGE+" - "+VERSION, description, metdata)
	if err := Display(model); err != nil {
		//fallback
		fmt.Println(MESSAGE + " - " + VERSION)
	}
}
