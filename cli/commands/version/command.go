package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	MESSAGE string = "Hype - 0.0.1"
)

func Welcome(cmd *cobra.Command, args []string) {
	version, _ := cmd.Flags().GetBool("version")
	_ = version
	model := FromMessage(MESSAGE)
	if err := Display(model); err != nil {
		fmt.Println(MESSAGE)
	}
}
