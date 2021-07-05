package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
	## summary
	ソート配列と同インデックスの値が合致しない数
*/
func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	ret := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != expected[i] {
			ret++
		}
	}
	return ret
}

func TestHeightChecker(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{1, 1, 4, 2, 1, 3}, 3},
		{[]int{5, 1, 2, 3, 4}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := heightChecker(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
