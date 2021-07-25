package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

/*
	## summary
*/
func findLHS(nums []int) int {
	ret := 0
	for i := 0; i < (1 << len(nums)); i++ {
		count, minVal, maxVal := 0, math.MaxInt32, math.MinInt32
		for j := 0; j < len(nums); j++ {
			if (i & (1 << j)) != 0 {
				minVal = min(minVal, nums[j])
				maxVal = max(maxVal, nums[j])
				count++
			}
		}
		if maxVal-minVal == 1 {
			ret = max(ret, count)
		}
	}
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestFindLHS(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{1, 3, 2, 2, 5, 2, 3, 7}, 5},
		{[]int{1, 2, 3, 4}, 2},
		{[]int{1, 1, 1, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := findLHS(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
