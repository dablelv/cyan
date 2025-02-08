// Package time provides some useful functions about time.
package time

import (
	"math"
	"time"
)

const (
	YFormatNum        = "2006"
	YMFormatNum       = "200601"
	DateFormatNum     = "20060102"
	DateHFormatNum    = "2006010215"
	DateHMFormatNum   = "200601021504"
	DateTimeFormatNum = "20060102150405"
	HFormatNum        = "15"
	HMFormatNum       = "1504"
	TimeFormatNum     = "150405"
	// Deprecated: please use standard library time.DateOnly since go 1.20
	DateFormat = "2006-01-02"
	// Deprecated: please use standard library time.TimeOnly since go 1.20
	TimeFormat = "15:04:05"
	// Deprecated: please use standard library time.DateTime since go 1.20
	DateTimeFormat      = "2006-01-02 15:04:05"
	DateTimeFormatMilli = "2006-01-02 15:04:05.000"
	DateTimeFormatMicro = "2006-01-02 15:04:05.000000"
	DateTimeFormatNano  = "2006-01-02 15:04:05.000000000"
)

// GetNowS gets unix timestamp in second.
func GetNowS() int64 {
	return time.Now().Unix()
}

// GetNowMs gets unix timestamp in millisecond.
func GetNowMs() int64 {
	return time.Now().UnixMilli()
}

// GetNowUs gets unix timestamp in microsecond.
func GetNowUs() int64 {
	return time.Now().UnixMicro()
}

// GetNowNs gets unix timestamp in nanosecond.
func GetNowNs() int64 {
	return time.Now().UnixNano()
}

// GetNowDate gets now date in YYYY-MM-DD.
func GetNowDate() string {
	return time.Now().Format(DateFormat)
}

// GetNowDate gets now time in hh:mm:ss.
func GetNowTime() string {
	return time.Now().Format(TimeFormat)
}

// GetNowDateTime gets now datetime in YYYY-MM-DD hh:mm:ss.
func GetNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetNowDateTimeZ gets now datetime with zone in YYYY-MM-DD hh:mm:ss Zone
// e.g. 2020-05-11 23:18:07 +08:00.
func GetNowDateTimeZ() string {
	return time.Now().Format("2006-01-02 15:04:05 Z07:00")
}

// GetDayBeginMoment gets the starting moment of one day.
func GetDayBeginMoment(t time.Time) time.Time {
	y, m, d := t.Date()
	n := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	return n
}

// GetDayBeginMomentTs gets the starting moment of one day specified by UNIX timestamp.
func GetDayBeginMomentTs(ts int64) time.Time {
	y, m, d := time.Unix(ts, 0).Date()
	n := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	return n
}

// GetDayEndMoment gets the ending moment of one day.
func GetDayEndMoment(t time.Time) time.Time {
	y, m, d := t.Date()
	n := time.Date(y, m, d, 23, 59, 59, 999999999, time.Local)
	return n
}

// GetDayEndMomentTs gets the ending moment of one day specified by UNIX timestamp.
func GetDayEndMomentTs(uts int64) time.Time {
	y, m, d := time.Unix(uts, 0).Date()
	n := time.Date(y, m, d, 23, 59, 59, 999999999, time.Local)
	return n
}

// GetDayElapsedS gets the elapsed seconds since the starting moment of one day.
func GetDayElapsedS(t time.Time) int64 {
	return t.Unix() - GetDayBeginMoment(t).Unix()
}

// GetDayElapsedMs gets the elapsed milliseconds since the starting moment of one day.
func GetDayElapsedMs(t time.Time) int64 {
	return (t.UnixNano() - GetDayBeginMoment(t).UnixNano()) / int64(time.Millisecond)
}

// GetDayElapsedUs gets the elapsed microseconds since the starting moment of one day.
func GetDayElapsedUs(t time.Time) int64 {
	return (t.UnixNano() - GetDayBeginMoment(t).UnixNano()) / int64(time.Microsecond)
}

// GetDayElapsedNs gets the elapsed nanoseconds since the starting moment of one day.
func GetDayElapsedNs(t time.Time) int64 {
	return t.UnixNano() - GetDayBeginMoment(t).UnixNano()
}

// GetDaysBtwTs gets the number of days between two timestamps and round down.
func GetDaysBtwTs(ts0, ts1 int64) int64 {
	return int64(math.Abs(float64(ts0-ts1))) / 86400
}

// GetHoursBtwTs gets the number of hours between two timestamps and round down.
func GetHoursBtwTs(ts0, ts1 int64) int64 {
	return int64(math.Abs(float64(ts0-ts1))) / 3600
}

// GetMinutesBtwTs gets the number of hours between two timestamps and round down.
func GetMinutesBtwTs(ts0, ts1 int64) int64 {
	return int64(math.Abs(float64(ts0-ts1))) / 60
}

// GetWeekday gets the weekday time.
func GetWeekday(t time.Time, w time.Weekday) time.Time {
	if t.Weekday() == w {
		return t
	}
	d := w - t.Weekday()
	if w == time.Sunday {
		d += 7
	} else if t.Weekday() == time.Sunday {
		d -= 7
	}
	return t.AddDate(0, 0, int(d))
}

// GetMonDate gets monday date in format 2006-01-02.
func GetMonDate(t time.Time) string {
	return GetWeekday(t, time.Monday).Format(DateFormat)
}

// GetTuesDate gets tuesday date in format 2006-01-02.
func GetTuesDate(t time.Time) string {
	return GetWeekday(t, time.Tuesday).Format(DateFormat)
}

