package github

import (
	"context"
	"testing"
)

func TestGetProtobufGetReleaseAssetURL(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	type args struct {
		ctx  context.Context
		tag  string
		arch string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:  context.TODO(),
				tag:  "v3.17.3",
				arch: "linux/386",
			},
			want:    "https://github.com/protocolbuffers/protobuf/releases/download/v3.17.3/protoc-3.17.3-linux-x86_32.zip",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetProtobufGetReleaseAssetURL(tt.args.ctx, tt.args.tag, tt.args.arch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProtobufGetReleaseAssetURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetProtobufGetReleaseAssetURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
