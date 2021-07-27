package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/shiv3/protoenv/adapter/goinstall"

	protoc_gen_go "github.com/shiv3/protoenv/adapter/github/protoc-gen-go"

	"github.com/spf13/cobra"
)

type Install struct {
	InstallDirectoryPath    string
	ShowVersionFormatSimple string
	TargetBinaryFileName    string
}

func NewInstall(parentCmd *cobra.Command, installDirectoryPath string, ShowVersionFormatSimple string, TargetBinaryFileName string) Install {
	install := Install{
		InstallDirectoryPath:    installDirectoryPath,
		ShowVersionFormatSimple: ShowVersionFormatSimple,
		TargetBinaryFileName:    TargetBinaryFileName,
	}
	cmd := &cobra.Command{
		Use:   "install (version)",
		Short: "install specified version",
		Long:  `install specified version`,
		RunE:  install.RunE,
	}
	cmd.PersistentFlags().BoolP("list", "l", false, "show install list flag")
	parentCmd.AddCommand(cmd)
	return install
}

func (i Install) RunE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	//c.installOptions.ShowVersionList {
	if list, err := cmd.PersistentFlags().GetBool("list"); err == nil && list {
		if err := i.showVersion(ctx); err != nil {
			return err
		}
		return nil
	}

	if len(args) >= 1 {
		err := i.installVersion(ctx, args[0])
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("requires a installing version or some flags")
}

func (i Install) showVersion(ctx context.Context) error {
	versions, err := protoc_gen_go.GetprotocGenGoRepoGoVersions(ctx)
	if err != nil {
		return err
	}
	for _, version := range versions {
		fmt.Printf(i.ShowVersionFormatSimple, version)
	}
	return nil
}

func (i Install) installVersion(ctx context.Context, version string) error {
	_, err := goinstall.GoInstall(ctx, "google.golang.org/protobuf", version, i.InstallDirectoryPath)
	if err != nil {
		return err
	}
	return nil
}
