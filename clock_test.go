package clock

import (
	"fmt"
	"testing"
	"time"
)

// Start with a base time fixed to noon, so we can add/subtract without unintentionally crossing into another day.
var baseTime = time.Date(2015, 10, 7, 12, 0, 0, 0, time.Local)

// A list of test cases each containing times with an earlier and later clock value.
// Changing baseTime's hours/minutes/seconds should affect it's Clock value.
var unEqualCases = []struct{Earlier, Later time.Time}{
	{baseTime, baseTime.Add(2 * time.Second)},
	{baseTime.Add(-2 * time.Second), baseTime},
	{baseTime, baseTime.Add(2 * time.Minute)},
	{baseTime.Add(-2 * time.Minute), baseTime},
	{baseTime, baseTime.Add(2 * time.Hour)},
	{baseTime.Add(-2 * time.Hour), baseTime},
}

// A list of test cases containing Times with equal clock values.
// Changing baseTime's years/months/days or milli/micor/nanoseconds should not affect it's Clock value.
var equalCases = []struct{T1, T2 time.Time}{
	{baseTime, baseTime},
	{baseTime, baseTime.AddDate(10,0,0)},
	{baseTime, baseTime.AddDate(-5,0,0)},
	{baseTime, baseTime.AddDate(0,8,0)},
	{baseTime, baseTime.AddDate(0,-5,0)},
	{baseTime, baseTime.AddDate(0,0,1)},
	{baseTime, baseTime.AddDate(0,0,-1)},

	{baseTime, baseTime.Add(2 * time.Nanosecond)},
	{baseTime.Add(2 * time.Nanosecond), baseTime},
	{baseTime, baseTime.Add(2 * time.Millisecond)},
	{baseTime.Add(2 * time.Millisecond), baseTime},
	{baseTime, baseTime.Add(2 * time.Microsecond)},
	{baseTime.Add(2 * time.Microsecond), baseTime},
}

func TestBefore(t *testing.T) {
	for _,testCase := range unEqualCases {
		if !Before(testCase.Earlier, testCase.Later) {
			t.Errorf("expected %s before %s", testCase.Earlier.String(), testCase.Later.String())
		}
	}
	for _,testCase := range equalCases {
		if Before(testCase.T1, testCase.T2) {
			t.Errorf("did not expect %s before %s", testCase.T1.String(), testCase.T2.String())
		}
	}
}

func TestAfter(t *testing.T) {
	for _,testCase := range unEqualCases {
		if !After(testCase.Later, testCase.Earlier) {
			t.Errorf("expected %s after %s", testCase.Later.String(), testCase.Earlier.String())
		}
	}
	for _,testCase := range equalCases {
		if After(testCase.T1, testCase.T2) {
			t.Errorf("did not expect %s after %s", testCase.T1.String(), testCase.T2.String())
		}
	}
}

func TestEquals(t *testing.T) {
	for _,testCase := range equalCases {
		if !Equals(testCase.T1, testCase.T2) {
			t.Errorf("expected %s equals %s", testCase.T1.String(), testCase.T2.String())
		}
	}
	for _,testCase := range unEqualCases {
		if Equals(testCase.Later, testCase.Earlier) {
			t.Errorf("did not expect %s equals %s", testCase.Later.String(), testCase.Earlier.String())
		}
	}
}

func ExampleAfter() {
	start := time.Now()
	after := start.Add(2 * time.Second)

	fmt.Println(After(after, start))
	// Output: true
}
