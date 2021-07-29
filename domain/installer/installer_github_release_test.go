package installer

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestInstallerGithubReleaseZip_Install(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	testInstallDir, _ := filepath.Abs("test_out")
	os.Mkdir(testInstallDir, os.ModePerm)
	defer func() {
		if err := os.RemoveAll(testInstallDir); err != nil {
			t.Fatal(err)
		}
	}()

	type fields struct {
		InstallType   InstallType
		InstallConfig InstallConfig
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				InstallType: "",
				InstallConfig: InstallConfig{
					TargetUrl:        "github.com/protocolbuffers/protobuf",
					TargetVersion:    "v3.17.3",
					TargetPath:       testInstallDir,
					TargetBinaryName: "protoc",
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InstallerGithubReleaseZip{
				InstallType:   tt.fields.InstallType,
				InstallConfig: tt.fields.InstallConfig,
				ArchMatcher: map[string]string{
					"darwin/386":   "osx-x86_64",
					"darwin/amd64": "osx-x86_64",
					"linux/386":    "linux-x86_32",
					"linux/amd64":  "linux-x86_64",
				},
			}
			if err := i.Install(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Install() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, err := os.Stat(filepath.Join(testInstallDir, tt.fields.InstallConfig.TargetBinaryName)); err != nil {
				t.Errorf("couldn't install %+v", err)
			}
		})
	}
}
