package main

import (
	"testing"
)

/*
  ## summary
*/
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestFindLUSlength(t *testing.T) {
	tests := []struct {
		in1 string
		in2 string
		out int
	}{
		{
			"aba", "cdc", 3,
		},
		{
			"aaa", "bbb", 3,
		},
		{
			"aaa", "aaa", -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in1+tt.in2, func(t *testing.T) {
			got := findLUSlength(tt.in1, tt.in2)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
