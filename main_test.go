package main

import (
	"testing"
	"time"
)

func Test_formatDuration(t *testing.T) {
	cases := map[string]time.Duration{
		"0:00":   -5 * time.Second,
		"0:01":   1 * time.Second,
		"0:05":   5 * time.Second,
		"0:30":   29*time.Second + 500*time.Millisecond,
		"0:59":   59 * time.Second,
		"1:00":   1 * time.Minute,
		"1:01":   1*time.Minute + 1*time.Second,
		"123:30": 123*time.Minute + 30*time.Second,
	}

	for expected, d := range cases {
		actual := formatDuration(d)
		if actual != expected {
			t.Fatalf(
				"bad: %#v, expected: %#v, d: %v",
				actual,
				expected,
				d,
			)
		}
	}
}

func Test_parseDuration(t *testing.T) {
	cases := map[string]time.Duration{
		"5":     5 * time.Second,
		"2m30s": 2*time.Minute + 30*time.Second,
	}

	for input, expected := range cases {
		actual, err := parseDuration(input)
		if err != nil {
			t.Fatal(err)
		}
		if actual != expected {
			t.Fatalf(
				"bad: %#v, expected: %#v, input: %v",
				actual,
				expected,
				input,
			)
		}
	}
}
