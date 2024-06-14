package version

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
)

var (
	MESSAGE string = "Hype - 0.0.2"
)

func Welcome(cmd *cobra.Command, args []string) {
	version, _ := cmd.Flags().GetBool("version")
	slog.Debug("welcome", "version", version)
	model := FromMessage(MESSAGE)
	if err := Display(model); err != nil {
		fmt.Println(MESSAGE)
	}
}
