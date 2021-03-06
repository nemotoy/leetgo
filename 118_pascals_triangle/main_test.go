package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	- 内部の配列数は与えられた数値と同値
	- 各配列の要素の最初・最後は1で確定。それ以外の要素は自身のindexに該当する要素と1つ前の要素を加算した値を持つ。
*/
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	result := make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		// 1つ前の要素を取得する
		prev := result[i-1]
		// 隣り合う要素を加算し、1要素として追加する。容量は1つ前の要素より1大きいと確定しているので予め決定する。
		mid := make([]int, len(prev)+1)
		for pi := 0; pi <= len(prev); pi++ {
			// 最初と最後はその要素を追加する。それ以外は前要素と加算した値を追加する。容量は確定しているので、appendは使わない。
			switch pi {
			case 0, len(prev):
				mid[pi] = 1 // 最初と最後は1で確定なので、prevを参照しない。
			default:
				mid[pi] = prev[pi] + prev[pi-1]
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
