package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	nums: half odd, other even
	Sort this conditions nums[i] is odd, i is odd, and whenever nums[i] is even, i is even.
*/
func sortArrayByParityII(nums []int) []int {
	ret := make([]int, len(nums))
	// even cursor starts at zero
	oddInd, evenInd := 1, 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num%2 == 0 {
			ret[evenInd] = num
			evenInd += 2
		} else {
			ret[oddInd] = num
			oddInd += 2
		}
	}
	return ret
}

func TestSortArrayByParityII(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{4, 2, 5, 7},
			[]int{4, 5, 2, 7},
		},
		{
			[]int{2, 3},
			[]int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := sortArrayByParityII(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
