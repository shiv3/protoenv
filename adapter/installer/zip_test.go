package installer

import "testing"

func Test_downloadZip(t *testing.T) {
	type args struct {
		dir string
		url string
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
				dir: ".",
				url: "http://www.golang-book.com/public/pdf/gobook.pdf",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := downloadZip(tt.args.dir, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("downloadZip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("downloadZip() got = %v, want %v", got, tt.want)
			}
		})
	}
}
