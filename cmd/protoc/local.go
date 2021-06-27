package protoc

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type Local struct {
	InstallDirectoryPath string
}

func NewLocal(parentCmd *cobra.Command, installDirectoryPath string) Local {
	local := Local{
		InstallDirectoryPath: installDirectoryPath,
	}
	cmd := &cobra.Command{
		Use:   "local",
		Short: "Set or show the local Go version",
		Long:  `Set or show the local Go version`,
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
		fmt.Printf(ShowVersionFormatSimple, v)
	}
	return nil
}
