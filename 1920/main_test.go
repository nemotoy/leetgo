package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, num := range nums {
		ans[i] = nums[num]
	}
	return ans
}

func TestBuildArray(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{0, 2, 1, 5, 3, 4},
			[]int{0, 1, 2, 4, 5, 3},
		},
		{
			[]int{5, 0, 1, 2, 3, 4},
			[]int{4, 5, 0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := buildArray(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
