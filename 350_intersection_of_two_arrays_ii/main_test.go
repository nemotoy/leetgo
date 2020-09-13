package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	2つの配列に含まれる整数を配列にして返す。同値は個別に扱う。
	- 配列の長さを短い方を元に、iterationする。
	- 単純化のため、返す配列は新規で割り当てる。

	- 返す配列は、長さの短い方から長い方に存在しない要素を削除し、返す方法も検討可能。
*/
func intersect(nums1 []int, nums2 []int) []int {

	n1Size := len(nums1)
	n2Size := len(nums2)
	result := []int{}

	if n1Size >= n2Size {
		for _, num2 := range nums2 {
			for i1, num1 := range nums1 {
				if num2 == num1 {
					// nums1から該当要素を除外する
					nums1 = append(nums1[:i1], nums1[i1+1:]...)
					result = append(result, num2)
					break
				}
			}
		}
	} else {
		for _, num1 := range nums1 {
			for i2, num2 := range nums2 {
				if num2 == num1 {
					nums2 = append(nums2[:i2], nums2[i2+1:]...)
					result = append(result, num1)
					break
				}
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
		{
			[]int{1, 2},
			[]int{1, 1},
			[]int{1},
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
