package main

import (
	"fmt"
	"testing"
)

/*
	## summary
*/
func isPowerOfThree(n int) bool {
	return false
}

func TestIsPowerOfThree(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{
			27, true,
		},
		{
			0, false,
		},
		{
			9, true,
		},
		{
			45, false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := isPowerOfThree(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
