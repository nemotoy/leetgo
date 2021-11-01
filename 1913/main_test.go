package main

import (
	"fmt"
	"reflect"
	"testing"
)

func maxProductDifference(nums []int) int {
	max1, max2, min1, min2 := 0, 0, 10001, 10001
	for _, num := range nums {
		if num >= max1 {
			max2 = max1
			max1 = num
		} else if num >= max2 {
			max2 = num
		}
		if num <= min1 {
			min2 = min1
			min1 = num
		} else if num <= min2 {
			min2 = num
		}
	}
	return max1*max2 - min1*min2
}

func TestMaxProductDifference(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{5, 6, 2, 7, 4}, 34},
		{[]int{4, 2, 5, 9, 7, 4, 8}, 64},
		{[]int{1, 6, 7, 5, 2, 4, 10, 6, 4}, 68},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := maxProductDifference(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
