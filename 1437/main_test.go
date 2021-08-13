package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func kLengthApart(nums []int, k int) bool {
	// initialize the counter of zero to k to pass if the first element is one.
	count := k
	for _, num := range nums {
		if num == 1 {
			if count < k {
				return false
			}
			// reset the counter to zero.
			count = 0
		} else {
			count++
		}
	}
	return true
}

func TestKLengthApart(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out bool
	}{
		{
			[]int{1, 0, 0, 0, 1, 0, 0, 1}, 2, true,
		},
		{
			[]int{1, 0, 0, 1, 0, 1}, 2, false,
		},
		{
			[]int{1, 1, 1, 1, 1}, 0, true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.in, tt.in2), func(t *testing.T) {
			got := kLengthApart(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
