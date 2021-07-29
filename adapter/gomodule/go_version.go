package gomodule

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
)

type jsonOut struct {
	Version   string
	Versions  []string
	GoVersion string
}

func GoGetVersions(ctx context.Context, target string) ([]string, error) {
	cmdString := []string{
		"go", "list", "-u", "-m", "-versions", "-json", target,
	}
	cmd := exec.Command(cmdString[0], cmdString[1:]...)
	//fmt.Printf("running...\n%s \n", strings.Join(cmdString, " "))
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			fmt.Printf("error:\n%s", ee.Stderr)
		}
	}
	var v jsonOut
	err = json.Unmarshal(out, &v)
	return v.Versions, err
}
