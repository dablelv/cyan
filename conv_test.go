package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap2StrSliceE(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		want1   []string
		wantErr bool
	}{
		{
			name:"map[string]string",
			args:args{map[string]string{"CN":"China", "HK":"Hong Kong","AU":"Australia"}},
			want:[]string{"CN","HK","AU"},
			want1:[]string{"China","Hong Kong","Australia"},
			wantErr: false,
		},
		{
			name:"map[string]int",
			args:args{map[string]int{"CN":86, "HK":852,"AU":61}},
			want:[]string{"CN","HK","AU"},
			want1:[]string{"86","852","61"},
			wantErr: false,
		},
		{
			name:"map[int]string",
			args:args{map[int]string{86:"China", 852:"Hong Kong",61:"Australia"}},
			want:[]string{"86","852","61"},
			want1:[]string{"China","Hong Kong","Australia"},
			wantErr: false,
		},
		{
			name:"map[string]string",
			args:args{map[string]string{"CN":"China", "HK":"Hong Kong","AU":"Australia"}},
			want:[]string{"CN","HK","AU"},
			want1:[]string{"China","Hong Kong","Australia"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Map2StrSliceE(tt.args.i)
			assert.Equal(t, tt.wantErr, err != nil, tt.name)
			assert.Equal(t, len(tt.want), len(got), tt.name)
			assert.Equal(t, len(tt.want1), len(got1), tt.name)
		})
	}
}
