package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	n == mat.len == mat[i].len
*/
func diagonalSum(mat [][]int) int {
	ret, n := 0, len(mat)
	for i := 0; i < n; i++ {
		ret += mat[i][i] + mat[i][n-i-1]
	}
	// If mat.length is odd, subtract the center number from result
	if n%2 == 1 {
		ret -= mat[n/2][n/2]
	}
	return ret
}

func TestDiagonalSum(t *testing.T) {
	tests := []struct {
		in  [][]int
		out int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			25,
		},
		{
			[][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			8,
		},
		{
			[][]int{
				{5},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := diagonalSum(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
