package conv

import (
	"reflect"
	"testing"
)

func TestSplitStrToIntSliceE(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "to int slice",
			args: args{
				s:   "1,2,3",
				sep: ",",
			},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				s:   "1,2,foo",
				sep: ",",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitStrToSliceE[int](tt.args.s, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitStrToSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitStrToSliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitStrToUintSliceE(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name    string
		args    args
		want    []uint
		wantErr bool
	}{
		{
			name: "to uint slice",
			args: args{
				s:   "1,2,3",
				sep: ",",
			},
			want:    []uint{1, 2, 3},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				s:   "1,2,foo",
				sep: ",",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitStrToSliceE[uint](tt.args.s, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitStrToSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitStrToSliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitStrToFloat64SliceE(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name: "to float64 slice",
			args: args{
				s:   "1,2,3",
				sep: ",",
			},
			want:    []float64{1, 2, 3},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				s:   "1,2,foo",
				sep: ",",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitStrToSliceE[float64](tt.args.s, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitStrToSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitStrToSliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitStrToBoolSliceE(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name    string
		args    args
		want    []bool
		wantErr bool
	}{
		{
			name: "to bool slice",
			args: args{
				s:   "1,0,1",
				sep: ",",
			},
			want:    []bool{true, false, true},
			wantErr: false,
		},
		{
			name: "to bool slice",
			args: args{
				s:   "true,false,true",
				sep: ",",
			},
			want:    []bool{true, false, true},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				s:   "1,0,2",
				sep: ",",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitStrToSliceE[bool](tt.args.s, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitStrToSliceE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitStrToSliceE() = %v, want %v", got, tt.want)
			}
		})
	}
}
