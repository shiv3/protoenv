package installer

import (
	"context"
)

type InstallType string

const (
	InstallTypeGoInstall        InstallType = "go_install"
	InstallTypeGitHubReleaseZip InstallType = "github_release_zip"
)

func NewInstallers(installConfig InstallConfig) map[InstallType]Installer {
	return map[InstallType]Installer{
		InstallTypeGoInstall:        NewInstallerGoInstall(InstallTypeGoInstall, installConfig),
		InstallTypeGitHubReleaseZip: NewInstallerGithubReleaseZip(InstallTypeGitHubReleaseZip, installConfig),
	}
}

type InstallOption interface {
}

type Installer interface {
	Install(ctx context.Context) error
}

type InstallConfig struct {
	TargetUrl        string
	TargetVersion    string
	TargetPath       string
	TargetBinaryName string
}
