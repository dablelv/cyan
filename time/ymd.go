package time

import (
	"fmt"
	"strconv"
	"time"
)

// Ymd represents a date as an integer in the format YYYYMMDD.
// For example, June 15, 2025 would be represented as 20250615.
type Ymd int

// Ok checks the yyyymmdd integer is valid.
func (y Ymd) Ok() bool {
	_, err := YmdToTime(y, nil)
	return err == nil
}

// String returns the ymd string format, like "20080808".
func (y Ymd) String() string {
	return strconv.Itoa(int(y))
}

// Next returns the next yyyymmdd.
func (y Ymd) Next() (Ymd, error) {
	return y.Add(1)
}

// Prev return the prev yyyymmdd.
func (y Ymd) Prev() (Ymd, error) {
	return y.Add(-1)
}

// Add returns a new Ymd value that is the result of adding the specified
// number of days to the current Ymd value.
func (y Ymd) Add(days int) (Ymd, error) {
	t, err := YmdToTime(y, nil)
	if err != nil {
		return Ymd(0), fmt.Errorf("invalid yymmdd: %d", int(y))
	}
	day := TimeToYmd(t.AddDate(0, 0, days))
	return Ymd(day), nil
}

// Sub return the duration of natural days for y - u (like: 20080808 - 20080731 = 8)
func (y Ymd) Sub(u Ymd) (int, error) {
	ty, ey := YmdToTime(y, nil)
	if ey != nil {
		return 0, fmt.Errorf("invalid yyyymmdd: %d", y)
	}

	tu, eu := YmdToTime(u, nil)
	if eu != nil {
		return 0, fmt.Errorf("invalid yyyymmdd: %d", u)
	}

	return int(ty.Sub(tu) / DayDuration), nil
}

// ToTime converts the Ymd value to a time.Time in the specified location.
// It panics if the Ymd value is invalid or cannot be converted to a time.Time.
func (y Ymd) ToTime(loc *time.Location) time.Time {
	t, err := YmdToTime(y, loc)
	if err != nil {
		panic(fmt.Errorf("invalid ymd(%d):%w", int(y), err))
	}
	return t
}

// YmdToTime converts a yyyymmdd integer like: 20080808 for the loc and returns the time value.
// Attention: the malform value will cause error, like 200808080, -20080808 etc.
// Local zone will be used if loc is nil.
func YmdToTime(day Ymd, loc *time.Location) (time.Time, error) {
	if day <= 0 {
		return time.Time{}, fmt.Errorf("minus day: %v", day)
	}
	return FromString(DayOnly, strconv.Itoa(int(day)), loc)
}

// TimeToYmd return the yyyymmdd integer, like "20080808", is very useful in finance.
// Must sure the year has four digit!
func TimeToYmd(t time.Time) Ymd {
	y, m, d := t.Date()
	return Ymd(y*10000 + int(m)*100 + d)
}
