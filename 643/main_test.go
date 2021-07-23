package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxVal := float64(sum)
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		maxVal = max(maxVal, float64(sum))
	}
	return maxVal / float64(k)
}

func max(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out float64
	}{
		{
			[]int{1, 12, -5, -6, 50, 3}, 4, 12.75000,
		},
		{
			[]int{5}, 1, 5.00000,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.in, tt.in2), func(t *testing.T) {
			got := findMaxAverage(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
