package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func matrixReshape(mat [][]int, r int, c int) [][]int {
	n, m := len(mat), len(mat[0])
	if n*m != r*c {
		return mat
	}
	ret := make([][]int, r)
	for i := range ret {
		ret[i] = make([]int, c)
	}
	for i := 0; i < r*c; i++ {
		ret[i/c][i%c] = mat[i/m][i%m]
	}
	return ret
}

func invert(n int) int {
	if n == 0 {
		return 1
	}
	return 0
}

func TestMatrixReshape(t *testing.T) {
	tests := []struct {
		in  [][]int
		in2 int
		in3 int
		out [][]int
	}{
		{
			[][]int{
				{1, 2},
				{3, 4},
			},
			1,
			4,
			[][]int{
				{1, 2, 3, 4},
			},
		},
		{
			[][]int{
				{1, 2},
				{3, 4},
			},
			2,
			4,
			[][]int{
				{1, 2},
				{3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d, %d", tt.in, tt.in2, tt.in3), func(t *testing.T) {
			got := matrixReshape(tt.in, tt.in2, tt.in3)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
