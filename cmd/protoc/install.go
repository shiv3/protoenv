package protoc

import (
	"context"
	"errors"
	"fmt"
	"runtime"

	"github.com/shiv3/protoenv/adapter/github"

	"github.com/spf13/cobra"
)

var (
	ShowVersionFormatSimple = "%s\n"
)

type Install struct {
	command *cobra.Command
}

type InstallOptions struct {
	ShowVersionList bool
}

func NewInstall(options InstallOptions) *Install {
	cmd := newInstallCmd(options)
	return &Install{
		command: &cobra.Command{
			Use:   "install",
			Short: "install specified version:",
			Long:  `install specified version:`,
			RunE:  cmd.run,
		},
	}
}

func (i Install) GetCommand() *cobra.Command {
	return i.command
}

type installCmd struct {
	installOptions InstallOptions
}

func newInstallCmd(installOptions InstallOptions) installCmd {
	return installCmd{installOptions: installOptions}
}

func (c *installCmd) run(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	if c.installOptions.ShowVersionList {
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

func (c *installCmd) showVersion(ctx context.Context) error {
	versions, err := github.GetProtobufVersions(ctx)
	if err != nil {
		return err
	}
	for _, version := range versions {
		fmt.Printf(ShowVersionFormatSimple, version)
	}
	return nil
}

func (c *installCmd) installVersion(ctx context.Context, version string) error {
	url, err := github.GetProtobufGetReleaseAssetURL(ctx, version, runtime.GOARCH)
	if err != nil {
		return err
	}

	return nil
}
