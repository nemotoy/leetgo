package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	islandに接する場合周線を減らす
*/
func islandPerimeter(grid [][]int) int {
	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// if cell is an island
			if grid[i][j] == 1 {
				// max perimeter of a cell is four.
				ret += 4
				// row
				if i > 0 && grid[i-1][j] == 1 {
					ret -= 2
				}
				// col
				if j > 0 && grid[i][j-1] == 1 {
					ret -= 2
				}
			}
		}
	}
	return ret
}

func TestIslandPerimeter(t *testing.T) {
	tests := []struct {
		in  [][]int
		out int
	}{
		{
			[][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			}, 16,
		},
		{
			[][]int{{1}}, 4,
		},
		{
			[][]int{{1, 0}}, 4,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := islandPerimeter(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
