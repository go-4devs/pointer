package pointer

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	var x time.Duration
	if *ToDuration(x) != x {
		t.Fail()
	}
}

func TestTime(t *testing.T) {
	var x time.Time
	if *ToTime(x) != x {
		t.Fail()
	}
}
