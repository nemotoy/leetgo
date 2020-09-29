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
	result := [][]int{}
	if numRows == 0 {
		return result
	}
	for i := 1; i <= numRows; i++ {
		// 1つ前の要素を取得する
		prev := result[i-1]
		fmt.Println(prev)
		// 隣り合う要素を加算し、1要素として追加する。
		mid := []int{}
		for pi := 0; pi < len(prev); pi++ {
			mid = append(mid, prev[pi]+prev[pi+1])
		}
		fmt.Println(mid)
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
