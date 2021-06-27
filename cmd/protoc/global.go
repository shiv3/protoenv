package protoc

import (
	"fmt"
	"github.com/spf13/cobra"
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
		return setVersion(getVersionsPath(i.InstallDirectoryPath), getGlobalVersionFilePath(i.InstallDirectoryPath), args[0])
	}
	if v, err := getVersion(getGlobalVersionFilePath(i.InstallDirectoryPath)); err != nil {
		return err
	} else {
		fmt.Printf(ShowVersionFormatSimple, v)
	}
	return nil
}
