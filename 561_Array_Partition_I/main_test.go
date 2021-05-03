package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
	## summary
*/
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func TestArrayPairSum(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{1, 4, 3, 2},
			4,
		},
		{
			[]int{6, 2, 6, 5, 1, 2},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := arrayPairSum(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
