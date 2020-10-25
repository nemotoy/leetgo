package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	i < j < k and nums[i] < nums[k] < nums[j].
*/
func find132pattern(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] < nums[k] && nums[k] < nums[j] {
					return true
				}
			}
		}
	}
	return false
}

func TestFind132pattern(t *testing.T) {
	tests := []struct {
		in  []int
		out bool
	}{
		{
			[]int{1, 2, 3, 4}, false,
		},
		{
			[]int{3, 1, 4, 2}, true,
		},
		{
			[]int{-1, 3, 2, 0}, true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := find132pattern(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
