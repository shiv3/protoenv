package github

import (
	"context"

	"github.com/google/go-github/v35/github"
)

func GetReleaseAssets(ctx context.Context, owner, repo string, tag string) ([]*github.ReleaseAsset, error) {
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetReleaseByTag(ctx, owner, repo, tag)
	if err != nil {
		return nil, err
	}
	return release.Assets, nil
}
