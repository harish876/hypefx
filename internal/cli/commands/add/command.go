package add

import (
	"embed"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/harish876/hypefx/internal/cli/commands"
	"github.com/harish876/hypefx/internal/utils"
	"github.com/spf13/cobra"
)

var (
	BASE_PATH   = "github.com/harish876/hypefx"
	MODULE_NAME string
)

func Add(cmd *cobra.Command, args []string, components embed.FS) {
	componentName := args[0]
	moduleName, err := commands.GetConfig("module")
	if err != nil || moduleName == nil {
		slog.Error("add", err)
	}
	if len(args) >= 1 {
		moduleName = args[0]
		commands.UpsertConfig("module", moduleName)
	}
	MODULE_NAME = moduleName.(string)
	currDir, err := os.Getwd()
	if err != nil {
		slog.Error("add", "Error Getting the current directory: ", err)
		os.Exit(1)
	}
	componentPath := filepath.Join(currDir, "components")
	copyDirectory("components", componentPath, components, componentName)
	fmt.Printf("Component %s Added. Dont forget to add other dependent components\n", componentName)
}

func copyDirectory(src, dst string, components embed.FS, componentName string) error {
	files, err := components.ReadDir(src)
	if err != nil {
		return err
	}
	for _, file := range files {
		sourcePath := filepath.Join(src, file.Name())
		destPath := filepath.Join(dst, file.Name())

		if !strings.Contains(sourcePath, componentName) {
			continue
		}

		if file.IsDir() {
			if err := os.MkdirAll(destPath, 0755); err != nil {
				return err
			}
			if err := copyDirectory(sourcePath, destPath, components, componentName); err != nil {
				return err
			}
		} else {
			if err := copyFile(sourcePath, destPath, components); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(src, dst string, components embed.FS) error {
	sourceFile, err := components.Open(src)
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
	// TODO: make this better -> hardcoded. use golang templates here
	utils.ReplaceFileContent(dst, BASE_PATH, MODULE_NAME)
	return nil
}
