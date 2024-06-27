package set

import (
	"fmt"
	"log/slog"

	"github.com/harish876/hypefx/internal/cli/commands"
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
			name:      "app_dir",
			shorthand: "a",
			usage:     "Set the Directory of your app directory. Default [pwd]/app",
		},
		{
			typ:       STRING,
			name:      "module",
			shorthand: "m",
			usage:     "Set the Go Module name of your project",
		},
		{
			typ:       STRING,
			name:      "route_dir",
			shorthand: "p",
			usage:     "Set the Base file path for your routes directory. Default [pwd]/routes/routes.go ",
		},
		{
			typ:       BOOL,
			name:      "routing",
			shorthand: "r",
			usage:     "Pass in the flag if you want to enable file based routing and automagic route generation.Default true",
		},
	}
)

func setConfig(cmd *cobra.Command, args []string) {
	for _, option := range OPTIONS {
		switch option.typ {
		case STRING:
			flag, _ := cmd.Flags().GetString(option.name)
			if flag != "" {
				if err := commands.UpsertConfig(option.name, flag); err != nil {
					slog.Error("setConfig", option.name, flag, err)
					DisplayError(fmt.Errorf("unable to set %s with value %s", option.name, flag))
				} else {
					DisplaySuccessMessage(fmt.Sprintf("Successfully set %s as %s", option.name, flag))
				}

			}
		case BOOL:
			flag, _ := cmd.Flags().GetBool(option.name)
			if flag {
				if err := commands.UpsertConfig(option.name, flag); err != nil {
					slog.Error("setConfig", option.name, flag, err)
					DisplayError(fmt.Errorf("unable to set %s with value %v", option.name, flag))
				} else {
					DisplaySuccessMessage(fmt.Sprintf("Successfully set %s as %v", option.name, flag))
				}
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
