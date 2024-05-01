package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/harish876/hypefx/cli/commands"
	"github.com/harish876/hypefx/cli/commands/add"
	generate "github.com/harish876/hypefx/cli/commands/generate"
	"github.com/spf13/cobra"
)

//go:embed components/*
var components embed.FS

func main() {

	rootCmd := &cobra.Command{
		Use:   "hello",
		Short: "A simple CLI tool to greet the user",
		Run:   commands.Welcome,
	}

	var generateCmd = &cobra.Command{
		Use:     "generate [project_name]",
		Short:   "Generates a new HypeFX Project Structure",
		Long:    `Generates a new HypeFX Project Structure, when a base path to the project i.e the go mod base path is provided.`,
		Args:    cobra.ExactArgs(1), // Require exactly one argument (project_name)
		Example: "hype generate foobar",
		Run:     generate.Generate,
	}

	var addCmd = &cobra.Command{
		Use:     "add [compoent_name]",
		Short:   "Add a new component from the component library",
		Long:    `Add a new component from the component library , and customise it as per your liking`,
		Args:    cobra.ExactArgs(1), // Require exactly one argument (component_name)
		Example: "hype add grid",
		Run: func(cmd *cobra.Command, args []string) {
			add.Add(cmd, args, components)
		},
	}

	rootCmd.Flags().StringP("name", "n", "", "Name of the person to greet")
	rootCmd.Flags().StringP("generate", "g", "", "Create a Folder according to the arguement")
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(addCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
