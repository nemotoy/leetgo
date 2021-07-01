package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func minTimeToVisitAllPoints(points [][]int) int {
	ret := 0
	for i := 0; i < len(points)-1; i++ {
		curx, cury := points[i][0], points[i][1]
		dstx, dsty := points[i+1][0], points[i+1][1]
		// 目的地と現在地の差の各絶対値で、最大の値を加算する
		ret += max(abs(dstx-curx), abs(dsty-cury))
	}
	return ret
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func abs(x int) int {
	if x < 1 {
		return -x
	}
	return x
}

func TestDiagonalSum(t *testing.T) {
	tests := []struct {
		in  [][]int
		out int
	}{
		{
			[][]int{{1, 1}, {3, 4}, {-1, 0}}, 7,
		},
		{
			[][]int{{3, 2}, {-2, 2}}, 5,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := minTimeToVisitAllPoints(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
