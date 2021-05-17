package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 合計が最大になるサブ配列の和を返す
func maxSubArray(nums []int) int {
	result, sum := nums[0], 0
	for i := 0; i < len(nums); i++ {
		t := sum + nums[i]
		result = max(result, t)
		if t <= 0 {
			sum = 0

		} else {
			sum = t
		}
	}

	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestMaxArray(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{5, 4, -1, 7, 8},
			23,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := maxSubArray(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
