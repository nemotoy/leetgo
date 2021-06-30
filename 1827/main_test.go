package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	nums[i] < nums[i+1] for all 0 <= i < nums.length-1
*/

func minOperations(nums []int) int {
	ret := 0
	for i := 0; i < len(nums)-1; i++ {
		for nums[i] >= nums[i+1] {
			nums[i+1] += 1
			ret++
		}
	}
	return ret
}

func minOperations2(nums []int) int {
	ret := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			// 等価も含むので1を加算する
			ret += nums[i] - nums[i+1] + 1
			nums[i+1] = nums[i] + 1
		}
	}
	return ret
}

func TestMinOperations(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{1, 1, 1}, 3},
		{[]int{1, 5, 2, 4, 1}, 14},
		{[]int{8}, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := minOperations2(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
