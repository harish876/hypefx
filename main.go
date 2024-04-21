package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate [foldername]",
	Short: "Create a new folder",
	Long:  `Create a new folder with the specified name.`,
	Run: func(cmd *cobra.Command, args []string) {
		folders := []string{"assets", "cmd"}
		for _, folder := range folders {
			err := os.Mkdir(folder, 0755)
			switch folder {
			case "cmd":
				dir := "cmd"
				filePath := filepath.Join(dir, "main.go")
				existingFilePath := "./examples/cmd/main.go"
				content, err := os.ReadFile(existingFilePath)
				if err != nil {
					fmt.Println(err)
				}
				if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
					fmt.Println(err)
				}

				fmt.Println("main.go created successfully at", filePath)

			}
			if err != nil {
				fmt.Println("Error creating folder:", err)
				return
			}
		}
		fmt.Println("Successfully Instantiated a Hype FX Project")
	},
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "hello",
		Short: "A simple CLI tool to greet the user",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				name = "HypeFX User"
			}
			fmt.Printf("Hello, %s!\n", name)
		},
	}

	rootCmd.Flags().StringP("name", "n", "", "Name of the person to greet")
	rootCmd.Flags().StringP("generate", "g", "", "Create a Folder according to the arguement")
	rootCmd.AddCommand(generateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
