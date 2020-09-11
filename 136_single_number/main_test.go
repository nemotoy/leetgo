package main

import (
	"fmt"
	"testing"
)

func singleNumber(nums []int) int {

	i := 0
	up := 0
	n := len(nums)
	for i < n {
		switch {
		case up < nums[i]:
			up = nums[i]
		case up == nums[i]:
			up = 0
		case up > nums[i]:
			if up == 0 {
				up = nums[i]
			}
		}
		i++
	}
	return up
}

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{2, 2, 1}, 1,
		},
		{
			[]int{4, 1, 2, 1, 2}, 4,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := singleNumber(tt.in)
			if got != tt.out {
				t.Errorf("got: %d, want: %d", got, tt.out)
			}
		})
	}
}
