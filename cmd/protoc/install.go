package protoc

import (
	"context"
	"errors"
	"fmt"
	"github.com/shiv3/protoenv/adapter/installer"
	"os"
	"path/filepath"
	"runtime"

	"github.com/shiv3/protoenv/adapter/github"

	"github.com/spf13/cobra"
)

var (
	ShowVersionFormatSimple = "%s\n"
)

type Install struct {
	InstallDirectoryPath string
}

func NewInstall(parentCmd *cobra.Command, installDirectoryPath string) Install {
	install := Install{
		InstallDirectoryPath: installDirectoryPath,
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

func (c Install) RunE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	//c.installOptions.ShowVersionList {
	if list, err := cmd.PersistentFlags().GetBool("list"); err == nil && list {
		if err := c.showVersion(ctx); err != nil {
			return err
		}
		return nil
	}

	if len(args) >= 1 {
		err := c.installVersion(ctx, args[0])
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("requires a installing version or some flags")
}

func (c Install) showVersion(ctx context.Context) error {
	versions, err := github.GetProtobufVersions(ctx)
	if err != nil {
		return err
	}
	for _, version := range versions {
		fmt.Printf(ShowVersionFormatSimple, version)
	}
	return nil
}

func (c Install) installVersion(ctx context.Context, version string) error {
	url, err := github.GetProtobufGetReleaseAssetURL(ctx, version, runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return err
	}
	targetDirPath := filepath.Join(c.InstallDirectoryPath, "versions", version)
	err = os.MkdirAll(targetDirPath, os.ModePerm)
	if err != nil {
		return err
	}
	filePath, err := installer.GetTargetFile(url, "protoc", targetDirPath)
	if err != nil {
		return err
	}
	fmt.Printf("installed protoc %s\n", filePath)
	if err := setGlobalVersion(getGlobalVersionFilePath(c.InstallDirectoryPath), version); err != nil {
		return err
	}
	return nil
}
