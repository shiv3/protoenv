package protoc

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type Versions struct {
	InstallDirectoryPath string
}

func NewVersions(parentCmd *cobra.Command, installDirectoryPath string) Versions {
	versions := Versions{
		InstallDirectoryPath: getVersionsPath(installDirectoryPath),
	}
	cmd := &cobra.Command{
		Use:   "versions",
		Short: "List all Go versions available to protoenv",
		Long:  `List all Go versions available to protoenv`,
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
			fmt.Printf(ShowVersionFormatSimple, dir.Name())
		}
	}
	return nil
}
