package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	偶数、奇数の順にソートして返す。各要素は順不同。
*/
// nolint: gocritic
func sortArrayByParity(A []int) []int {
	result := make([]int, len(A))
	return result
}

func TestSortArrayByParity(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{3, 1, 2, 4},
			[]int{2, 4, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := sortArrayByParity(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
