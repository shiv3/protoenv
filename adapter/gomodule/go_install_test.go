package gomodule

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestGoInstall(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	type args struct {
		ctx         context.Context
		target      string
		tag         string
		installPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test go install protoenv",
			args: args{
				ctx:         context.Background(),
				target:      "github.com/shiv3/protoenv",
				tag:         "v0.0.0",
				installPath: "./",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := GoInstall(tt.args.ctx, tt.args.target, tt.args.tag, tt.args.installPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoInstall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			abs, _ := filepath.Abs("protoenv")
			if _, err := os.Stat(abs); err != nil {
				t.Errorf("couldn't install protoenv: %+v", err)
			}
			os.Remove(abs)
		})
	}
}
