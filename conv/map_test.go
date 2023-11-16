package conv

import (
	"reflect"
	"testing"
)

func TestStruct2Map(t *testing.T) {
	type args struct {
		obj any
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "convert success",
			args: args{
				struct {
					I int
					S string
				}{I: 1, S: "a"}},
			want: map[string]any{"I": 1, "S": "a"},
		},
		{
			name: "convert struct contains unexported fields success",
			args: args{
				struct {
					I int
					S string
					s string
				}{I: 1, S: "a", s: "a"}},
			want: map[string]any{"I": 1, "S": "a"},
		},
		{
			name: "convert empty struct success",
			args: args{struct{}{}},
			want: map[string]any{},
		},
		{
			name: "not struct failed",
			args: args{1},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructToMap(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Struct2Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStruct2MapStr(t *testing.T) {
	type args struct {
		obj any
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "convert success",
			args: args{
				struct {
					I int
					S string
				}{I: 1, S: "a"}},
			want: map[string]string{"I": "1", "S": "a"},
		},
		{
			name: "convert struct contains unexported fields success",
			args: args{
				struct {
					I int
					S string
					s string
				}{I: 1, S: "a", s: "a"}},
			want: map[string]string{"I": "1", "S": "a"},
		},
		{
			name: "convert empty struct success",
			args: args{struct{}{}},
			want: map[string]string{},
		},
		{
			name: "not struct failed",
			args: args{1},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructToMapStr(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Struct2MapString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapStr(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "map[string]string",
			args: args{map[string]string{"foo": "foo", "bar": "bar", "baz": "baz"}},
			want: map[string]string{"foo": "foo", "bar": "bar", "baz": "baz"},
		},
		{
			name: "map[string]any",
			args: args{map[string]any{"foo": 1, "bar": 2, "baz": 3}},
			want: map[string]string{"foo": "1", "bar": "2", "baz": "3"},
		},
		{
			name: "map[string]any failed",
			args: args{map[string]any{"foo": 1, "bar": 2, "baz": []int{}}},
			want: nil,
		},
		{
			name: "map[any]string",
			args: args{map[any]string{1: "foo", 2: "bar", 3: "baz"}},
			want: map[string]string{"1": "foo", "2": "bar", "3": "baz"},
		},
		{
			name: "map[any]string failed",
			args: args{map[any]string{uintptr(1): "foo", 2: "bar", 3: "baz"}},
			want: nil,
		},
		{
			name: "map[any]any",
			args: args{map[any]any{1: 1, 2: 2, 3: 3}},
			want: map[string]string{"1": "1", "2": "2", "3": "3"},
		},
		{
			name: "map[any]any key failed",
			args: args{map[any]any{uintptr(1): 1, 2: 2, 3: 3}},
			want: nil,
		},
		{
			name: "map[any]any value failed",
			args: args{map[any]any{1: uintptr(1), 2: 2, 3: 3}},
			want: nil,
		},
		{
			name: "json",
			args: args{`{"foo":"foo","bar":"bar","baz":"baz"}`},
			want: map[string]string{"foo": "foo", "bar": "bar", "baz": "baz"},
		},
		{
			name: "string to map failed",
			args: args{`{"foo":1`},
			want: nil,
		},
		{
			name: "int to map failed",
			args: args{1},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMapStr(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapStrStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
