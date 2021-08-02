package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	row -> col
*/
func transpose(matrix [][]int) [][]int {
	ret := make([][]int, len(matrix[0]))
	// initialize an instance that is type of slice of slice
	for i := 0; i < len(matrix[0]); i++ {
		ret[i] = make([]int, len(matrix))
	}
	for i, row := range matrix {
		for j := 0; j < len(matrix[0]); j++ {
			ret[j][i] = row[j]
		}
	}
	return ret
}

func TestTranspose(t *testing.T) {
	tests := []struct {
		in  [][]int
		out [][]int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := transpose(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
