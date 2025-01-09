package time

import "time"

const (
	DayTime = "20060102150405"
	DayOnly = "20060102"
)

const (
	DayDuration time.Duration = time.Hour * 24
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
