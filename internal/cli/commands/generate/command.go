package generate

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/harish876/hypefx/internal/cli/commands"
	"github.com/harish876/hypefx/internal/utils"
	"github.com/spf13/cobra"
)

var (
	BASE_PATH = "github.com/harish876/hypefx/internal/cli/commands/generate/scaffolding"
)

var GenerateCmd = &cobra.Command{
	Use:     "generate [project_name/module_name](optional)",
	Short:   "Generates a new HypeFX Project Structure",
	Long:    `Generates a new HypeFX Project Structure, when a base path to the project i.e the go mod base path is provided.`,
	Example: "hype generate",
	Run:     generate,
}

//go:embed scaffolding/*
var embeddedFiles embed.FS

func generate(cmd *cobra.Command, args []string) {
	view := NewModel()
	var message string
	var errorInterface error
	var moduleName interface{}
	if len(args) >= 1 {
		moduleName = args[0]
		commands.UpsertConfig("module", moduleName.(string))
	} else {
		moduleName, err := commands.GetConfig("module")
		if err != nil {
			errorInterface = errors.Join(err, fmt.Errorf("unable to find module name. use hype set --module [module_name]"))
			DisplayError(errorInterface)
			return
		}
		if moduleName == nil && len(args) == 0 {
			errorInterface = fmt.Errorf("unable to find module name. use hype generate [module_name] to automagically perform the initializations")
			DisplayError(errorInterface)
			return
		}
	}
	initCmd := exec.Command("go", "mod", "init", moduleName.(string))
	_, err := initCmd.CombinedOutput()
	if err != nil {
		view.errorInterface = fmt.Errorf("error initializing Go module. run go mod init [module_name] manually")
		DisplayError(errorInterface)
		return
	}
	if err := copyDirectory("scaffolding", ".", moduleName.(string)); err != nil {
		view.errorInterface = fmt.Errorf("unable to scaffold project")
		DisplayError(errorInterface)
		return
	}
	message = fmt.Sprintf("Successfully Instantiated a Hype FX Project in module %s", moduleName.(string))
	DisplayMessage(message)
}

func copyDirectory(src, dst, moduleName string) error {
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
			if err := copyDirectory(sourcePath, destPath, moduleName); err != nil {
				return err
			}
		} else {
			if err := copyFile(sourcePath, destPath, moduleName); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(src, dst, moduleName string) error {
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
	utils.ReplaceFileContent(dst, BASE_PATH, moduleName)
	return nil
}
