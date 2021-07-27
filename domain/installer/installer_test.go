package installer

//
//import (
//	"context"
//	"os"
//	"path/filepath"
//	"testing"
//)
//
//func TestInstallerImpl_Install(t *testing.T) {
//	if testing.Short() {
//		t.Skip()
//	}
//
//	testInstallDir, _ := filepath.Abs("test_out")
//	defer func() {
//		if err := os.RemoveAll(testInstallDir); err != nil {
//			t.Fatal(err)
//		}
//	}()
//
//	type fields struct {
//		InstallConfig InstallConfig
//	}
//	type args struct {
//		ctx    context.Context
//		option InstallOption
//	}
//	tests := []struct {
//		name     string
//		fields   fields
//		args     args
//		wantFile string
//		wantErr  bool
//	}{
//		{
//			name: "go install",
//			fields: fields{
//				InstallConfig: InstallConfig{
//					TargetUrl:     "google.golang.org/protobuf/cmd/protoc-gen-go",
//					TargetVersion: "v1.27.1",
//					TargetPath:    testInstallDir,
//				},
//			},
//			args: args{
//				ctx: nil,
//				option: InstallOption{
//					InstallType: InstallTypeGoInstall,
//					ArchiveType: ArchiveTypeRaw,
//				},
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			i := &InstallerImpl{
//				InstallConfig: tt.fields.InstallConfig,
//			}
//			if err := i.Install(tt.args.ctx, tt.args.option); (err != nil) != tt.wantErr {
//				t.Errorf("Install() error = %v, wantErr %v", err, tt.wantErr)
//			}
//			if _, err := os.Stat(filepath.Join(testInstallDir, tt.wantFile)); err != nil {
//				t.Errorf("couldn't install %s: %+v", tt.wantFile, err)
//			}
//		})
//	}
//}
