package clock

// Clock defines the datatype with which we can do time based comparisons
type Clock interface {
	Clock() (hours int, minutes int, seconds int)
}

// After will tell you if c1 is later than c2
func After(c1 Clock, c2 Clock) bool {
	return timeToSeconds(c1) > timeToSeconds(c2)
}

// Before will tell you if c1 is earlier than c2
func Before(c1 Clock, c2 Clock) bool {
	return !After(c1, c2)
}

func timeToSeconds(c Clock) int {
	hours, mins, secs := c.Clock()
	return hours*60*60 + mins*60 + secs
}
