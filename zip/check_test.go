package zip

import "testing"

func TestIsZipFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "file isn't zip",
			args:    args{"check.go"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "file not foud",
			args:    args{"check.g"},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsZipFile(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsZipFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsZipFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
