package time

import "time"

//
// Desc: Common conversion about time.
//

// DateTimeToTs converts datetime in YYYY-MM-DD hh:mm:ss to unix timestamp.
func DateTimeToTs(dt string) int64 {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(DateTimeFormat, dt, loc)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// TsToDateTime converts unix timestamp to local datetime in YYYY-MM-DD hh:mm:ss.
func TsToDateTime(ts int64) string {
	return time.Unix(ts, 0).Format(DateTimeFormat)
}
