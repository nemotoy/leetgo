package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	2つの配列に含まれる整数を配列にして返す。
	- 配列の長さを小さい方を元に、iterationする。
*/
func intersect(nums1 []int, nums2 []int) []int {

	n1Size := len(nums1)
	n2Size := len(nums2)
	result := []int{}

	if n1Size >= n2Size {
		for _, num2 := range nums2 {
			f := false
			for _, num1 := range nums1 {
				if num2 == num1 {
					f = true
					break
				}
			}
			if f {
				result = append(result, num2)
			}
		}
	} else {
		for _, num1 := range nums1 {
			f := false
			for _, num2 := range nums2 {
				if num2 == num1 {
					f = true
					break
				}
			}
			if f {
				result = append(result, num1)
				fmt.Println(result)
			}
		}
	}

	return result
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
			got := intersect(tt.in1, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
