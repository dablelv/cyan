package time

import "time"

//
// A time counter to count time interval.
//

// TimeCounter is used to count time interval.
type TimeCounter struct {
	time.Time
	int64
}

// NewTimeCounter creates a time counter.
func NewTimeCounter() (t *TimeCounter) {
	t = new(TimeCounter)
	t.Reset()
	return
}

// Reset sets start time to now.
func (t *TimeCounter) Reset() {
	t.Time = time.Now()
	t.int64 = t.Time.UnixNano()
}

// GetD returns the time interval from the beginning to now in time.Duration.
func (t *TimeCounter) GetD() time.Duration {
	return time.Since(t.Time)
}

// GetS returns the time interval from the beginning to now in second.
func (t *TimeCounter) GetS() int64 {
	return (time.Now().UnixNano() - t.int64) / int64(time.Second)
}

// GetMs returns the time interval from the beginning to now in millisecond.
func (t *TimeCounter) GetMs() int64 {
	return (time.Now().UnixNano() - t.int64) / int64(time.Millisecond)
}

// GetUs returns the time interval from the beginning to now in microsecond.
func (t *TimeCounter) GetUs() int64 {
	return (time.Now().UnixNano() - t.int64) / int64(time.Microsecond)
}

// GetNs returns the time interval from the beginning to now in nanosecond.
func (t *TimeCounter) GetNs() int64 {
	return time.Now().UnixNano() - t.int64
}

// TimeCost counts time cost.
func TimeCost() func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		return time.Since(start)
	}
}
