package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func projectionArea(grid [][]int) int {
	l := len(grid)
	ret := 0
	for i := 0; i < l; i++ {
		bestRow, bestCol := 0, 0
		for j := 0; j < l; j++ {
			if grid[i][j] > 0 {
				ret++
			}
			bestRow = max(bestRow, grid[i][j])
			bestCol = max(bestCol, grid[j][i])
		}
		ret += bestRow + bestCol
	}
	return ret
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestProjectionArea(t *testing.T) {
	tests := []struct {
		in  [][]int
		out int
	}{
		{
			[][]int{{1, 2}, {3, 4}}, 17,
		},
		{
			[][]int{{2}}, 5,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := projectionArea(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
