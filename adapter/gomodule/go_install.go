package gomodule

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func GoInstall(ctx context.Context, target, tag, installPath string) error {
	cmdString := []string{
		"go", "install", getInstallPath(target, tag),
	}
	cmd := exec.Command(cmdString[0], cmdString[1:]...)
	if !path.IsAbs(installPath) {
		installPath, _ = filepath.Abs(installPath)
	}
	env := fmt.Sprintf("GOBIN=%s", installPath)
	cmd.Env = append(os.Environ(), env)
	fmt.Printf("running...\n%s %s \n", env, strings.Join(cmdString, " "))
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			fmt.Printf("error:\n%s", ee.Stderr)
		}
	}
	fmt.Printf("%s", out)
	return err
}

// output install path
func GoInstallTarget(ctx context.Context, url, owner, repo string, tag string) (string, error) {
	exec.Command("go install", getInstallPath(getTargetPath(url, owner, repo), tag))
	return "", nil
}

func getInstallPath(target, tag string) string {
	return fmt.Sprintf("%s@%s", target, tag)
}

func getTargetPath(url, owner, repo string) string {
	return path.Join(url, owner, repo)
}
