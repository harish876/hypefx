package main

import (
	"embed"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/harish876/hypefx/internal/cli/commands/add"
	"github.com/harish876/hypefx/internal/cli/commands/generate"
	"github.com/harish876/hypefx/internal/cli/commands/set"
	"github.com/harish876/hypefx/internal/cli/commands/unset"
	"github.com/harish876/hypefx/internal/utils"
	"github.com/harish876/hypefx/internal/cli/commands/version"
	"github.com/spf13/cobra"
)

//go:embed components/*
var components embed.FS

func main() {
	loggerConfig := utils.FromConfig(filepath.Join("tmp", "hypefx.log"))
	jsonHandler := slog.NewJSONHandler(loggerConfig, &slog.HandlerOptions{
		Level: utils.FromLogLevel("Debug"),
	})
	slog.SetDefault(slog.New(jsonHandler))

	rootCmd := &cobra.Command{
		Use:   "hypefx",
		Short: "A simple CLI tool to bootstrap Go + HTMX + Templ Projects.",
		Run:   version.Welcome,
	}

	//this needs access to components.to be refactored
	var addCmd = &cobra.Command{
		Use:     "add [compoent_name] [project_name/module_name](optional)",
		Short:   "Add a new component from the component library",
		Long:    `Add a new component from the component library , and customise it as per your liking`,
		Args:    cobra.ExactArgs(1), // Require exactly two arguments (component_name, module_name)
		Example: "hype add grid foobar",
		Run: func(cmd *cobra.Command, args []string) {
			add.Add(cmd, args, components)
		},
	}
	//Init Command Flags
	slog.Info("CLI Initialization done")
	rootCmd.Flags().BoolP("version", "v", false, "Display CLI Version")
	set.InitFlags()

	// Add Commands
	rootCmd.AddCommand(generate.GenerateCmd)
	rootCmd.AddCommand(set.SetCmd)
	rootCmd.AddCommand(unset.UnsetCmd)
	rootCmd.AddCommand(addCmd)

	if err := rootCmd.Execute(); err != nil {
		slog.Error("rootCmd", err)
		os.Exit(1)
	}
}
