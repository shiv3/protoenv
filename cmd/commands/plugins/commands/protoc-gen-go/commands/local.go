package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type Local struct {
	InstallDirectoryPath    string
	ShowVersionFormatSimple string
	TargetBinaryFileName    string
}

func NewLocal(parentCmd *cobra.Command, installDirectoryPath string, ShowVersionFormatSimple string, TargetBinaryFileName string) Local {
	local := Local{
		InstallDirectoryPath:    installDirectoryPath,
		ShowVersionFormatSimple: ShowVersionFormatSimple,
		TargetBinaryFileName:    TargetBinaryFileName,
	}
	cmd := &cobra.Command{
		Use:   "local",
		Short: fmt.Sprintf("Set or show the local %s version", TargetBinaryFileName),
		Long:  fmt.Sprintf(`Set or show the local %s version`, TargetBinaryFileName),
		RunE:  local.RunE,
	}
	parentCmd.AddCommand(cmd)
	return local
}

func (i *Local) RunE(cmd *cobra.Command, args []string) error {
	currentDirectory, err := os.Getwd()
	if err != nil {
		return err
	}
	localVersionFilePath := getLocalVersionFilePath(currentDirectory)
	if len(args) > 0 {
		return setVersion(getVersionsPath(i.InstallDirectoryPath), localVersionFilePath, args[0])
	}
	if v, err := getVersion(localVersionFilePath); err != nil {
		return err
	} else {
		fmt.Printf(i.ShowVersionFormatSimple, v)
	}
	return nil
}
