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
		Use:   "generate [foldername]",
		Short: "Create a new folder",
		Long:  `Create a new folder with the specified name.`,
		Run:   generate.Generate,
	}

	rootCmd.Flags().StringP("name", "n", "", "Name of the person to greet")
	rootCmd.Flags().StringP("generate", "g", "", "Create a Folder according to the arguement")
	rootCmd.AddCommand(generateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
