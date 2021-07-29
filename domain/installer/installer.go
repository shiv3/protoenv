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
		return NewInstallerGithubReleaseZip(InstallTypeGitHubReleaseZip, installConfig, map[string]string{
			"darwin/386":   "osx-x86_64",
			"darwin/amd64": "osx-x86_64",
			"linux/386":    "linux-x86_32",
			"linux/amd64":  "linux-x86_64",
		})
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
