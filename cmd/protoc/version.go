package protoc

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

var versionFileName = ".protoc-version"

type Version struct {
	GlobalVersionFilePath string
}

func NewVersion(parentCmd *cobra.Command, installDirectoryPath string) Version {
	version := Version{
		GlobalVersionFilePath: getGlobalVersionFilePath(installDirectoryPath),
	}
	cmd := &cobra.Command{
		Use:   "version",
		Short: fmt.Sprintf("Show the current %s version and its origin", TargetBinaryFileName),
		Long:  fmt.Sprintf(`Show the current %s version and its origin`, TargetBinaryFileName),
		RunE:  version.RunE,
	}
	parentCmd.AddCommand(cmd)
	return version
}

func getLocalVersionFilePath(currentDirectory string) string {
	return filepath.Join(currentDirectory, versionFileName)
}
func getGlobalVersionFilePath(installDirectoryPath string) string {
	return filepath.Join(installDirectoryPath, versionFileName)
}

func (i *Version) RunE(cmd *cobra.Command, args []string) error {
	v, err := getVersion(i.GlobalVersionFilePath)
	if err != nil {
		return err
	}
	fmt.Printf(ShowVersionFormatSimple, v)
	return nil
}

func getVersion(versionFilePath string) (string, error) {
	if _, err := os.Stat(versionFilePath); os.IsNotExist(err) {
		return "", fmt.Errorf("version file not found: %w", err)
	}
	b, err := os.ReadFile(versionFilePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func setVersion(versionsPath, versionFilePath string, version string) error {
	if _, err := os.Stat(filepath.Join(versionsPath, version)); os.IsNotExist(err) {
		return fmt.Errorf("goenv: version '%s' not installed", version)
	}
	err := ioutil.WriteFile(versionFilePath, []byte(version), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("set global version: %s\n", version)
	return nil
}

func GetShimsFileDir(installFilePath string) string {
	return filepath.Join(installFilePath, "shims")
}
