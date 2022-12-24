package time

import "time"

//
// Desc: A time counter to count time interval.
//

// TimeCounter is used to count time interval.
type TimeCounter struct {
	time.Time
	int64
}

// NewTimeCounter create a time counter.
func NewTimeCounter() (t *TimeCounter) {
	t = new(TimeCounter)
	t.Set()
	return t
}

// Set start timing.
func (t *TimeCounter) Set() {
	t.Time = time.Now()
	t.int64 = t.Time.UnixNano()
}

// GetD return the time interval from the beginning to now in time.Duration.
func (t *TimeCounter) GetD() time.Duration {
	return time.Since(t.Time)
}

// GetS return the time interval from the beginning to now in second.
func (t *TimeCounter) GetS() int64 {
	return (time.Now().UnixNano() - t.int64) / int64(time.Second)
}

// GetMs return the time interval from the beginning to now in millisecond.
func (t *TimeCounter) GetMs() int64 {
	return (time.Now().UnixNano() - t.int64) / int64(time.Millisecond)
}

// GetUs return the time interval from the beginning to now in microsecond.
func (t *TimeCounter) GetUs() int64 {
	return (time.Now().UnixNano() - t.int64) / int64(time.Microsecond)
}

// GetNs return the time interval from the beginning to now in nanosecond.
func (t *TimeCounter) GetNs() int64 {
	return time.Now().UnixNano() - t.int64
}

// TimeCost count time cost.
func TimeCost() func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		return time.Since(start)
	}
}
