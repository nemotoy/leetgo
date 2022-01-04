package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	３つ中2つ以上の配列に含まれる整数。値は任意の順でよい。
*/
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	count := [3][101]int{{}}
	for _, n := range nums1 {
		count[0][n] = 1
	}
	for _, n := range nums2 {
		count[1][n] = 1
	}
	for _, n := range nums3 {
		count[2][n] = 1
	}
	ret := []int{}
	for i := 0; i < 101; i++ {
		if count[0][i]+count[1][i]+count[2][i] > 1 {
			ret = append(ret, i)
		}
	}
	return ret
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestTwoOutOfThree(t *testing.T) {
	tests := []struct {
		in  []int
		in2 []int
		in3 []int
		out []int
	}{
		{
			[]int{1, 1, 3, 2},
			[]int{2, 3},
			[]int{3},
			[]int{2, 3},
		},
		{
			[]int{3, 1},
			[]int{2, 3},
			[]int{1, 2},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v,%v,%v", tt.in, tt.in2, tt.in3), func(t *testing.T) {
			got := twoOutOfThree(tt.in, tt.in2, tt.in3)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
