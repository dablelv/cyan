package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerLTVer(t *testing.T) {
	res0, err0 := VerLTVer("1.0.4", "1.0.5")
	res1, err1 := VerLTVer("1.0.5", "1.0.4")
	assert.NoError(t, err0)
	assert.Equal(t, true, res0)
	assert.NoError(t, err1)
	assert.Equal(t, false, res1)
}

func TestVerGTVer(t *testing.T) {
	type args struct {
		ver0 string
		ver1 string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "greaterthan",
			args:    args{"1.0.5", "1.0.4"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "lessthan",
			args:    args{"1.0.5", "1.0.4"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "equal",
			args:    args{"1.0.5", "1.0.5"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "lessthan",
			args:    args{"1.0.5", "2.0.4"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "argill",
			args:    args{"1.0.5.5", "2.0.4"},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerGTVer(tt.args.ver0, tt.args.ver1)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerGTVer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerGTVer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
