package time

import (
	"fmt"
	"strconv"
	"time"
)

// FromYmd parses a yyyymmdd integer like: 20080808 for the loc and returns the time value.
// Attention: the malform value will cause error, like 200808080, 20080808 etc.
// Local zone will be used if loc is nil.
func FromYmd(day int, loc *time.Location) (time.Time, error) {
	if day <= 0 {
		return time.Time{}, fmt.Errorf("minus day: %v", day)
	}
	return FromString(DayOnly, strconv.Itoa(int(day)), loc)
}

// ToYmd return the yyyymmdd integer, like "20080808", is very useful in finance.
// Must sure the year has four digit!
func ToYmd(t time.Time) int {
	y, m, d := t.Date()
	return y*10000 + int(m)*100 + d
}

type Ymd int

// Ok checks the yyyymmdd integer is valid.
func (y Ymd) Ok() bool {
	_, err := FromYmd(int(y), nil)
	return err == nil
}

// String returns the ymd string format, like "20080808".
func (y Ymd) String() string {
	return strconv.Itoa(int(y))
}

// Add returns the yyyymmdd integer after adding days.
func (y Ymd) Add(days int) (Ymd, error) {
	t, err := FromYmd(int(y), nil)
	if err != nil {
		return Ymd(0), fmt.Errorf("invalid yymmdd: %d", int(y))
	}
	day := ToYmd(t.AddDate(0, 0, days))
	return Ymd(day), nil
}

// Next returns the next yyyymmdd.
func (y Ymd) Next() (Ymd, error) {
	return y.Add(1)
}

// Prev return the prev yyyymmdd.
func (y Ymd) Prev() (Ymd, error) {
	return y.Add(-1)
}

// Sub return the duration of natural days for y - u (like: 20080808 - 20080731 = 8)
func (y Ymd) Sub(u Ymd) (int, error) {
	ty, ey := FromYmd(int(y), nil)
	if ey != nil {
		return 0, fmt.Errorf("invalid yyyymmdd: %d", y)
	}

	tu, eu := FromYmd(int(u), nil)
	if eu != nil {
		return 0, fmt.Errorf("invalid yyyymmdd: %d", u)
	}

	return int(ty.Sub(tu) / DayDuration), nil
}
