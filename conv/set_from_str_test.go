package conv

import (
	"reflect"
	"testing"
)

func TestSplitStrToSet(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{"nil string", args{"", ","}, nil},
		{"nil separator", args{"foo", ""}, map[string]struct{}{"f": {}, "o": {}}},
		{"string to set", args{"f,o,o", ","}, map[string]struct{}{"f": {}, "o": {}}},
		{"string ended with sep", args{"f,o,o,", ","}, map[string]struct{}{"f": {}, "o": {}, "": {}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitStrToSet(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitStrToSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
