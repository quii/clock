package clock
import "time"

// Clock defines the datatype with which we can do time based comparisons. It's worth noting that time.Time implements this so you can just pass those through to the functions in this lib
type Clock interface {
	Clock() (hours int, minutes int, seconds int)
}

// After will tell you if c1 is later than c2
func After(c1 Clock, c2 Clock) bool {
	return timeToSeconds(c1) > timeToSeconds(c2)
}

// Before will tell you if c1 is earlier than c2
func Before(c1 Clock, c2 Clock) bool {
	return timeToSeconds(c1) < timeToSeconds(c2)
}

// Equals will tell you if c1 is equal to c2
func Equals(c1 Clock, c2 Clock) bool {
	return timeToSeconds(c1) == timeToSeconds(c2)
}

// Between returns the duration in seconds between c1 and c2.
// Note that the result is negative if c1 is after c2.
func Between(c1 Clock, c2 Clock) time.Duration {
	return time.Duration(timeToSeconds(c2) - timeToSeconds(c1)) * time.Second
}

// Until returns the duration in seconds from c1 until c2, going forward in time.
// Result is always non-negative.
func Until(c1 Clock, c2 Clock) time.Duration {
	if After(c1, c2) {
		return 24 * time.Hour - Between(c2, c1)
	} else {
		// Before(c1, c2) or Equals(c1, c2)
		return Between(c1, c2)
	}
}

func timeToSeconds(c Clock) int {
	hours, mins, secs := c.Clock()
	return hours*60*60 + mins*60 + secs
}
