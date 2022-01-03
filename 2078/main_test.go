package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	異なる色の最大距離

	左右ぞれぞれから異なる色を探して、両者の距離を比較する
*/
func maxDistance(colors []int) int {
	n, i, j := len(colors), 0, len(colors)-1
	for colors[0] == colors[j] {
		j--
	}
	for colors[n-1] == colors[i] {
		i++
	}
	return max(n-1-i, j)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestMaxDistance(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{1, 1, 1, 6, 1, 1, 1}, 3},
		{[]int{1, 8, 3, 8, 3}, 4},
		{[]int{0, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := maxDistance(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
