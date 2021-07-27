package github

import (
	"context"
	"testing"
)

func TestGetVersions(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	type args struct {
		ctx    context.Context
		target string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test GetReleaseVersions",
			args: args{
				ctx:    context.Background(),
				target: "github.com/shiv3/protoenv",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetReleaseVersions(tt.args.ctx, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReleaseVersions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
