package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func smallerNumbersThanCurrent(nums []int) []int {
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
		ret[i] = count
	}
	return ret
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{8, 1, 2, 2, 3},
			[]int{4, 0, 1, 1, 3},
		},
		{
			[]int{6, 5, 4, 8},
			[]int{2, 1, 0, 3},
		},
		{
			[]int{7, 7, 7, 7},
			[]int{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := smallerNumbersThanCurrent(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
