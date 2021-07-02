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

// binary search
func countNegatives2(grid [][]int) int {
	ret := 0
	for i := 0; i < len(grid); i++ {
		start, end := 0, len(grid[i])-1
		for start <= end {
			mid := (start + end) / 2
			if grid[i][mid] < 0 {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
		ret += len(grid[i]) - start
	}
	return ret
}

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
			got := countNegatives2(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
