package protoc

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

type Version struct {
	GlobalVersionFilePath string
}

func NewVersion(parentCmd *cobra.Command, installDirectoryPath string) Version {
	version := Version{
		GlobalVersionFilePath: getGlobalVersionFilePath(installDirectoryPath),
	}
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the current protoc version and its origin",
		Long:  `Show the current protoc version and its origin`,
		RunE:  version.RunE,
	}
	parentCmd.AddCommand(cmd)
	return version
}

func getGlobalVersionFilePath(installDirectoryPath string) string {
	return filepath.Join(installDirectoryPath, "version")
}

func (i *Version) RunE(cmd *cobra.Command, args []string) error {
	v, err := getGlobalVersion(i.GlobalVersionFilePath)
	if err != nil {
		return err
	}
	fmt.Printf(ShowVersionFormatSimple, v)
	return nil
}
