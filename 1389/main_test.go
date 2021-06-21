package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func createTargetArray(nums []int, index []int) []int {
	ret := make([]int, len(nums))
	for i, ind := range index {
		if ind < i {
			for j := i; j > ind; j-- {
				ret[j] = ret[j-1]
			}
		}
		ret[ind] = nums[i]
	}
	return ret
}

func TestCreateTargetArray(t *testing.T) {
	tests := []struct {
		in1 []int
		in2 []int
		out []int
	}{
		{
			[]int{0, 1, 2, 3, 4},
			[]int{0, 1, 2, 2, 1},
			[]int{0, 4, 1, 3, 2},
		},
		{
			[]int{1, 2, 3, 4, 0},
			[]int{0, 1, 2, 3, 0},
			[]int{0, 1, 2, 3, 4},
		},
		{
			[]int{1},
			[]int{0},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums: %v, index: %v", tt.in1, tt.in2), func(t *testing.T) {
			got := createTargetArray(tt.in1, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
