package cmp

import (
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestVerGTVer(t *testing.T) {
	type args struct {
		v0 string
		v1 string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"err", args{"1.1.5", "1.0.20.1"}, false, true},
		{"gt", args{"1.1.5", "1.0.20"}, true, false},
		{"eq", args{"1.1.5", "1.1.5"}, false, false},
		{"lt", args{"1.1.0", "1.1.5"}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerGTVer(tt.args.v0, tt.args.v1)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerGTVer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerGTVer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerGEVer(t *testing.T) {
	assert := internal.NewAssert(t, "TestVerGEVer")
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
		{"arg ill", args{"1.0.5.5", "2.0.4"}, false, true},
		{"greater than", args{"1.0.5", "1.0.4"}, true, false},
		{"equal", args{"1.0.5", "1.0.5"}, true, false},
		{"less than", args{"1.0.4", "1.0.5"}, false, false},
	}
	for _, tt := range tests {
		got, err := VerGEVer(tt.args.ver0, tt.args.ver1)
		assert.Equal(tt.want, got)
		assert.Equal(tt.wantErr, err != nil)
	}
}

func TestVerLTVer(t *testing.T) {
	type args struct {
		v0 string
		v1 string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"err", args{"1.1.5", "1.0.20.1"}, false, true},
		{"gt", args{"1.1.5", "1.0.20"}, false, false},
		{"eq", args{"1.1.5", "1.1.5"}, false, false},
		{"lt", args{"1.1.0", "1.1.5"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerLTVer(tt.args.v0, tt.args.v1)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerLTVer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerLTVer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerLEVer(t *testing.T) {
	assert := internal.NewAssert(t, "TestVerLEVer")
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
		{"err", args{"1.0.5.5", "2.0.4"}, false, true},
		{"less than", args{"1.0.4", "1.0.5"}, true, false},
		{"equal", args{"1.0.5", "1.0.5"}, true, false},
		{"greater than", args{"1.0.5", "1.0.4"}, false, false},
	}
	for _, tt := range tests {
		got, err := VerLEVer(tt.args.ver0, tt.args.ver1)
		assert.Equal(tt.want, got)
		assert.Equal(tt.wantErr, err != nil)
	}
}
