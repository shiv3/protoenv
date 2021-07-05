package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type Versions struct {
	InstallDirectoryPath    string
	ShowVersionFormatSimple string
	TargetBinaryFileName    string
}

func NewVersions(parentCmd *cobra.Command, installDirectoryPath string, ShowVersionFormatSimple string, TargetBinaryFileName string) Versions {
	versions := Versions{
		InstallDirectoryPath:    getVersionsPath(installDirectoryPath),
		ShowVersionFormatSimple: ShowVersionFormatSimple,
		TargetBinaryFileName:    TargetBinaryFileName,
	}
	cmd := &cobra.Command{
		Use:   "versions",
		Short: fmt.Sprintf("List all %s versions available to protoenv", TargetBinaryFileName),
		Long:  fmt.Sprintf(`List all %s versions available to protoenv`, TargetBinaryFileName),
		RunE:  versions.RunE,
	}
	parentCmd.AddCommand(cmd)
	return versions
}

func getVersionsPath(installDirectoryPath string) string {
	return filepath.Join(installDirectoryPath, "versions")
}

func (i *Versions) RunE(cmd *cobra.Command, args []string) error {
	if _, err := os.Stat(i.InstallDirectoryPath); os.IsNotExist(err) {
		return fmt.Errorf("versions dir not found: %w", err)
	}
	dirs, err := os.ReadDir(i.InstallDirectoryPath)
	if err != nil {
		return err
	}

	// only show local directories
	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Printf(i.ShowVersionFormatSimple, dir.Name())
		}
	}
	return nil
}
