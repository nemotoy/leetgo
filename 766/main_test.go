package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func isToeplitzMatrix(matrix [][]int) bool {
	groups := make(map[int]int)
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if _, ok := groups[row-col]; !ok {
				groups[row-col] = matrix[row][col]
			} else if groups[row-col] != matrix[row][col] {
				return false
			}
		}
	}
	return true
}

// 左上の要素と同値か
func isToeplitzMatrix2(matrix [][]int) bool {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if row > 0 && col > 0 && matrix[row-1][col-1] != matrix[row][col] {
				return false
			}
		}
	}
	return true
}

func TestIsToeplitzMatrix(t *testing.T) {
	tests := []struct {
		in  [][]int
		out bool
	}{
		{
			[][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}, true,
		},
		{
			[][]int{{1, 2}, {2, 2}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := isToeplitzMatrix2(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
