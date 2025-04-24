package time

import "time"

// InCloseIntvl checks the target time is within closed interval [left, right].
func InCloseIntvl(target, left, right time.Time) bool {
	if right.Before(left) {
		return false
	}
	return !target.Before(left) && !target.After(right)
}

// InLeftCloseIntvl checks the target time is within left-closed interval [left, right).
func InLeftCloseIntvl(target, left, right time.Time) bool {
	if !right.After(left) {
		return false
	}
	return !target.Before(left) && target.Before(right)
}

// InRightCloseIntvl checks the target time is within right-closed interval (left, right].
func InRightCloseIntvl(target, left, right time.Time) bool {
	if !right.After(left) {
		return false
	}
	return target.After(left) && !target.After(right)
}

// InOpenIntvl checks the target time is within open interval (left, right).
func InOpenIntvl(target, left, right time.Time) bool {
	if !right.After(left) {
		return false
	}
	return target.After(left) && target.Before(right)
}