// GetWedDate gets wednesday date in format 2006-01-02.
func GetWedDate(t time.Time) string {
	return GetWeekday(t, time.Wednesday).Format(DateFormat)
}

// GetThursDate gets thursday date in format 2006-01-02.
func GetThursDate(t time.Time) string {
	return GetWeekday(t, time.Thursday).Format(DateFormat)
}

// GetFriDate gets friday date in format 2006-01-02.
func GetFriDate(t time.Time) string {
	return GetWeekday(t, time.Friday).Format(DateFormat)
}

// GetSatDate gets saturday date in format 2006-01-02.
func GetSatDate(t time.Time) string {
	return GetWeekday(t, time.Saturday).Format(DateFormat)
}

// GetSunDate gets sunday date in format 2006-01-02.
func GetSunDate(t time.Time) string {
	return GetWeekday(t, time.Sunday).Format(DateFormat)
}

// IsLeapYear checks the year whether is leap year.
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

// IsSameYear checks the unix timestamp whether is the same year.
func IsSameYear(ts1, ts2 int64) bool {
	t1 := time.Unix(ts1, 0)
	t2 := time.Unix(ts2, 0)
	return t1.Format(YFormatNum) == t2.Format(YFormatNum)
}

// IsSameMonth checks the unix timestamp whether is the same month.
func IsSameMonth(ts1, ts2 int64) bool {
	t1 := time.Unix(ts1, 0)
	t2 := time.Unix(ts2, 0)
	return t1.Format(YMFormatNum) == t2.Format(YMFormatNum)
}

// IsSameDay checks the unix timestamp whether is the same day.
func IsSameDay(ts1, ts2 int64) bool {
	t1 := time.Unix(ts1, 0)
	t2 := time.Unix(ts2, 0)
	return t1.Format(DateFormatNum) == t2.Format(DateFormatNum)
}

// IsSameHour checks the unix timestamp whether is the same hour.
func IsSameHour(ts1, ts2 int64) bool {
	t1 := time.Unix(ts1, 0)
	t2 := time.Unix(ts2, 0)
	return t1.Format(DateHFormatNum) == t2.Format(DateHFormatNum)
}

// IsSameMinute checks the unix timestamp whether is the same minute.
func IsSameMinute(ts1, ts2 int64) bool {
	t1 := time.Unix(ts1, 0)
	t2 := time.Unix(ts2, 0)
	return t1.Format(DateHMFormatNum) == t2.Format(DateHMFormatNum)
}

// IsSameWeek checks the unix timestamp whether is the same week.
func IsSameWeek(ts1, ts2 int64) bool {
	t1 := time.Unix(ts1, 0)
	t2 := time.Unix(ts2, 0)
	return GetMonDate(t1) == GetMonDate(t2)
}

// GetDayMomentNanoTs gets the nanosecond timestamp based on specified day and moment.
// If you want get more info about timezone, please refer to [IANA Time Zone Database](https://www.iana.org/time-zones).
// Common timezone are: Asia/Shanghai, America/New_York, Europe/London etc.
func GetDayMomentNanoTs(t time.Time, timezone string, hour, minute, second, nsec int) (int64, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 0, err
	}

	year, month, day := t.In(loc).Date()
	r := time.Date(year, month, day, hour, minute, second, nsec, loc)
	return r.UnixNano(), nil
}

// GetDayMomentTime gets the time.Time based on specified day and moment.
func GetDayMomentTime(t time.Time, timezone string, hour, minute, second, nsec int) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	year, month, day := t.In(loc).Date()
	return time.Date(year, month, day, hour, minute, second, nsec, loc), nil
}

// NowUtc returns now UTC time.
func NowUtc() time.Time {
	return time.Now().UTC()
}

// NowEst returns now EST(Eastern Standard Time) time.
func NowEst() time.Time {
	return time.Now().In(LocAmericaNewYork)
}

// NowCst returns now CST(China Standard Time) time.
func NowCst() time.Time {
	return time.Now().In(LocAsiaShanghai)
}

// TodayMidnight returns the time of today's midnight (00:00:00).
func TodayMidnight() time.Time {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, n.Location())
}

// IsCloseIn check the cond of left <= target <= right.
func IsCloseIn(target, left, right time.Time) bool {
	if right.Before(left) {
		return false
	}
	return !target.Before(left) && !target.After(right)
}

// IsLeftCloseIn check the cond of left <= target < right.
func IsLeftCloseIn(target, left, right time.Time) bool {
	if !right.After(left) {
		return false
	}
	return !target.Before(left) && target.Before(right)
}

// IsRightCloseIn check the cond of left < target <= right.
func IsRightCloseIn(target, left, right time.Time) bool {
	if !right.After(left) {
		return false
	}
	return target.After(left) && !target.After(right)
}

// IsOpenIn check the cond of left < target < right.
func IsOpenIn(target, left, right time.Time) bool {
	if !right.After(left) {
		return false
	}
	return target.After(left) && target.Before(right)
}

// Sec return "unix second" of the given time.
func Sec(t time.Time) int64 {
	return t.Unix()
}

// Ms return "unix millisecond" of the given time.
func Ms(t time.Time) int64 {
	return t.UnixMilli()
}

// Us return "unix microsecond" of the given time.
func Us(t time.Time) int64 {
	return t.UnixMicro()
}

// Ns return "unix nanosecond" of the given time.
func Ns(t time.Time) int64 {
	return t.UnixNano()
}
