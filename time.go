package pointer

import "time"

// ToTime get pointer by time.Time
func ToTime(in time.Time) *time.Time {
	return &in
}

// ToDuration get pointer by time.Duration
func ToDuration(in time.Duration) *time.Duration {
	return &in
}
