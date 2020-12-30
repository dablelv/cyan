package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerGTVer(t *testing.T) {
	res0, err0 := VerGTVer("1.0.5", "1.0.4")
	res1, err1 := VerGTVer("1.0.4", "1.0.5")
	res2, err2 := VerGTVer("2.0", "1.0")
	assert.NoError(t, err0)
	assert.Equal(t, true, res0)
	assert.NoError(t, err1)
	assert.Equal(t, false, res1)
	assert.Error(t, err2)
	assert.Equal(t, false, res2)
}

func TestVerLTVer(t *testing.T) {
	res0, err0 := VerLTVer("1.0.4", "1.0.5")
	res1, err1 := VerLTVer("1.0.5", "1.0.4")
	assert.NoError(t, err0)
	assert.Equal(t, true, res0)
	assert.NoError(t, err1)
	assert.Equal(t, false, res1)
}

func TestVerGEVer(t *testing.T) {
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
			name:    "equal",
			args:    args{"1.0.5", "1.0.5"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "lessthan",
			args:    args{"1.0.4", "1.0.5"},
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
		got, err := VerGEVer(tt.args.ver0, tt.args.ver1)
		assert.Equal(t, tt.want, got, tt.name)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
	}
}

func TestVerLEVer(t *testing.T) {
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
			name:    "lessthan",
			args:    args{"1.0.4", "1.0.5"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "equal",
			args:    args{"1.0.5", "1.0.5"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "greaterthan",
			args:    args{"1.0.5", "1.0.4"},
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
		got, err := VerLEVer(tt.args.ver0, tt.args.ver1)
		assert.Equal(t, tt.want, got, tt.name)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
	}
}

func TestVerGEVerMore(t *testing.T) {
	type args struct {
		ver0 string
		ver1 string
		sep  string
		num  int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "greaterthan",
			args:    args{"2-1-1-1", "2-1-1-0", "-", 4},
			want:    true,
			wantErr: false,
		},
		{
			name:    "equal",
			args:    args{"2-1-1-1", "2-1-1-1", "-", 4},
			want:    true,
			wantErr: false,
		},
		{
			name:    "lessthan",
			args:    args{"2-1-1-0", "2-1-1-1", "-", 4},
			want:    false,
			wantErr: false,
		},
		{
			name:    "argill",
			args:    args{"2-1-1", "2-1-0", "-", 4},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := VerGEVerMore(tt.args.ver0, tt.args.ver1, tt.args.sep, tt.args.num)
		assert.Equal(t, tt.want, got, tt.name)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
	}
}

func TestVerLEVerMore(t *testing.T) {
	type args struct {
		ver0 string
		ver1 string
		sep  string
		num  int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "lessthan",
			args:    args{"2-1-1-0", "2-1-1-1", "-", 4},
			want:    true,
			wantErr: false,
		},
		{
			name:    "equal",
			args:    args{"2-1-1-1", "2-1-1-1", "-", 4},
			want:    true,
			wantErr: false,
		},
		{
			name:    "greaterthan",
			args:    args{"2-1-1-1", "2-1-1-0", "-", 4},
			want:    false,
			wantErr: false,
		},
		{
			name:    "argill",
			args:    args{"2-1-1", "2-1-0", "-", 4},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := VerLEVerMore(tt.args.ver0, tt.args.ver1, tt.args.sep, tt.args.num)
		assert.Equal(t, tt.want, got, tt.name)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
	}
}
