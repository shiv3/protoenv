package installer

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/shiv3/protoenv/adapter/archiver"

	"github.com/shiv3/protoenv/adapter/github"
)

type InstallerGithubReleaseZip struct {
	InstallType   InstallType
	InstallConfig InstallConfig
	ArchMatcher   map[string]string
}

func NewInstallerGithubReleaseZip(installType InstallType, installConfig InstallConfig) *InstallerGithubReleaseZip {
	return &InstallerGithubReleaseZip{InstallType: installType, InstallConfig: installConfig}
}

func (i *InstallerGithubReleaseZip) Install(ctx context.Context) error {
	url, err := getReleaseAssetURL(ctx, i.InstallConfig.TargetUrl, i.InstallConfig.TargetVersion, i.ArchMatcher)
	if err != nil {
		return err
	}
	if i.InstallConfig.TargetBinaryName == "" {
		return fmt.Errorf("require target binary file name")
	}
	_, err = archiver.GetTargetFile(url, i.InstallConfig.TargetBinaryName, i.InstallConfig.TargetPath)
	return err
}

func getReleaseAssetURL(ctx context.Context, target, tag string, matcher map[string]string) (string, error) {
	urls, err := github.GetReleaseAssetsInstallURL(ctx, target, tag)
	if err != nil {
		return "", err
	}
	for _, url := range urls {
		filename := filepath.Base(url)
		archString := archMarch(runtime.GOOS, runtime.GOARCH, matcher)
		if strings.Contains(filename, "protoc") && strings.Contains(filename, archString) {
			return url, nil
		}
	}
	return "", fmt.Errorf("Unknown version")
}

func archMarch(os, arch string, matcher map[string]string) string {
	k := fmt.Sprintf("%s/%s", os, arch)
	if v, ok := matcher[k]; ok {
		return v
	}
	return ""
}
