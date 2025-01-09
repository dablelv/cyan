package time

import (
	"fmt"
	"time"
)

// GetLocByUtcOffset gets a time.Location based on the offset to UTC.
func GetLocByUtcOffset(offset time.Duration) *time.Location {
	if offset == time.Duration(0) {
		return time.UTC
	}

	hours := int(offset.Hours())
	minutes := int(offset.Minutes() - float64(hours)*60)
	sign := "+"
	if offset < 0 {
		sign = "-"
		hours = -hours
		minutes = -minutes
	}

	// custom zone name.
	name := fmt.Sprintf("UTC%s%02d:%02d", sign, hours, minutes)

	loc := time.FixedZone(name, int(offset.Seconds()))
	return loc
}

func loadLocation(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(fmt.Sprintf("Failed to load location: %s", err))
	}
	return loc
}
