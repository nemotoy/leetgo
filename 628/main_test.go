package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
	## summary
	3つの積の最大数
*/
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	return max(nums[0]*nums[1]*nums[len(nums)-1], nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3])
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestMaximumProduct(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4}, 24},
		{[]int{-1, -2, -3}, -6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := maximumProduct(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
