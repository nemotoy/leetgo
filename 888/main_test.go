package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	aliceSizes[i], bobSizes[j] を1度交換し、aliceSizes[i...n] = bobSizes[j...l]
	0: alice exchange, 1: bob exchange
	alice candy x, bob candy y
	Sa - x + y = Sb - y + x
	y = x + (Sb-Sa)/2
*/
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sa, sb := 0, 0
	for _, c := range aliceSizes {
		sa += c
	}
	for _, c := range bobSizes {
		sb += c
	}
	delta := (sb - sa) / 2

	setB := make(map[int]int)
	for _, c := range bobSizes {
		setB[c] = 1
	}
	for _, c := range aliceSizes {
		if _, ok := setB[c+delta]; ok {
			return []int{c, c + delta}
		}
	}
	return nil
}

func TestFairCandySwap(t *testing.T) {
	tests := []struct {
		in1 []int
		in2 []int
		out []int
	}{
		{
			[]int{1, 1},
			[]int{2, 2},
			[]int{1, 2},
		},
		{
			[]int{1, 2},
			[]int{2, 3},
			[]int{1, 2},
		},
		{
			[]int{2},
			[]int{1, 3},
			[]int{2, 3},
		},
		{
			[]int{1, 2, 5},
			[]int{2, 4},
			[]int{5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := fairCandySwap(tt.in1, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
