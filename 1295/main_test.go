package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	how many of them contain an even number of digits.
*/
func findNumbers(nums []int) int {
	ret := 0
	for _, n := range nums {
		digits := 0
		for n != 0 {
			n /= 10
			digits++
		}
		// if event number this digits
		if digits%2 == 0 {
			ret++
		}
	}
	return ret
}

func TestFindNumbers(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{12, 345, 2, 6, 7896}, 2},
		{[]int{555, 901, 482, 1771}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := findNumbers(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
