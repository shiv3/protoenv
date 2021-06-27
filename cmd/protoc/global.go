package protoc

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

type Global struct {
	GlobalVersionFilePath string
}

func NewGlobal(parentCmd *cobra.Command, installDirectoryPath string) Global {
	global := Global{
		GlobalVersionFilePath: getGlobalVersionFilePath(installDirectoryPath),
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
		return setGlobalVersion(i.GlobalVersionFilePath, args[0])
	}
	if v, err := getGlobalVersion(i.GlobalVersionFilePath); err != nil {
		return err
	} else {
		fmt.Printf(ShowVersionFormatSimple, v)
	}
	return nil
}

func getGlobalVersion(globalVersionFilePath string) (string, error) {
	if _, err := os.Stat(globalVersionFilePath); os.IsNotExist(err) {
		return "", fmt.Errorf("please install protoc (cannot find protoc versions: %w)", err)
	}
	b, err := os.ReadFile(globalVersionFilePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func setGlobalVersion(globalVersionFilePath string, version string) error {
	err := ioutil.WriteFile(globalVersionFilePath, []byte(version), 0644)
	if err != nil {
		return err
	}
	fmt.Printf("set global version: %s\n", version)
	return nil
}
