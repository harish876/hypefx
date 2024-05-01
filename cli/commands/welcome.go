package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Welcome(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	if name == "" {
		name = "HypeFX Version 0"
	}
	fmt.Printf("Hello, %s!\n", name)
}
