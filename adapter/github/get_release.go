package github

import (
	"context"

	"github.com/google/go-github/v35/github"
)

func GetReleaseAssetsInstallURL(ctx context.Context, target string, tag string) ([]string, error) {
	owner, repo, err := ParseUrl(target)
	if err != nil {
		return nil, err
	}
	assets, err := GetGithubReleaseAssets(ctx, owner, repo, tag)
	if err != nil {
		return nil, err
	}
	res := make([]string, 0, len(assets))
	for _, asset := range assets {
		url := asset.GetBrowserDownloadURL()
		res = append(res, url)
	}
	return res, nil
}

func GetGithubReleaseAssets(ctx context.Context, owner, repo string, tag string) ([]*github.ReleaseAsset, error) {
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetReleaseByTag(ctx, owner, repo, tag)
	if err != nil {
		return nil, err
	}
	return release.Assets, nil
}
