package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Welcome(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	if name == "" {
		name = "HypeFX User-1"
	}
	fmt.Printf("Hello, %s!\n", name)
}
