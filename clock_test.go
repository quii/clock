package clock

import (
	"fmt"
	"testing"
	"time"
)

// Start with a base time fixed to noon, so we can add/subtract without unintentionally crossing into another day.
var baseTime = time.Date(2015, 10, 7, 12, 0, 0, 0, time.Local)

// A list of test cases each containing times with an earlier and later Clock value.
// Changing baseTime's hours/minutes/seconds should affect it's Clock values.
var unEqualCases = []struct{Earlier, Later time.Time}{
	{baseTime, baseTime.Add(2 * time.Second)},
	{baseTime.Add(-2 * time.Second), baseTime},
	{baseTime, baseTime.Add(2 * time.Minute)},
	{baseTime.Add(-2 * time.Minute), baseTime},
	{baseTime, baseTime.Add(2 * time.Hour)},
	{baseTime.Add(-2 * time.Hour), baseTime},
}

// A list of test cases containing Times with equal Clock values.
// Changing baseTime's years/months/days or milli/micro/nanoseconds should not affect it's Clock values.
var equalCases = []struct{T1, T2 time.Time}{
	{baseTime, baseTime},
	{baseTime, baseTime.AddDate(10,0,0)},
	{baseTime, baseTime.AddDate(-5,0,0)},
	{baseTime, baseTime.AddDate(0,8,0)},
	{baseTime, baseTime.AddDate(0,-5,0)},
	{baseTime, baseTime.AddDate(0,0,1)},
	{baseTime, baseTime.AddDate(0,0,-1)},

	{baseTime, baseTime.Add(2 * time.Nanosecond)},
	{baseTime, baseTime.Add(2 * time.Millisecond)},
	{baseTime, baseTime.Add(2 * time.Microsecond)},
}

func TestBefore(t *testing.T) {
	for _,testCase := range unEqualCases {
		if !Before(testCase.Earlier, testCase.Later) {
			t.Errorf("expected %s before %s", testCase.Earlier, testCase.Later)
		}
	}
	for _,testCase := range equalCases {
		if Before(testCase.T1, testCase.T2) {
			t.Errorf("did not expect %s before %s", testCase.T1, testCase.T2)
		}
	}
}

func TestAfter(t *testing.T) {
	for _,testCase := range unEqualCases {
		if !After(testCase.Later, testCase.Earlier) {
			t.Errorf("expected %s after %s", testCase.Later, testCase.Earlier)
		}
	}
	for _,testCase := range equalCases {
		if After(testCase.T1, testCase.T2) {
			t.Errorf("did not expect %s after %s", testCase.T1, testCase.T2)
		}
	}
}

func TestEquals(t *testing.T) {
	for _,testCase := range equalCases {
		if !Equals(testCase.T1, testCase.T2) {
			t.Errorf("expected %s equals %s", testCase.T1, testCase.T2)
		}
	}
	for _,testCase := range unEqualCases {
		if Equals(testCase.Later, testCase.Earlier) {
			t.Errorf("did not expect %s equals %s", testCase.Later, testCase.Earlier)
		}
	}
}

func TestBetween(t *testing.T) {
	for _,testCase := range []struct{
		T1, T2 time.Time
		// Expected duration between
		Expected time.Duration
	} {
		{baseTime, baseTime.Add(10 * time.Second), 10 * time.Second},
		{baseTime, baseTime.Add(23 * time.Hour), -1 * time.Hour},
		{baseTime, baseTime.Add(24 * time.Hour), 0},
		{baseTime, baseTime.Add(24 * time.Hour - (1 * time.Second)), -1 * time.Second},
		{baseTime, baseTime.Add(-1 * time.Second), -1 * time.Second},
		{baseTime, baseTime.AddDate(1, 1, 1), 0},
	} {
		if got := Between(testCase.T1, testCase.T2); got != testCase.Expected {
			t.Errorf("expected %s between %s and %s, but got %s", testCase.Expected, testCase.T1, testCase.T2, got)
		}
	}
}

func TestUntil(t *testing.T) {
	for _,testCase := range []struct {
		T1, T2 time.Time
		// Expected duration until
		Expected time.Duration
	} {
		{baseTime, baseTime.Add(100 * time.Second), 100 * time.Second},
		{baseTime, baseTime.Add(25 * time.Hour), 1 * time.Hour},
		{baseTime, baseTime.Add(24 * time.Hour), 0},
		{baseTime, baseTime.Add(24 * time.Hour - (1 * time.Second)), 24 * time.Hour - (1 * time.Second)},
		{baseTime, baseTime.Add(-1 * time.Second), 24 * time.Hour - (1 * time.Second)},
		{baseTime, baseTime.AddDate(1, 1, 1), 0},
	} {
		if got := Until(testCase.T1, testCase.T2); got != testCase.Expected {
			t.Errorf("expected %s between %s and %s, but got %s", testCase.Expected, testCase.T1, testCase.T2, got)
		}
	}
}

func ExampleAfter() {
	start := time.Now()
	after := start.Add(2 * time.Second)

	fmt.Println(After(after, start))
	// Output: true
}
