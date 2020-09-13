package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func intersect(nums1 []int, nums2 []int) []int {
	return []int{}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		in1 []int
		in2 []int
		out []int
	}{
		{
			[]int{1, 2, 2, 1},
			[]int{2, 2},
			[]int{2, 2},
		},
		{
			[]int{4, 9, 5},
			[]int{9, 4, 9, 8, 4},
			[]int{4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := intersect(tt.in2, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
