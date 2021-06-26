package main

import (
	"fmt"
	"testing"
)

/*
	## summary
	gainの各要素を加算した各結果の最大値
*/
func largestAltitude(gain []int) int {
	hights := make([]int, len(gain)+1)
	hights[0] = 0
	for i, h := range gain {
		hights[i+1] = hights[i] + h
	}
	ret := 0
	for _, h := range hights[1:] {
		if ret < h {
			ret = h
		}
	}
	return ret
}

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{-5, 1, 5, 0, -7}, 1,
		},
		{
			[]int{-4, -3, -2, -1, 4, 3, 2}, 0,
		},
		{
			[]int{28, 0, -8, -99, 11, 62, -35, 68, 2, 12, -71, 13, 66, -28}, 49,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := largestAltitude(tt.in)
			if got != tt.out {
				t.Errorf("got: %d, want: %d", got, tt.out)
			}
		})
	}
}
