//go:build test

package main

import (
	"embed"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/harish876/hypefx/internal/cli/commands"
	"github.com/harish876/hypefx/internal/cli/commands/add"
	"github.com/harish876/hypefx/internal/cli/commands/build"
	"github.com/harish876/hypefx/internal/cli/commands/generate"
	"github.com/harish876/hypefx/internal/cli/commands/set"
	"github.com/harish876/hypefx/internal/cli/commands/unset"
	"github.com/harish876/hypefx/internal/cli/commands/version"
	"github.com/harish876/hypefx/internal/utils"
	"github.com/spf13/cobra"
)

//go:embed components/*
var components embed.FS

func main() {
	loggerConfig := utils.FromConfig(filepath.Join("tmp", "hypefx.log"))
	jsonHandler := slog.NewJSONHandler(loggerConfig, &slog.HandlerOptions{
		AddSource: true,
		Level:     utils.FromLogLevel("debug"),
	})
	slog.SetDefault(slog.New(jsonHandler))

	rootCmd := &cobra.Command{
		Use:   "hype-test",
		Short: "A simple CLI tool to bootstrap Go + HTMX + Templ Projects.",
		Run: func(cmd *cobra.Command, args []string) {
			version.Welcome(cmd, args, map[string]string{
				"buildEnv": "test",
			})
		},
		ValidArgs: []string{"generate", "build", "set", "unset", "add"},
	}

	//this needs access to components.to be refactored
	var addCmd = &cobra.Command{
		Use:       "add [compoent_name] [project_name/module_name](optional)",
		Short:     "Add a new component from the component library",
		Long:      `Add a new component from the component library , and customise it as per your liking`,
		Args:      cobra.ExactArgs(1),
		Example:   "hype add grid",
		ValidArgs: []string{"notfound", "grid", "input", "dropdown", "notification"},
		Run: func(cmd *cobra.Command, args []string) {
			add.Add(cmd, args, components)
		},
	}
	//Init Command Flags
	slog.Info("CLI Initialization done. Build tag is test")
	rootCmd.Flags().BoolP("version", "v", false, "Display CLI Version")
	set.InitFlags()

	//add autocompletion
	addCmd.CompletionOptions.DisableDefaultCmd = true

	// Add Command
	rootCmd.AddCommand(generate.GenerateCmd)
	rootCmd.AddCommand(set.SetCmd)
	rootCmd.AddCommand(unset.UnsetCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(build.BuildCmd)
	rootCmd.AddCommand(commands.CompletionCmd)

	if err := rootCmd.Execute(); err != nil {
		slog.Error("rootCmd", err)
		os.Exit(1)
	}
}
