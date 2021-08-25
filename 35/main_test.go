package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
 */
func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if target == num || target < num {
			return i
		}
	}
	return len(nums)
}

func searchInsert2(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %d", tt.in, tt.in2), func(t *testing.T) {
			got := searchInsert2(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
