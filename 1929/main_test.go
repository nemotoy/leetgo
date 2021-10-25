package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
  ans[i] == nums[i] and ans[i + n] == nums[i]
*/
func getConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}
	return ans
}

func TesGetConcatenation(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{1, 2, 1},
			[]int{1, 2, 1, 1, 2, 1},
		},
		{
			[]int{1, 3, 2, 1},
			[]int{1, 3, 2, 1, 1, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := getConcatenation(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
