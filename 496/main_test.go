package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0, len(nums1))
	for _, num1 := range nums1 {
		added := -1
		for j, num2 := range nums2 {
			if num1 == num2 {
				// search next greater element
				for _, n := range nums2[j:] {
					if num2 < n {
						added = n
						break
					}
				}
			}
		}
		ans = append(ans, added)
	}
	return ans
}

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		in1 []int
		in2 []int
		out []int
	}{
		{
			[]int{4, 1, 2},
			[]int{1, 3, 4, 2},
			[]int{-1, 3, -1},
		},
		{
			[]int{2, 4},
			[]int{1, 2, 3, 4},
			[]int{3, -1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("in1: %v, in2: %v", tt.in1, tt.in2), func(t *testing.T) {
			got := nextGreaterElement(tt.in1, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
