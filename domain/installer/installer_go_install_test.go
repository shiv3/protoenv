package installer

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestInstallerGoInstall_Install(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	testInstallDir, _ := filepath.Abs("test_out")
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
					TargetUrl:        "google.golang.org/protobuf/cmd/protoc-gen-go",
					TargetVersion:    "v1.27.1",
					TargetPath:       testInstallDir,
					TargetBinaryName: "protoc-gen-go",
				},
			},
			args: args{
				ctx: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InstallerGoInstall{
				InstallType:   tt.fields.InstallType,
				InstallConfig: tt.fields.InstallConfig,
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
