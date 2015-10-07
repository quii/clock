package clock

import (
	"fmt"
	"testing"
	"time"
)

func TestItComparesTimesCorrectly(t *testing.T) {
	startTime := time.Now()
	earlier := startTime.Add(-2 * time.Second)
	later := startTime.Add(2 * time.Second)

	if !After(startTime, earlier) {
		t.Error("start time should be seen as being after the earlier time")
	}

	if !Before(startTime, later) {
		t.Error("the start time should be seen as before the later time")
	}
}

func ExampleAfter() {
	start := time.Now()
	after := start.Add(2 * time.Second)

	fmt.Println(After(after, start))
	// Output: true
}
