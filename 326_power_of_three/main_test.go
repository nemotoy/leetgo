package main

import (
	"fmt"
	"testing"
)

/*
	## summary
*/
func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	// 3で割り続ける。最終結果が1（3/3=1）か否かで評価する
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
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
