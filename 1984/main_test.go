package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
 */
func minimumDifference(nums []int, k int) int {
	// 昇順
	sort.Ints(nums)
	// 初期値
	ret := nums[k-1] - nums[0]
	for i := k; i < len(nums); i++ {
		ret = min(ret, nums[i]-nums[i-k+1])
	}
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestMinimumDifference(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out int
	}{
		{[]int{90}, 1, 0},
		{[]int{9, 4, 1, 7}, 2, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.in, tt.in2), func(t *testing.T) {
			got := minimumDifference(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
