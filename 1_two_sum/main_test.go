package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func twoSum(nums []int, target int) []int {
	return []int{}
}

func Test(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums:%v target:%d", tt.in, tt.in2), func(t *testing.T) {
			got := twoSum(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
