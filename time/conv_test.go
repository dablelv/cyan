package time

import "testing"

// Note that not suitable for adding unittests, because Local timezone may be different.
func TestDateTime2UTS(t *testing.T) {
	type args struct {
		dt string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateTime2UTS(tt.args.dt); got != tt.want {
				t.Errorf("DateTime2UTS() = %v, want %v", got, tt.want)
			}
		})
	}
}
