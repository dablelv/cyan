package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDaysBtwTs(t *testing.T) {
	type args struct {
		ts0 int64
		ts1 int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1day",
			args: args{1605300000, 1605386401},
			want: 1,
		},
		{
			name: "1day in ireverse ts order",
			args: args{1605386401, 1605300000},
			want: 1,
		},
		{
			name: "11days in ireverse ts order",
			args: args{1606250400, 1605300000},
			want: 11,
		},
	}
	for _, tt := range tests {
		got := GetDaysBtwTs(tt.args.ts0, tt.args.ts1)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestGetHoursBtwTs(t *testing.T) {
	type args struct {
		ts0 int64
		ts1 int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1hour",
			args: args{1605340000, 1605343712},
			want: 1,
		},
		{
			name: "1hour in ireverse ts order",
			args: args{1605343712, 1605340000},
			want: 1,
		},
		{
			name: "11hours in ireverse ts order",
			args: args{1605343712, 1605304112},
			want: 11,
		},
	}
	for _, tt := range tests {
		got := GetHoursBtwTs(tt.args.ts0, tt.args.ts1)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestGetMinutesBtwTs(t *testing.T) {
	type args struct {
		ts0 int64
		ts1 int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1minute",
			args: args{1605340000, 1605340061},
			want: 1,
		},
		{
			name: "1minute in ireverse ts order",
			args: args{1605340061, 1605340000},
			want: 1,
		},
		{
			name: "11minutes in ireverse ts order",
			args: args{1605340661, 1605340000},
			want: 11,
		},
	}
	for _, tt := range tests {
		got := GetMinutesBtwTs(tt.args.ts0, tt.args.ts1)
		assert.Equal(t, tt.want, got, tt.name)
	}
}
