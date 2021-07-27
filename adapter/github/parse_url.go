package github

import (
	"fmt"
	"net/url"
	"strings"
)

func ParseUrl(target string) (string, string, error) {
	u, err := url.Parse(target)
	if err != nil {
		return "", "", err
	}
	if u.Scheme == "" {
		u, err = url.Parse("https://" + target)
		if err != nil {
			return "", "", err
		}
	}

	if u.Hostname() != "github.com" {
		return "", "", fmt.Errorf("hostname: %s is not github.com", u.Hostname())
	}
	v := strings.Split(u.Path, "/")
	owner, repo := v[1], v[2]
	return owner, repo, nil
}
