package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func shuffle(nums []int, n int) []int {
	return []int{}
}

func TestShuffle(t *testing.T) {
	tests := []struct {
		in  []int
		inN int
		out []int
	}{
		{
			[]int{2, 5, 1, 3, 4, 7},
			3,
			[]int{2, 3, 5, 4, 1, 7},
		},
		{
			[]int{1, 2, 3, 4, 4, 3, 2, 1},
			4,
			[]int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			[]int{1, 1, 2, 2},
			2,
			[]int{1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.in, tt.inN), func(t *testing.T) {
			got := shuffle(tt.in, tt.inN)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
