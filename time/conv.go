package time

import "time"

//
// Desc: Common conversion about time.
//

// DateTime2UTS converts datetime in YYYY-MM-DD hh:mm:ss to unix timestamp
func DateTime2UTS(dt string) int64 {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(DateTimeFormat, dt, loc)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// UTS2DateTime converts unix timestamp to datetime in YYYY-MM-DD hh:mm:ss.
func UTS2DateTime(uts int64) string {
	return time.Unix(uts, 0).Format(DateTimeFormat)
}
