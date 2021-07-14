package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func findMaxConsecutiveOnes(nums []int) int {
	cur := 0
	tmp := 0
	for _, n := range nums {
		if n == 1 {
			cur++
		} else {
			if tmp < cur {
				tmp = cur
			}
			cur = 0
		}
	}
	return max(cur, tmp)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{1, 1, 0, 1, 1, 1}, 3,
		},
		{
			[]int{1, 0, 1, 1, 0, 1}, 2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := findMaxConsecutiveOnes(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
