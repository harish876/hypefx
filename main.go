package main

import (
	"fmt"
	"os"

	"github.com/harish876/hypefx/cli/commands"
	generate "github.com/harish876/hypefx/cli/commands/generate"
	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{
		Use:   "hello",
		Short: "A simple CLI tool to greet the user",
		Run:   commands.Welcome,
	}

	var generateCmd = &cobra.Command{
		Use:     "generate [foldername]",
		Short:   "Generates a new HypeFX Project Structure",
		Long:    `Generates a new HypeFX Project Structure, when a base path to the project i.e the go mod base path is provided.`,
		Args:    cobra.ExactArgs(1), // Require exactly one argument (project_name)
		Example: "hype generate foobar",
		Run:     generate.Generate,
	}

	rootCmd.Flags().StringP("name", "n", "", "Name of the person to greet")
	rootCmd.Flags().StringP("generate", "g", "", "Create a Folder according to the arguement")
	rootCmd.AddCommand(generateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
