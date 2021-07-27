package github

import (
	"context"

	"github.com/google/go-github/v35/github"
)

func GetReleaseVersions(ctx context.Context, target string) ([]string, error) {
	owner, repo, err := ParseUrl(target)
	if err != nil {
		return nil, err
	}
	return GetGithubReleaseVersions(ctx, owner, repo)
}

func GetGithubReleaseVersions(ctx context.Context, owner, repo string) ([]string, error) {
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
