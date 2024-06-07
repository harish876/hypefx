package set

import (
	"github.com/harish876/hypefx/cli/commands"
	"github.com/spf13/cobra"
)

const (
	DEFAULT_STRING_VAL = ""
	DEFAULT_BOOL_VAL   = false
)

const (
	STRING = iota
	BOOL
)

type Flags struct {
	typ       int
	name      string
	shorthand string
	usage     string
}

var (
	OPTIONS = []Flags{
		{
			typ:       STRING,
			name:      "module",
			shorthand: "m",
			usage:     "Set the Go Module name of your project",
		},
		{
			typ:       BOOL,
			name:      "routing",
			shorthand: "r",
			usage:     "Pass in the flag if you want to enable file based routing",
		},
		{
			typ:       STRING,
			name:      "file",
			shorthand: "f",
			usage:     "Pass in the path to the log file to see logs",
		},
		{
			typ:       STRING,
			name:      "level",
			shorthand: "l",
			usage:     "Pass in the log level [INFO | DEBUG | ERROR]",
		},
	}
)

func setConfig(cmd *cobra.Command, args []string) {
	for _, option := range OPTIONS {
		switch option.typ {
		case STRING:
			flag, _ := cmd.Flags().GetString(option.name)
			if flag != "" {
				commands.UpsertConfig(option.name, flag)
			}
		case BOOL:
			flag, _ := cmd.Flags().GetBool(option.name)
			if flag {
				commands.UpsertConfig(option.name, flag)
			}
		}
	}
}

var SetCmd = &cobra.Command{
	Use:     "set",
	Short:   "Set a config option globally",
	Long:    `Creates a hypeconfig.json file, where options can be persisted globally for ease of use. Use --help to check the allowed flags`,
	Example: "hype set --module github.com/harish876/hypefx",
	Run:     setConfig,
}

func InitFlags() {
	for _, option := range OPTIONS {
		switch option.typ {
		case STRING:
			SetCmd.Flags().StringP(option.name, option.shorthand, DEFAULT_STRING_VAL, option.usage)
		case BOOL:
			SetCmd.Flags().BoolP(option.name, option.shorthand, DEFAULT_BOOL_VAL, option.usage)
		default:
			SetCmd.Flags().StringP(option.name, option.shorthand, DEFAULT_STRING_VAL, option.usage)
		}

	}
}
