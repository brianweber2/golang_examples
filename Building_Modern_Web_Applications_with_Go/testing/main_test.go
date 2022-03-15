package main

import (
	"testing"
)

var tests = []struct {
	name     string
	dividend float32
	divsor   float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divsor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error, but did not get one")
			}
		} else {
			if err != nil {
				t.Error("got an error when we should not have", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}
