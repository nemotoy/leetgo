package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	discount if j is the minimum index such that j > i and prices[j] <= prices[i]
*/
func finalPrices(prices []int) []int {
	ret := make([]int, len(prices))
	for i, p1 := range prices {
		for _, p2 := range prices[i+1:] {
			if p1 >= p2 {
				p1 -= p2
				break
			}
		}
		ret[i] = p1
	}
	return ret
}

func TestFinalPrices(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{8, 4, 6, 2, 3},
			[]int{4, 2, 4, 2, 3},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{10, 1, 1, 6},
			[]int{9, 0, 1, 6},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := finalPrices(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
