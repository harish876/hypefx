package main

import (
	"embed"
	"log"
	"os"
	"path/filepath"

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
	logger, err := utils.NewLogger(filepath.Join("hypefx.log"))

	if err != nil {
		log.Fatalf("Unable to initialise logger %v", err)
	}

	rootCmd := &cobra.Command{
		Use:   "hypefx",
		Short: "A simple CLI tool to bootstrap Go + HTMX + Templ Projects.",
		Run:   version.Welcome,
	}

	//this needs access to components
	var addCmd = &cobra.Command{
		Use:     "add [compoent_name] [project_name/module_name](optional)",
		Short:   "Add a new component from the component library",
		Long:    `Add a new component from the component library , and customise it as per your liking`,
		Args:    cobra.ExactArgs(1),
		Example: "hype add grid",
		Run: func(cmd *cobra.Command, args []string) {
			add.Add(cmd, args, components)
		},
	}
	//Init Command Flags
	logger.Info("CLI Initialization done")
	rootCmd.Flags().BoolP("version", "v", false, "Display CLI Version")
	set.InitFlags()

	// Add Commands
	rootCmd.AddCommand(generate.GenerateCmd)
	rootCmd.AddCommand(set.SetCmd)
	rootCmd.AddCommand(unset.UnsetCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(build.BuildCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.Error("rootCmd", err)
		os.Exit(1)
	}
}
