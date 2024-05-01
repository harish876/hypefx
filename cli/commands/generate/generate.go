package commands

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/harish876/hypefx/cli/commands/generate/utils"
	"github.com/spf13/cobra"
)

var (
	BASE_PATH = "github.com/harish876/hypefx/cli/commands/generate/scaffolding"
)

//go:embed scaffolding/*
var embeddedFiles embed.FS

func Generate(cmd *cobra.Command, args []string) {
	projectName := args[0]
	if projectName == "" {
		fmt.Println("Please specify a project name. This would be the 'go mod init' module name")
		os.Exit(1)
	}
	err := copyDirectory("scaffolding", ".", projectName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Successfully Instantiated a Hype FX Project")
}

func copyDirectory(src, dst, projectName string) error {
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
			if err := copyDirectory(sourcePath, destPath, projectName); err != nil {
				return err
			}
		} else {
			if err := copyFile(sourcePath, destPath, projectName); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(src, dst, projectName string) error {
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
	// TODO: make this better
	utils.ReplaceFileContent(dst, BASE_PATH, projectName)
	return nil
}
