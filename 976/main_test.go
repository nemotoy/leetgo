package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
	## summary
	面積が0でない三角形の周囲長
*/
func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i > 1; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

func TestLargestPerimeter(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{2, 1, 2}, 5},
		{[]int{1, 2, 1}, 0},
		{[]int{3, 2, 3, 4}, 10},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := largestPerimeter(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
