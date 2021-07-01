package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	return the maximum value of (nums[i]-1)*(nums[j]-1)
*/
func maxProduct(nums []int) int {
	m, n := 0, 0
	for _, num := range nums {
		if m < num {
			n = m
			m = num
		} else if n < num {
			n = num
		}
	}
	return (m - 1) * (n - 1)
}

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{3, 4, 5, 2}, 12},
		{[]int{1, 5, 4, 5}, 16},
		{[]int{3, 7}, 12},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := maxProduct(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
