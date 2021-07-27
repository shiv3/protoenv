package installer

import (
	"context"

	"github.com/shiv3/protoenv/adapter/goinstall"
)

type InstallerGoInstall struct {
	InstallType   InstallType
	InstallConfig InstallConfig
}

func NewInstallerGoInstall(installType InstallType, installConfig InstallConfig) *InstallerGoInstall {
	return &InstallerGoInstall{InstallType: installType, InstallConfig: installConfig}
}

func (i *InstallerGoInstall) Install(ctx context.Context) error {
	return goinstall.GoInstall(ctx, i.InstallConfig.TargetUrl, i.InstallConfig.TargetVersion, i.InstallConfig.TargetPath)
}
