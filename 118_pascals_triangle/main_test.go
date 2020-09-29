package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	if numRows == 0 {
		return result
	}
	if numRows == 1 {
		result[0] = []int{1}
		return result
	}
	// numRows >= 2 は確定。
	result[0] = []int{1}
	result[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		// 1つ前の要素を取得する
		prev := result[i-1]
		// 隣り合う要素を加算し、1要素として追加する。
		mid := []int{}
		for pi := 0; pi <= len(prev); pi++ {
			// 最初と最後はその要素を追加する。それ以外は前要素と加算した値を追加する。
			switch pi {
			case 0:
				mid = append(mid, prev[pi])
			case len(prev):
				mid = append(mid, prev[pi-1]) // piの基底値が0なので、pi-1が最終要素。
			default:
				mid = append(mid, prev[pi]+prev[pi-1])
			}
		}
		result[i] = mid
	}
	return result
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		in  int
		out [][]int
	}{
		{
			5,
			[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := generate(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
