package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type Global struct {
	InstallDirectoryPath    string
	ShowVersionFormatSimple string
	TargetBinaryFileName    string
}

func NewGlobal(parentCmd *cobra.Command, installDirectoryPath string, ShowVersionFormatSimple string, TargetBinaryFileName string) Global {
	global := Global{
		InstallDirectoryPath:    installDirectoryPath,
		ShowVersionFormatSimple: ShowVersionFormatSimple,
		TargetBinaryFileName:    TargetBinaryFileName,
	}
	cmd := &cobra.Command{
		Use:   "global",
		Short: fmt.Sprintf("Set or show the global %s version", TargetBinaryFileName),
		Long:  fmt.Sprintf(`Set or show the global %s version`, TargetBinaryFileName),
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
		return setShims(GetShimsFileDir(i.InstallDirectoryPath), i.TargetBinaryFileName, i.InstallDirectoryPath, version)
	}
	v, err := getVersion(getGlobalVersionFilePath(i.InstallDirectoryPath))
	if err != nil {
		return err
	}
	fmt.Printf(i.ShowVersionFormatSimple, v)
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
