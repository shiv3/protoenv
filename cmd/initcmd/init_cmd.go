package initcmd

import (
	"fmt"
	"github.com/shiv3/protoenv/config"
	"github.com/spf13/cobra"
	"os"
)

type Init struct {
	RootDirectoryPath string
}

func NewInit(parentCmd *cobra.Command, config config.Config) *Init {
	init := &Init{RootDirectoryPath: config.RootDirectoryPath}
	cmd := &cobra.Command{
		Use:   "init",
		Short: "A brief description of your command",
		Long:  ``,
		RunE:  init.RunE,
	}
	parentCmd.AddCommand(cmd)
	return init
}

func (i *Init) RunE(cmd *cobra.Command, args []string) error {
	if _, err := os.Stat(i.RootDirectoryPath); os.IsNotExist(err) {
		return os.Mkdir(i.RootDirectoryPath, os.ModePerm)
	}
	fmt.Printf(`
export PROTOENV_ROOT=%s
`, i.RootDirectoryPath)
	return nil
}
