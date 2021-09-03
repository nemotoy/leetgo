package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
  in: 昇順整数の配列、整数
  out: arrに存在しないk番目の整数
*/
func findKthPositive(arr []int, k int) int {
	for i, n := 0, 1; n <= 1000; n++ {
		if i < len(arr) && arr[i] == n {
			i++
		} else {
			k--
			if k == 0 {
				return n
			}
		}
	}
	return 1000 + k
}

func TestFindKthPositive(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out int
	}{
		{[]int{2, 3, 4, 7, 11}, 5, 9},
		{[]int{1, 2, 3, 4}, 2, 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.in, tt.in2), func(t *testing.T) {
			got := findKthPositive(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
