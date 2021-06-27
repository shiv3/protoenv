package github

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
)

var (
	owner = "protocolbuffers"
	repo  = "protobuf"
)

func GetProtobufVersions(ctx context.Context) ([]string, error) {
	return GetVersions(ctx, owner, repo)
}

func GetProtobufGetReleaseAssetURL(ctx context.Context, tag string,os ,arch string) (string, error) {
	assets, err := GetReleaseAssets(ctx, owner, repo, tag)
	if err != nil {
		return "", err
	}
	for _, asset := range assets {
		url := asset.GetBrowserDownloadURL()
		filename := filepath.Base(url)
		archString :=  Get(os,arch)
		if strings.Contains(filename,"protoc") && strings.Contains(filename,archString) {
			return url, nil
		}
	}
	return "", fmt.Errorf("Unknown version")
}

func Get(os,arch string) string {
	switch fmt.Sprintf("%s/%s",os,arch) {
	case "darwin/386":
		return "osx-x86_64"
	case "darwin/amd64":
		return "osx-x86_64"
	case "linux/386":
		return "linux-x86_32"
	case "linux/amd64":
		return "linux-x86_64"
	default:
		return "Unknown"
	}
}
