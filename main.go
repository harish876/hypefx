package main

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var source string

//go:embed scaffolding/*
var embeddedFiles embed.FS

var generateCmd = &cobra.Command{
	Use:   "generate [foldername]",
	Short: "Create a new folder",
	Long:  `Create a new folder with the specified name.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := copyDirectory("scaffolding", ".")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Successfully Instantiated a Hype FX Project")
	},
}

func copyDirectory(src, dst string) error {
	files, err := embeddedFiles.ReadDir(src)
	if err != nil {
		return err
	}

	for _, file := range files {
		sourcePath := filepath.Join(src, file.Name())
		destPath := filepath.Join(dst, file.Name())

		if file.IsDir() {
			if err := os.MkdirAll(destPath, 0755); err != nil {
				return err
			}
			if err := copyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(sourcePath, destPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	sourceFile, err := embeddedFiles.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}
	return nil
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
	rootCmd.Flags().StringVarP(&source, "source", "s", "", "The source for generation")
	rootCmd.AddCommand(generateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
