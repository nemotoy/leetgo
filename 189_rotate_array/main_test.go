package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func rotate(nums []int, k int) {

}

func TestRotate(t *testing.T) {
	tests := []struct {
		in1 []int
		in2 int
		out []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums: %v, k: %d", tt.in1, tt.in2), func(t *testing.T) {
			rotate(tt.in1, tt.in2)
			got := tt.in1
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
