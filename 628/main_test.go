// nolint:gocritic
package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"testing"
)

/*
	## summary
	3つの積の最大数
*/
// sort
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	return max(nums[0]*nums[1]*nums[len(nums)-1], nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3])
}

//
func maximumProduct2(nums []int) int {
	min1, min2 := math.MaxInt32, math.MaxInt32
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	for _, n := range nums {
		if n <= min1 {
			min1, min2 = n, min1
		} else if n <= min2 {
			min2 = n
		}
		if n >= max1 {
			max1, max2, max3 = n, max1, max2
		} else if n >= max2 {
			max2, max3 = n, max2
		} else if n >= max3 {
			max3 = n
		}
	}
	return max(min1*min2*max1, max1*max2*max3)
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
			got := maximumProduct2(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
