package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/harish876/hypefx/cli/commands/add"
	"github.com/harish876/hypefx/cli/commands/generate"
	"github.com/harish876/hypefx/cli/commands/set"
	"github.com/harish876/hypefx/cli/commands/unset"
	"github.com/harish876/hypefx/cli/commands/version"
	"github.com/spf13/cobra"
)

//go:embed components/*
var components embed.FS

func main() {

	rootCmd := &cobra.Command{
		Use:   "hypefx",
		Short: "A simple CLI tool to bootstrap Go + HTMX + Templ Projects.",
		Run:   version.Welcome,
	}

	var addCmd = &cobra.Command{
		Use:     "add [compoent_name] [project_name/module_name]",
		Short:   "Add a new component from the component library",
		Long:    `Add a new component from the component library , and customise it as per your liking`,
		Args:    cobra.ExactArgs(2), // Require exactly two arguments (component_name, module_name)
		Example: "hype add grid foobar",
		Run: func(cmd *cobra.Command, args []string) {
			add.Add(cmd, args, components)
		},
	}
	//Init Command Flags
	rootCmd.Flags().BoolP("version", "v", false, "Display CLI Version")
	set.InitFlags()

	// Add Commands
	rootCmd.AddCommand(generate.GenerateCmd)
	rootCmd.AddCommand(set.SetCmd)
	rootCmd.AddCommand(unset.UnsetCmd)
	rootCmd.AddCommand(addCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
