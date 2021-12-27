package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
	## summary
	numsをinc orderした結果において、targetと値が同値であるindexのリストを返す
*/
func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	ret := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			ret = append(ret, i)
		}
	}
	return ret
}

func TestTargetIndices(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out []int
	}{
		{
			[]int{1, 2, 5, 2, 3},
			2,
			[]int{1, 2},
		},
		{
			[]int{1, 2, 5, 2, 3},
			3,
			[]int{3},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.in, tt.in2), func(t *testing.T) {
			got := targetIndices(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
