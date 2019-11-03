package handler

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type test struct {
	input    string
	expected *event
}

func TestParseEvent(t *testing.T) {
	now := time.Date(2020, 1, 1, 14, 0, 0, 0, time.Local)
	tests := []test{
		{
			input:    " 10s ",
			expected: nil,
		},
		{
			input:    "10s Some numbers 3s2 ",
			expected: makeTest("Some numbers 3s2", now.Add(-10*time.Second), now),
		},
		{
			input:    " 5 m hello!",
			expected: makeTest("hello!", now.Add(-5*time.Minute), now),
		},
		{
			input:    " 2h  l9l  ",
			expected: makeTest("l9l", now.Add(-2*time.Hour), now),
		},
		{
			input:    "...this is a task 35s",
			expected: makeTest("...this is a task", now.Add(-35*time.Second), now),
		},
		{
			input:    " ...   8m",
			expected: makeTest("...", now.Add(-8*time.Minute), now),
		},
		{
			input:    "Task 1 h",
			expected: makeTest("Task", now.Add(-1*time.Hour), now),
		},
	}

	for _, tt := range tests {
		actual := parseEvent(tt.input, now)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("parseEvent(\"%s\", now): expected %s, actual %s",
				tt.input, prettyPrint(tt.expected), prettyPrint(actual))
		}
	}
}

func makeTest(description string, start, end time.Time) *event {
	e := event{description: description, start: start, end: end}
	return &e
}

func prettyPrint(tt *event) string {
	if tt == nil {
		return "nil"
	}
	return fmt.Sprintf("(\"%s\", \"%s\", \"%s\")", tt.description, tt.start, tt.end)
}
