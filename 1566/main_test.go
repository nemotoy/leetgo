package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	k回以上繰り返される長さmのパターンが存在する場合はtrueを返し、存在しない場合はfalseを返す。
*/
func containsPattern(arr []int, m int, k int) bool {
	count := 0
	for i := 0; i+m < len(arr); i++ {
		// 部分配列の長さはmなので、i+mが同値になる
		if arr[i] == arr[i+m] {
			count++
		} else {
			count = 0
		}
		if count == (k-1)*m {
			return true
		}
	}
	return false
}

func TestContainsPattern(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		in3 int
		out bool
	}{
		{[]int{1, 2, 4, 4, 4, 4}, 1, 3, true},
		{[]int{1, 2, 1, 2, 1, 1, 1, 3}, 2, 2, true},
		{[]int{1, 2, 1, 2, 1, 3}, 2, 3, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d, %d", tt.in, tt.in2, tt.in3), func(t *testing.T) {
			got := containsPattern(tt.in, tt.in2, tt.in3)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
