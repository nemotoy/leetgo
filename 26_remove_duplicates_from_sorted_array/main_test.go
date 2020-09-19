package main

import (
	"fmt"
	"testing"
)

/*
	## summary
*/
func removeDuplicates(nums []int) int {
	return 0
}

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{1, 1, 2},
			2,
		},
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			if got := removeDuplicates(tt.in); got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
