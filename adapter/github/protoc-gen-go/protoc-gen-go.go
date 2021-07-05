package protoc_gen_go

import (
	"context"
	"fmt"
	"github.com/shiv3/protoenv/adapter/github"
	"path/filepath"
	"strings"
)

var (
	owner = "protocolbuffers"
	repo  = "protobuf-go"
)

func GetprotocGenGoRepoGoVersions(ctx context.Context) ([]string, error) {
	return github.GetVersions(ctx, owner, repo)
}

func GetprotocGenGoRepoGoGetReleaseAssetURL(ctx context.Context, tag string, os, arch string) (string, error) {
	assets, err := github.GetReleaseAssets(ctx, owner, repo, tag)
	if err != nil {
		return "", err
	}
	for _, asset := range assets {
		url := asset.GetBrowserDownloadURL()
		filename := filepath.Base(url)
		archString := Get(os, arch)
		if strings.Contains(filename, "protoc") && strings.Contains(filename, archString) {
			return url, nil
		}
	}
	return "", fmt.Errorf("Unknown version")
}

func Get(os, arch string) string {
	switch fmt.Sprintf("%s/%s", os, arch) {
	case "darwin/amd64":
		return "darwin.amd64"
	case "linux/386":
		return "linux.386"
	case "linux/amd64":
		return "linux.amd64"
	default:
		return "Unknown"
	}
}
