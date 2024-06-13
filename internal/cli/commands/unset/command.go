package unset

import (
	"github.com/harish876/hypefx/internal/cli/commands"
	"github.com/spf13/cobra"
)

var UnsetCmd = &cobra.Command{
	Use:     "unset [option_name]",
	Short:   "unset a config option globally by passing the option name",
	Long:    `Removes the config option from the hypeconfig.json file`,
	Args:    cobra.ExactArgs(1),
	Example: "hype unset module",
	Run:     unsetConfig,
}

func unsetConfig(cmd *cobra.Command, args []string) {
	key := args[0]
	commands.DeleteConfig(key)
}
