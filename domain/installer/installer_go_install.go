package installer

import (
	"context"

	"github.com/shiv3/protoenv/adapter/gomodule"
)

type InstallerGoInstall struct {
	InstallType   InstallType
	InstallConfig InstallConfig
}

func NewInstallerGoInstall(installType InstallType, installConfig InstallConfig) *InstallerGoInstall {
	return &InstallerGoInstall{InstallType: installType, InstallConfig: installConfig}
}

func (i *InstallerGoInstall) Install(ctx context.Context, options ...InstallOption) error {
	return gomodule.GoInstall(ctx, i.InstallConfig.TargetUrl, i.InstallConfig.TargetVersion, i.InstallConfig.TargetPath)
}

func (i *InstallerGoInstall) SetVersion(version string) {
	i.InstallConfig.TargetVersion = version
}
