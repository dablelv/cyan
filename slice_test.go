package util

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseIntSlice(t *testing.T) {
	intSlice := []int{1, 2, 3}
	assert.Equal(t, []int{3, 2, 1}, ReverseIntSlice(intSlice))
}

func TestSumSlice(t *testing.T) {
	f32Slice := []float32{1.1, 2.2, 3.3}
	f64 := SumSlice(f32Slice)
	assert.Equal(t, float32(6.6), float32(f64))
}

func TestMinSliceE(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"err int", args{""}, nil, true},
		{"zero int", args{[]int{}}, 0, false},
		{"min int", args{[]int{1, 2, 3}}, 1, false},
		{"zero uint", args{[]uint{}}, uint(0), false},
		{"min uint", args{[]uint{1, 2, 3}}, uint(1), false},
		{"zero float", args{[]float64{}}, 0.0, false},
		{"min float", args{[]float64{1.0, 2.0, 3.0}}, 1.0, false},
	}
	for _, tt := range tests {
		got, err := MinSliceE(tt.args.slice)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestMaxSliceE(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"err int", args{""}, nil, true},
		{"zero int", args{[]int{}}, 0, false},
		{"max int", args{[]int{1, 2, 3}}, 3, false},
		{"zero uint", args{[]uint{}}, uint(0), false},
		{"max uint", args{[]uint{1, 2, 3}}, uint(3), false},
		{"zero float", args{[]float64{}}, 0.0, false},
		{"max float", args{[]float64{1.0, 2.0, 3.0}}, 3.0, false},
	}
	for _, tt := range tests {
		got, err := MaxSliceE(tt.args.slice)
		assert.Equal(t, tt.wantErr, err != nil, tt.name)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestIsContains(t *testing.T) {
	type args struct {
		slice  interface{}
		target interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"string slice contain",
			args:args{
				slice:[]string{"foo","bar","baz"},
				target:"baz",
			},
			want:true,
		},
		{
			name:"int32 slice contain",
			args:args{
				slice:[]int32{1,2,3},
				target:int32(1),
			},
			want:true,
		},
		{
			name:"int32 slice not contain because type isn't equal",
			args:args{
				slice:[]int32{1,2,3},
				target:1,
			},
			want:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsContains(tt.args.slice, tt.args.target); got != tt.want {
				t.Errorf("IsContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueSliceE(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:"unique string slice",
			args:args{[]string{"foo","bar","bar","baz"}},
			want:[]string{"foo","bar","baz"},
			wantErr:false,
		},
		{
			name:"unique zero length string slice",
			args:args{[]string{}},
			want:[]string{},
			wantErr:false,
		},
		{
			name:"unique uint64 slice",
			args:args{[]uint64{1,2,2,3}},
			want:[]uint64{1,2,3},
			wantErr:false,
		},
		{
			name:"unique zero length uint64 slice",
			args:args{[]uint64{}},
			want:[]uint64{},
			wantErr:false,
		},
		{
			name:"input isn't a slice",
			args:args{nil},
			want:nil,
			wantErr:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UniqueSliceE(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("UniqueSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueSliceE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueStrSlice(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name:"unique string slice",
			args:args{[]string{"foo", "bar", "bar", "baz"}},
			want:[]string{"foo","bar","baz"},
		},
		{
			name:"unique zero length string slice",
			args:args{[]string{}},
			want:[]string{},
		},
		{
			name:"unique nil string slice",
			args:args{nil},
			want:[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueStrSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseSliceE(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:"reverse string slice",
			args:args{[]string{"foo","bar","baz"}},
			want:[]string{"baz", "bar", "foo"},
			wantErr:false,
		},
		{
			name:"reverse zero length string slice",
			args:args{[]string{}},
			want:[]string{},
			wantErr:false,
		},
		{
			name:"reverse uint64 slice",
			args:args{[]uint64{1,2,3}},
			want:[]uint64{3,2,1},
			wantErr:false,
		},
		{
			name:"reverse zero length uint64 slice",
			args:args{[]uint64{}},
			want:[]uint64{},
			wantErr:false,
		},
		{
			name:"input isn't a slice",
			args:args{nil},
			want:nil,
			wantErr:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReverseSliceE(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReverseSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSliceE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseStrSlice(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name:"reverse string slice",
			args:args{[]string{"foo","bar","baz"}},
			want:[]string{"baz", "bar", "foo"},
		},
		{
			name:"reverse zero length string slice",
			args:args{[]string{}},
			want:[]string{},
		},
		{
			name:"reverse nil string slice",
			args:args{[]string{}},
			want:[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseStrSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinSliceWithSepE(t *testing.T) {
	type args struct {
		slice interface{}
		sep   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:"join string slice",
			args:args{[]string{"foo","bar","baz"}, ","},
			want:"foo,bar,baz",
			wantErr:false,
		},
		{
			name:"join string array",
			args:args{[3]string{"foo","bar","baz"}, ","},
			want:"foo,bar,baz",
			wantErr:false,
		},
		{
			name:"join int slice",
			args:args{[]int{1,2,3}, ","},
			want:"1,2,3",
			wantErr:false,
		},
		{
			name:"join int array",
			args:args{[3]int{1,2,3}, ","},
			want:"1,2,3",
			wantErr:false,
		},
		{
			name:"join nil",
			args:args{nil, ","},
			want:"",
			wantErr:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JoinSliceWithSepE(tt.args.slice, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("JoinSliceWithSepE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JoinSliceWithSepE() got = %v, want %v", got, tt.want)
			}
		})
	}
}