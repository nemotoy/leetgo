package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func countNegatives(grid [][]int) int {
	ret := 0
	for _, matrix := range grid {
		for _, v := range matrix {
			if v < 0 {
				ret++
			}
		}
	}
	return ret
}

// func max(x, y int) int {
// 	if x < y {
// 		return y
// 	}
// 	return x
// }

// func abs(x int) int {
// 	if x < 1 {
// 		return -x
// 	}
// 	return x
// }

func TestCountNegatives(t *testing.T) {
	tests := []struct {
		in  [][]int
		out int
	}{
		{
			[][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}, 8,
		},
		{
			[][]int{{3, 2}, {1, 0}}, 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := countNegatives(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
