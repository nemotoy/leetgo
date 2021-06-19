package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	candies[i] + extraCandies > each candies
*/
func kidsWithCandies(candies []int, extraCandies int) []bool {
	ret := make([]bool, len(candies))
	for i, c := range candies {
		tmp := true
		for _, c2 := range candies {
			if c+extraCandies < c2 {
				tmp = false
				break
			}
		}
		ret[i] = tmp
	}
	return ret
}

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out []bool
	}{
		{
			[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true},
		},
		{
			[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.in, tt.in2), func(t *testing.T) {
			got := kidsWithCandies(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
