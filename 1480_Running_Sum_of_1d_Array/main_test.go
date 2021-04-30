package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	配列の各要素に0~indexまで加算した配列を返す
*/
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func TestRunningSum(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{1, 3, 6, 10},
		},
		{
			[]int{1, 1, 1, 1, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{3, 1, 2, 10, 1},
			[]int{3, 4, 6, 16, 17},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := runningSum(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
