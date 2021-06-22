package github

import (
	"context"

	"github.com/google/go-github/v35/github"
)

func GetVersions(ctx context.Context, owner, repo string) ([]string, error) {
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(ctx, owner, repo, nil)
	if err != nil {
		return nil, err
	}
	versions := make([]string, 0, len(releases))
	for _, release := range releases {
		versions = append(versions, release.GetTagName())
	}
	return versions, nil
}
