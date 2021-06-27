package protoc

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type Global struct {
	InstallDirectoryPath string
}

func NewGlobal(parentCmd *cobra.Command, installDirectoryPath string) Global {
	global := Global{
		InstallDirectoryPath: installDirectoryPath,
	}
	cmd := &cobra.Command{
		Use:   "global",
		Short: "Set or show the global Go version",
		Long:  `Set or show the global Go version`,
		RunE:  global.RunE,
	}
	parentCmd.AddCommand(cmd)
	return global
}

func (i *Global) RunE(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		version := args[0]
		if err := setVersion(getVersionsPath(i.InstallDirectoryPath), getGlobalVersionFilePath(i.InstallDirectoryPath), version); err != nil {
			return err
		}
		return setShims(GetShimsFileDir(i.InstallDirectoryPath), TargetBinaryFileName, i.InstallDirectoryPath, version)
	}
	v, err := getVersion(getGlobalVersionFilePath(i.InstallDirectoryPath))
	if err != nil {
		return err
	}
	fmt.Printf(ShowVersionFormatSimple, v)
	return nil
}

func setShims(shimsDir, targetFileName, installPath, version string) error {
	if _, err := os.Stat(filepath.Join(shimsDir)); os.IsNotExist(err) {
		err = os.MkdirAll(shimsDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	shimsPath := filepath.Join(shimsDir, targetFileName)
	if _, err := os.Lstat(shimsPath); err == nil {
		if err := os.Remove(shimsPath); err != nil {
			return err
		}
	}
	return nil
}
