package generate

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/harish876/hypefx/cli/commands"
	"github.com/harish876/hypefx/cli/commands/utils"
	"github.com/spf13/cobra"
)

var (
	MODULE_NAME string
	BASE_PATH   = "github.com/harish876/hypefx/cli/commands/generate/scaffolding"
)

//go:embed scaffolding/*
var embeddedFiles embed.FS

func generate(cmd *cobra.Command, args []string) {
	moduleName, err := commands.GetConfig("module")
	if err != nil {
		fmt.Println(err)
	}
	if len(args) >= 1 {
		moduleName = args[0]
		commands.UpsertConfig("module", moduleName)
	}
	MODULE_NAME = moduleName.(string)
	if err := copyDirectory("scaffolding", "."); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully Instantiated a Hype FX Project in module %s\n", MODULE_NAME)
}

var GenerateCmd = &cobra.Command{
	Use:     "generate [project_name/module_name](optional)",
	Short:   "Generates a new HypeFX Project Structure",
	Long:    `Generates a new HypeFX Project Structure, when a base path to the project i.e the go mod base path is provided.`,
	Example: "hype generate foobar",
	Run:     generate,
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
	// TODO: make this better
	if MODULE_NAME == "" {
		fmt.Println("Module Name is required. It is not set")
		os.Exit(1)
	}
	utils.ReplaceFileContent(dst, BASE_PATH, MODULE_NAME)
	return nil
}
