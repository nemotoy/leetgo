package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
	## summary
	非昇順な回答を返す
*/
func minSubsequence(nums []int) []int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums)
	ret := make([]int, 0, len(nums))
	tmp := 0
	for _, v := range nums {
		ret = append(ret, v)
		tmp += v
		if tmp > sum-tmp {
			break
		}
	}
	return ret
}

func TestMinSubsequence(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{4, 3, 10, 9, 8},
			[]int{10, 9},
		},
		{
			[]int{4, 4, 7, 6, 7},
			[]int{7, 7, 6},
		},
		{
			[]int{6},
			[]int{6},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := minSubsequence(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
