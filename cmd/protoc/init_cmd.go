package protoc

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type Init struct {
	InstallDirectoryPath string
}

func NewInit(parentCmd *cobra.Command, installDirectoryPath string) Init {
	init := Init{
		InstallDirectoryPath: installDirectoryPath,
	}
	cmd := &cobra.Command{
		Use:   "init",
		Short: fmt.Sprintf("Set or show the global Go version",TargetBinaryFileName),
		Long:  fmt.Sprintf(`Set or show the global Go version`,TargetBinaryFileName),
		RunE:  init.RunE,
	}
	parentCmd.AddCommand(cmd)
	return init
}

func (i *Init) RunE(cmd *cobra.Command, args []string) error {
	if _, err := os.Stat(i.InstallDirectoryPath); os.IsNotExist(err) {
		return os.Mkdir(i.InstallDirectoryPath, os.ModePerm)
	}
	fmt.Printf(`export PATH=$PATH:%s`, GetShimsFileDir(i.InstallDirectoryPath))
	return nil
}
