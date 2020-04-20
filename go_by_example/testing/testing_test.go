package main

import (
	"fmt"
	"testing"
)

// Now test code is in main package, but typically it should be inside the same package as the code it tests

// Simple function to test
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error reports of test failure but keeps tests running
		// t.Fail reports of a failure and stops all the tests
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Test inputs can be repetitive, so it's better to use table driven style
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// One loop goes over every test and performs a subtest
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		// t.Run runs a new subtest
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
