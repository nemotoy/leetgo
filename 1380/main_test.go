package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	A lucky number is an element of the matrix such that it is the minimum element in its row and maximum in its column.
*/
func luckyNumbers(matrix [][]int) []int {
	criteria := make(map[int]int)
	for _, row := range matrix {
		min := row[0]
		for _, ele := range row[1:] {
			if min > ele {
				min = ele
			}
		}
		criteria[min] += 1
	}

	for i := 0; i < len(matrix[0]); i++ {
		max := 0
		for j := 0; j < len(matrix); j++ {
			if max < matrix[j][i] {
				max = matrix[j][i]
			}
		}
		criteria[max] += 1
	}
	ret := make([]int, 0, len(criteria))
	for k, v := range criteria {
		if v == 2 {
			ret = append(ret, k)
		}
	}
	return ret
}

func TestLuckyNumbers(t *testing.T) {
	tests := []struct {
		in  [][]int
		out []int
	}{
		{
			[][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}, []int{15},
		},
		{
			[][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}, []int{12},
		},
		{
			[][]int{{7, 8}, {1, 2}}, []int{7},
		},
		{
			[][]int{{36376, 85652, 21002, 4510}, {68246, 64237, 42962, 9974}, {32768, 97721, 47338, 5841}, {55103, 18179, 79062, 46542}}, []int{},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := luckyNumbers(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
