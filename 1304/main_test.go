package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	合計が0になるn個の一意の整数の組み合わせ
*/
func sumZero(n int) []int {
	ret := make([]int, n)
	start, end := 0, n-1
	for start < end {
		ret[start] = -end
		ret[end] = end
		start++
		end--
	}
	return ret
}

func TestSumZero(t *testing.T) {
	tests := []struct {
		in  int
		out []int
	}{
		{5, []int{-4, -3, 0, 3, 4}}, // or []int{-7, -1, 1, 3, 4}
		{3, []int{-2, 0, 2}},        // or []int{-1, 0, 1}
		{1, []int{0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("n: %d", tt.in), func(t *testing.T) {
			got := sumZero(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
