package time

import (
	"testing"
	"time"

	"github.com/dablelv/cyan/internal"
)

func TestGetDaysBtwTs(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetDaysBtwTs")
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
		assert.Equal(tt.want, got)
	}
}

func TestGetHoursBtwTs(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetHoursBtwTs")
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
		assert.Equal(tt.want, got)
	}
}

func TestGetMinutesBtwTs(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetMinutesBtwTs")
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
		assert.Equal(tt.want, got)
	}
}

func TestIsLeapYear(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsLeapYear")

	assert.Equal(true, IsLeapYear(2000))
	assert.Equal(false, IsLeapYear(2001))
	assert.Equal(false, IsLeapYear(2100))
}

func TestIsSameYear(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsSameYear")
	type args struct {
		uts1 int64
		uts2 int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"same year", args{1577808000, 1577808000}, true},      // "2020-01-01T00:00:00Z08:00"
		{"not same year", args{1577808000, 1609430400}, false}, // "2020-01-01 00:00:00" & "2021-01-01 00:00:00"
	}
	for _, tt := range tests {
		got := IsSameYear(tt.args.uts1, tt.args.uts2)
		assert.Equal(tt.want, got)
	}
}

func TestIsSameMonth(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsSameMonth")
	type args struct {
		uts1 int64
		uts2 int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"same month", args{1577808000, 1577808000}, true},      // "2020-01-01T00:00:00Z08:00"
		{"not same month", args{1577808000, 1580486400}, false}, // "2020-01-01 00:00:00" & "2020-02-01 00:00:00"
	}
	for _, tt := range tests {
		got := IsSameMonth(tt.args.uts1, tt.args.uts2)
		assert.Equal(tt.want, got)
	}
}

func TestIsSameDay(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsSameDay")
	type args struct {
		uts1 int64
		uts2 int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"same day", args{1577808000, 1577808000}, true},      // "2020-01-01T00:00:00Z08:00"
		{"not same day", args{1577808000, 1577894400}, false}, // "2020-01-01T00:00:00Z08:00" & "2020-01-02T00:00:00Z08:00"
	}
	for _, tt := range tests {
		got := IsSameDay(tt.args.uts1, tt.args.uts2)
		assert.Equal(tt.want, got)
	}
}

func TestIsSameHour(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsSameHour")
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
		assert.Equal(tt.want, got)
	}
}

func TestIsSameMinute(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsSameMinute")
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
		assert.Equal(tt.want, got)
	}
}

func TestIsSameWeek(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsSameWeek")
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
		assert.Equal(tt.want, got)
	}
}

func TestGetNowUtc(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetNowUtc")
	name, offset := GetNowUtc().Zone()
	assert.Equal("UTC", name)
	assert.Equal(0, offset)
}

func TestGetBeijingTime(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetBeijingTime")

	got, err := GetBeijingTime(time.DateTime, "2022-11-04 08:30:00")
	assert.IsNil(err)
	assert.Equal("2022-11-04T08:30:00+08:00", got.Format(time.RFC3339))

	_, err = GetBeijingTime(time.RFC3339, "2022-11-04 08:30:00")
	assert.IsNotNil(err)
}
