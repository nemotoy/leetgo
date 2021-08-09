package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func minMoves(nums []int) int {
	minVal, sumVal := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		minVal = min(minVal, nums[i])
		sumVal += nums[i]
	}
	return sumVal - minVal*len(nums)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestMinMoves(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{1, 1, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := minMoves(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
