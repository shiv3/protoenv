package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/shiv3/protoenv/adapter/gomodule"

	"github.com/shiv3/protoenv/domain/installer"

	"github.com/spf13/cobra"
)

type Install struct {
	InstallDirectoryPath    string
	ShowVersionFormatSimple string
	TargetUrl               string
	installer               installer.Installer
}

func NewInstall(parentCmd *cobra.Command, installDirectoryPath string, ShowVersionFormatSimple string, TargetBinaryFileName string) Install {
	install := Install{
		InstallDirectoryPath:    installDirectoryPath,
		ShowVersionFormatSimple: ShowVersionFormatSimple,
		TargetUrl:               "google.golang.org/protobuf/cmd/protoc-gen-go",
		installer: installer.NewInstaller(installer.InstallTypeGoInstall, installer.InstallConfig{
			TargetUrl:        "google.golang.org/protobuf/cmd/protoc-gen-go",
			TargetVersion:    "",
			TargetPath:       installDirectoryPath,
			TargetBinaryName: TargetBinaryFileName,
		}),
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
	versions, err := gomodule.GoGetVersions(ctx, i.TargetUrl)
	if err != nil {
		return err
	}
	for _, version := range versions {
		fmt.Printf(i.ShowVersionFormatSimple, version)
	}
	return nil
}

func (i Install) installVersion(ctx context.Context, version string) error {
	i.installer.SetVersion(version)
	return i.installer.Install(ctx)
}
