package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	Return the sum of all the unique elements of nums.
*/
func sumOfUnique(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, n := range nums {
		m[n] += 1
	}
	ret := 0
	for k, v := range m {
		if v == 1 {
			ret += k
		}
	}
	return ret
}

func sumOfUnique2(nums []int) int {
	tmp := [101]int{}
	for _, n := range nums {
		tmp[n] += 1
	}
	ret := 0
	for k, v := range tmp {
		if v == 1 {
			ret += k
		}
	}
	return ret
}

func TestSumOfUnique(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{1, 2, 3, 2}, 4}, {[]int{1, 1, 1, 1, 1}, 0}, {[]int{1, 2, 3, 4, 5}, 15},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := sumOfUnique2(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
