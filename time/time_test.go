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
		{"1 day", args{1605300000, 1605386401}, 1},
		{"1 day in reverse ts order", args{1605386401, 1605300000}, 1},
		{"11 days in reverse ts order", args{1606250400, 1605300000}, 11},
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
		{"1 hour", args{1605340000, 1605343712}, 1},
		{"1 hour in reverse ts order", args{1605343712, 1605340000}, 1},
		{"11 hours in reverse ts order", args{1605343712, 1605304112}, 11},
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
		{"1 minute", args{1605340000, 1605340061}, 1},
		{"1 minute in reverse ts order", args{1605340061, 1605340000}, 1},
		{"11 minutes in reverse ts order", args{1605340661, 1605340000}, 11},
	}
	for _, tt := range tests {
		got := GetMinutesBtwTs(tt.args.ts0, tt.args.ts1)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestIsSameYear(t *testing.T) {
	type args struct {
		uts1 int64
		uts2 int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"same year", args{1577808000, 1609430399}, true},      // "2020-01-01 00:00:00" & "2020-12-31 23:59:59"
		{"not same year", args{1577808000, 1609430400}, false}, // "2020-01-01 00:00:00" & "2021-01-01 00:00:00"
	}
	for _, tt := range tests {
		got := IsSameYear(tt.args.uts1, tt.args.uts2)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestIsSameMonth(t *testing.T) {
	type args struct {
		uts1 int64
		uts2 int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"same month", args{1577808000, 1580486399}, true},      // "2020-01-01 00:00:00" & "2020-01-31 23:59:59"
		{"not same month", args{1577808000, 1580486400}, false}, // "2020-01-01 00:00:00" & "2020-02-01 00:00:00"
	}
	for _, tt := range tests {
		got := IsSameMonth(tt.args.uts1, tt.args.uts2)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestIsSameDay(t *testing.T) {
	type args struct {
		uts1 int64
		uts2 int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"same day", args{1577808000, 1577894399}, true},      // "2020-01-01 00:00:00" & "2020-01-01 23:59:59"
		{"not same day", args{1577808000, 1577894400}, false}, // "2020-01-01 00:00:00" & "2020-01-02 00:00:00"
	}
	for _, tt := range tests {
		got := IsSameDay(tt.args.uts1, tt.args.uts2)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestIsSameHour(t *testing.T) {
	type args struct {
		uts1 int64
		uts2 int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"same hour", args{1577808000, 1577811599}, true},      // "2020-01-01 00:00:00" & "2020-01-01 00:59:59"
		{"not same hour", args{1577808000, 1577811600}, false}, // "2020-01-01 00:00:00" & "2020-01-01 01:00:00"
	}
	for _, tt := range tests {
		got := IsSameHour(tt.args.uts1, tt.args.uts2)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestIsSameMinute(t *testing.T) {
	type args struct {
		uts1 int64
		uts2 int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"same hour", args{1577808000, 1577808059}, true},      // "2020-01-01 00:00:00" & "2020-01-01 00:00:59"
		{"not same hour", args{1577808000, 1577808060}, false}, // "2020-01-01 00:00:00" & "2020-01-01 00:01:00"
	}

	for _, tt := range tests {
		got := IsSameMinute(tt.args.uts1, tt.args.uts2)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestIsSameWeek(t *testing.T) {
	type args struct {
		uts1 int64
		uts2 int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"same week", args{1577808000, 1578153600}, true},      // "2020-01-01 00:00:00" & "2020-01-05 00:00:00"
		{"not same week", args{1577808000, 1578412800}, false}, // "2020-01-01 00:00:00" & "2020-01-08 00:00:00"
	}
	for _, tt := range tests {
		got := IsSameWeek(tt.args.uts1, tt.args.uts2)
		assert.Equal(t, tt.want, got, tt.name)
	}
}
