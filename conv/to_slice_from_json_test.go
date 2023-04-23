package conv

import (
	"reflect"
	"testing"
)

func TestJSONToSlice(t *testing.T) {
	type args struct {
		data []byte
	}

	// int slice test.
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"int to slice",
			args{[]byte("[1,2,3]")},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JSONToSlice[[]int](tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONToSlice() = %v, want %v", got, tt.want)
			}
		})
	}

	// float32 slice test.
	tests1 := []struct {
		name string
		args args
		want []float32
	}{
		{
			"float32 to slice",
			args{[]byte("[1.1,2.2,3.3]")},
			[]float32{1.1, 2.2, 3.3},
		},
	}
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got := JSONToSlice[[]float32](tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONToSlice() = %v, want %v", got, tt.want)
			}
		})
	}

	// string slice test.
	tests2 := []struct {
		name string
		args args
		want []string
	}{
		{
			"string to slice",
			args{[]byte(`["foo","bar","baz"]`)},
			[]string{"foo", "bar", "baz"},
		},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got := JSONToSlice[[]string](tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
