package time

import (
	"fmt"
	"time"
)

// FromString parses a formatted string and returns the time value it represents.
// Local zone will be used if loc is nil
func FromString(layout, value string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		loc = time.Local
	}

	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return t, fmt.Errorf("invalid time string: %w", err)
	}
	return t, nil
}

// FromSec returns the given zone Time corresponding to the "Unix second".
// Local zone will be used if loc is nil.
func FromSec(sec int64, loc *time.Location) time.Time {
	t := time.Unix(sec, 0)
	if loc != nil {
		t = t.In(loc)
	}
	return t
}

// FromMs returns the given zone Time corresponding to the "Unix millisecond"
// Local zone will be used if loc is nil
func FromMs(ms int64, loc *time.Location) time.Time {
	t := time.UnixMilli(ms)
	if loc != nil {
		t = t.In(loc)
	}
	return t
}

// FromUs returns the given zone Time corresponding to the "Unix microsecond"
// Local zone will be used if loc is nil
func FromUs(us int64, loc *time.Location) time.Time {
	t := time.UnixMicro(us)
	if loc != nil {
		t = t.In(loc)
	}
	return t
}

// FromNs converts a nanosecond timestamp to time.Time in the specified location.
func FromNs(ns int64, loc *time.Location) time.Time {
	sec := ns / 1e9
	nsec := ns % 1e9

	t := time.Unix(sec, nsec)
	if loc != nil {
		t = t.In(loc)
	}
	return t
}

// GetBeijingTime gets Beijing Time from time layout and value string.
// The location name Asia/Shanghai from IANA Time Zone Database standards for Beijing Time.
func GetBeijingTime(layout, value string) (t time.Time, err error) {
	loc, err := time.LoadLocation(AsiaShanghai)
	if err != nil {
		return
	}
	return time.ParseInLocation(layout, value, loc)
}

// GetTimezoneTime gets time.Time based on specified timezone, layout and value string.
func GetTimezoneTime(timezone string, layout, value string) (t time.Time, err error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return
	}
	return time.ParseInLocation(layout, value, loc)
}
