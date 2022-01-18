package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func checkValid(matrix [][]int) bool {
	l := len(matrix)
	for i := 0; i < l; i++ {
		row := make(map[int]int, l)
		col := make(map[int]int, l)
		for j := 0; j < l; j++ {
			row[matrix[i][j]]++
			col[matrix[j][i]]++
		}
		if len(row) != l || len(col) != l {
			return false
		}
	}
	return true
}

func TestCheckValid(t *testing.T) {
	tests := []struct {
		in  [][]int
		out bool
	}{
		{
			[][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}}, true,
		},
		{
			[][]int{{1, 1, 1}, {1, 2, 3}, {1, 2, 3}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := checkValid(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
