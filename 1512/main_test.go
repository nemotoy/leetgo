package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	nums[i] == nums[j]
*/
func numIdenticalPairs(nums []int) int {
	ret, count := 0, [101]int{}
	for i := 0; i < len(nums); i++ {
		ret += count[nums[i]]
		count[nums[i]]++
	}
	return ret
}

func TestNumIdenticalPairs(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{1, 2, 3, 1, 1, 3},
			4,
		},
		{
			[]int{1, 1, 1, 1},
			6,
		},
		{
			[]int{1, 2, 3},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := numIdenticalPairs(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
