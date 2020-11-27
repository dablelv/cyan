package util

import (
	"math"
	"time"
)

//
// Part 0: Get some useful infomation about time
//

// GetNowS get unix timestamp in second
func GetNowS() int64 {
	return time.Now().Unix()
}

// GetNowMs get unix timestamp in millisecond
func GetNowMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GetNowUs get unix timestamp in microsecond
func GetNowUs() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// GetNowNs get unix timestamp in nanosecond
func GetNowNs() int64 {
	return time.Now().UnixNano()
}

// GetNowDate get now date in YYYY-MM-DD
func GetNowDate() string {
	return time.Now().Format("2006-01-02")
}

// GetNowDate get now time in hh:mm:ss
func GetNowTime() string {
	return time.Now().Format("15:04:05")
}

// GetNowDateTime get now datetime in YYYY-MM-DD hh:mm:ss
func GetNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetNowDateTimeZ get now datetime with zone in YYYY-MM-DD hh:mm:ss Zone
// e.g. 2020-05-11 23:18:07 +08:00
func GetNowDateTimeZ() string {
	return time.Now().Format("2006-01-02 15:04:05 Z07:00")
}

// GetDayBeginMoment get the starting moment of one day
func GetDayBeginMoment(t time.Time) time.Time {
	y, m, d := t.Date()
	n := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	return n
}

// GetDayEndMoment get the ending moment of one day
func GetDayEndMoment(t time.Time) time.Time {
	y, m, d := t.Date()
	n := time.Date(y, m, d, 23, 59, 59, 999999999, time.Local)
	return n
}

// GetDayElapsedS get the elapsed seconds since the starting moment of one day
func GetDayElapsedS(t time.Time) int64 {
	return t.Unix() - GetDayBeginMoment(t).Unix()
}

// GetDayElapsedMs get the elapsed milliseconds since the starting moment of one day
func GetDayElapsedMs(t time.Time) int64 {
	return (t.UnixNano() - GetDayBeginMoment(t).UnixNano()) / int64(time.Millisecond)
}

// GetDayElapsedUs get the elapsed microseconds since the starting moment of one day
func GetDayElapsedUs(t time.Time) int64 {
	return (t.UnixNano() - GetDayBeginMoment(t).UnixNano()) / int64(time.Microsecond)
}

// GetDayElapsedNs get the elapsed nanoseconds since the starting moment of one day
func GetDayElapsedNs(t time.Time) int64 {
	return t.Unix() - GetDayBeginMoment(t).Unix()
}

// GetDaysBtwTs calculates the number of days between two timestamps and round down
func GetDaysBtwTs(ts0, ts1 int64) int64 {
	return int64(math.Abs(float64(ts0-ts1))) / 86400
}

// GetHoursBtwTs calculates the number of hours between two timestamps and round down
func GetHoursBtwTs(ts0, ts1 int64) int64 {
	return int64(math.Abs(float64(ts0-ts1))) / 3600
}

// GetMinutesBtwTs calculates the number of hours between two timestamps and round down
func GetMinutesBtwTs(ts0, ts1 int64) int64 {
	return int64(math.Abs(float64(ts0-ts1))) / 60
}

//
// Part 1: Common conversion about time
//

// DateTime2UTs convert datetime in YYYY-MM-DD hh:mm:ss to unix timestamp
func DateTime2UTs(dt string) int64 {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", dt, loc)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// UTs2DateTime convert unix timestamp to datetime in YYYY-MM-DD hh:mm:ss
func UTs2DateTime(uts int64) string {
	return time.Unix(uts, 0).Format("2006-01-02 15:04:05")
}

//
// Part 2: A time counter to count time interval
//

// TimeCounter is used to count time interval
type TimeCounter struct {
	int64
}

// NewTimeCounter create a time counter
func NewTimeCounter() (t *TimeCounter) {
	t = new(TimeCounter)
	t.Set()
	return t
}

// Set start timing
func (t *TimeCounter) Set() {
	t.int64 = time.Now().UnixNano()
}

// GetS return the time interval from the beginning to now in second
func (t *TimeCounter) GetS() int64 {
	return (time.Now().UnixNano() - t.int64) / int64(time.Second)
}

// GetMs return the time interval from the beginning to now in millisecond
func (t *TimeCounter) GetMs() int64 {
	return (time.Now().UnixNano() - t.int64) / int64(time.Millisecond)
}

// GetUs return the time interval from the beginning to now in microsecond
func (t *TimeCounter) GetUs() int64 {
	return (time.Now().UnixNano() - t.int64) / int64(time.Microsecond)
}

// GetNs return the time interval from the beginning to now in nanosecond
func (t *TimeCounter) GetNs() int64 {
	return time.Now().UnixNano() - t.int64
}
