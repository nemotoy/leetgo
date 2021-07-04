package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	その配列内のすべての要素をその右側の要素の中で最大の要素に置き換え、最後の要素を-1に置き換える
*/
func replaceElements(arr []int) []int {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		max := 0
		for _, v := range arr[i+1:] {
			if max < v {
				max = v
			}
		}
		arr[i] = max
	}
	arr[l-1] = -1
	return arr
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestReplaceElements(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{17, 18, 5, 4, 6, 1},
			[]int{18, 6, 6, 6, 1, -1},
		},
		{
			[]int{400},
			[]int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := replaceElements(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
