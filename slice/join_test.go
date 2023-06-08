package slice

import (
	"testing"

	"github.com/dablelv/go-huge-util/internal"
)

func TestJoin(t *testing.T) {
	assert := internal.NewAssert(t, "TestJoin")

	got := Join([]int{1})
	assert.Equal("1", got)
}

func TestJoinE(t *testing.T) {
	assert := internal.NewAssert(t, "TestJoinE")

	got, err := JoinE([]int{1})
	assert.IsNil(err)
	assert.Equal("1", got)
}

func TestJoinWithSep(t *testing.T) {
	assert := internal.NewAssert(t, "TestJoinWithSep")

	got := JoinWithSep([]int{1}, "")
	assert.Equal("1", got)
}

func TestJoinWithSepE(t *testing.T) {
	type args struct {
		slice any
		sep   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "join string slice",
			args:    args{[]string{"foo", "bar", "baz"}, ","},
			want:    "foo,bar,baz",
			wantErr: false,
		},
		{
			name:    "join string array",
			args:    args{[3]string{"foo", "bar", "baz"}, ","},
			want:    "foo,bar,baz",
			wantErr: false,
		},
		{
			name:    "join int slice",
			args:    args{[]int{1, 2, 3}, ","},
			want:    "1,2,3",
			wantErr: false,
		},
		{
			name:    "join int array",
			args:    args{[3]int{1, 2, 3}, ","},
			want:    "1,2,3",
			wantErr: false,
		},
		{
			name:    "join struct slice",
			args:    args{[]struct{}{{}}, ","},
			want:    "",
			wantErr: true,
		},
		{
			name:    "join nil",
			args:    args{nil, ","},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JoinWithSepE(tt.args.slice, tt.args.sep)
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
