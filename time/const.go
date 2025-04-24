package time

import "time"

const (
	YearOnly         = "2006"
	YearMonth        = "200601"
	DayOnly          = "20060102"
	DayHour          = "2006010215"
	DayHourMinute    = "200601021504"
	DayHourMinuteSec = "20060102150405"
	HourOnly         = "15"
	HourMinute       = "1504"
	HourMinuteSec    = "150405"
	DateTimeMilli    = "2006-01-02 15:04:05.000"
	DateTimeMicro    = "2006-01-02 15:04:05.000000"
	DateTimeNano     = "2006-01-02 15:04:05.000000000"
	DateTimeZone     = "2006-01-02 15:04:05 Z07:00"
)

const (
	HoursPerDay                  = 24
	DayDuration    time.Duration = time.Hour * HoursPerDay
	SecondsPerHour               = int64(time.Hour / time.Second)
	SecondsPerDay                = int64(DayDuration / time.Second)
)

// Common timezone strings defined as constants.
// Those timezone standard for a location name corresponding to a file in the IANA(Internet Assigned Numbers Authority) Time Zone database.
const (
	AsiaShanghai      = "Asia/Shanghai"       // China Standard Time (CST)
	AmericaNewYork    = "America/New_York"    // Eastern Standard Time (EST/EDT)
	AmericaLosAngeles = "America/Los_Angeles" // Pacific Standard Time (PST/PDT)
	EuropeLondon      = "Europe/London"       // UK Time (GMT/BST)
	EuropeBerlin      = "Europe/Berlin"       // Central European Time (CET/CEST)
	AsiaTokyo         = "Asia/Tokyo"          // Japan Standard Time (JST)
	AustraliaSydney   = "Australia/Sydney"    // Australian Eastern Standard Time (AEST/AEDT)
	AsiaHongKong      = "Asia/Hong_Kong"      // Hong Kong Time (HKT)
	AsiaSingapore     = "Asia/Singapore"      // Singapore Time (SGT)
	PacificAuckland   = "Pacific/Auckland"    // New Zealand Standard Time (NZST/NZDT)
	EuropeParis       = "Europe/Paris"        // Central European Time (CET/CEST)
	AsiaSeoul         = "Asia/Seoul"          // Korea Standard Time (KST)
)

// Location constants.
var (
	LocAsiaShanghai      = loadLocation(AsiaShanghai)
	LocAmericaNewYork    = loadLocation(AmericaNewYork)
	LocAmericaLosAngeles = loadLocation(AmericaLosAngeles)
	LocEuropeLondon      = loadLocation(EuropeLondon)
	LocEuropeBerlin      = loadLocation(EuropeBerlin)
	LocAsiaTokyo         = loadLocation(AsiaTokyo)
	LocAustraliaSydney   = loadLocation(AustraliaSydney)
	LocAsiaHongKong      = loadLocation(AsiaHongKong)
	LocAsiaSingapore     = loadLocation(AsiaSingapore)
	LocPacificAuckland   = loadLocation(PacificAuckland)
	LocEuropeParis       = loadLocation(EuropeParis)
	LocAsiaSeoul         = loadLocation(AsiaSeoul)
)
