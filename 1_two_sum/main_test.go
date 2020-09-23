package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i+1 < l; i++ {
		for j := i + 1; j < l; j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func Test(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 3},
			6,
			[]int{0, 2},
		},
		{
			[]int{2, 5, 5, 11},
			10,
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums:%v target:%d", tt.in, tt.in2), func(t *testing.T) {
			got := twoSum(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
