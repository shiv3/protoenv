package installer

import (
	"context"
)

type InstallType string

const (
	InstallTypeGoInstall        InstallType = "go_install"
	InstallTypeGitHubReleaseZip InstallType = "github_release_zip"
)

func NewInstaller(installType InstallType, installConfig InstallConfig) Installer {
	switch installType {
	case InstallTypeGoInstall:
		return NewInstallerGoInstall(InstallTypeGoInstall, installConfig)
	case InstallTypeGitHubReleaseZip:
		return NewInstallerGithubReleaseZip(InstallTypeGitHubReleaseZip, installConfig)
	}
	return nil
}

type Installer interface {
	Install(ctx context.Context, options ...InstallOption) error
	SetVersion(version string)
}
type InstallOption struct {
}

type InstallConfig struct {
	TargetUrl        string
	TargetVersion    string
	TargetPath       string
	TargetBinaryName string
}
