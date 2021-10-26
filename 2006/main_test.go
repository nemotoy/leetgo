package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	Return the number of pairs (i, j) where i < j such that |nums[i] - nums[j]| == k.
	The value of |x| is defined as:
	 x if x >= 0.
	-x if x < 0.
*/
func countKDifference(nums []int, k int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i+1; j++ {
			if abs(nums[i]-nums[j]) == k {
				ret++
			}
		}
	}
	return ret
}

func countKDifference2(nums []int, k int) int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	ret := 0
	for _, v := range nums {
		c := v - k
		ret += m[c]
	}
	return ret
}

func abs(x int) int {
	if x < 1 {
		return -x
	}
	return x
}

func TestCountKDifference(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out int
	}{
		{
			[]int{1, 2, 2, 1}, 1, 4,
		},
		{
			[]int{1, 3}, 3, 0,
		},
		{
			[]int{3, 2, 1, 5, 4}, 2, 3,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := countKDifference2(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
