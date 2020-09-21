package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	階段の上階に登る。一度に1か2階段登れる。取り得る選択肢を返す。
*/
func climbStairs(n int) int {
	m := make([]int, n)
	return calc(n, 0, m)
}

func calc(n, i int, m []int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if m[i] > 0 {
		return m[i]
	}
	m[i] = calc(n, i+1, m) + calc(n, i+2, m)
	return m[i]
}

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{
			2, 2,
		},
		{
			3, 3,
		},
		{
			4, 5,
		},
		{
			5, 8,
		},
		{
			6, 13,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := climbStairs(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
